package lark

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	v1 "star/api/lark/v1"
	"star/internal/dao"
	"star/internal/logic/users"
	"star/internal/model/entity"
	"star/utility/lark"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
)

// ExportWords 导出单词到飞书
func (s *Lark) ExportWords(ctx context.Context, req *v1.ExportWordsReq) (res *v1.ExportWordsRes, err error) {
	res = &v1.ExportWordsRes{}

	// 1. 从 token 中获取用户 ID
	uid, err := users.New().GetUid(ctx)
	if err != nil {
		return nil, fmt.Errorf("获取用户信息失败: %w", err)
	}

	// 2. 查询用户的单词
	var words []*entity.Words
	err = dao.Words.Ctx(ctx).Where(dao.Words.Columns().Uid, uid).Scan(&words)
	if err != nil {
		return nil, fmt.Errorf("查询单词失败: %w", err)
	}

	if len(words) == 0 {
		res.Message = "没有找到单词数据"
		return res, nil
	}

	// 2. 准备表头和数据
	headers := []string{"单词", "定义", "例句", "中文翻译", "发音", "熟练度", "创建时间", "更新时间"}
	var rows [][]string
	for _, word := range words {
		rows = append(rows, []string{
			word.Word,
			word.Definition,
			word.ExampleSentence,
			word.ChineseTranslation,
			word.Pronunciation,
			strconv.Itoa(int(word.ProficiencyLevel)),
			word.CreatedAt.String(),
			word.UpdatedAt.String(),
		})
	}

	// 3. 使用默认或提供的多维表格
	config := gcfg.Instance()
	baseTokenVar, _ := config.Get(ctx, "lark.baseToken", "Cx6ubIEy3apIFVsssQdleh4tgeb")
	baseToken := baseTokenVar.String()
	if req.SpreadsheetToken != "" {
		baseToken = req.SpreadsheetToken
	}

	// 4. 追加到多维表格
	g.Log().Infof(ctx, "准备导出 %d 个单词到多维表格: %s", len(words), baseToken)
	client := lark.NewClient()
	err = client.AppendSpreadsheetRows(ctx, baseToken, "单词学习表", headers, rows)
	if err != nil {
		g.Log().Errorf(ctx, "导出到多维表格失败: %v", err)
		return nil, fmt.Errorf("导出到多维表格失败: %w", err)
	}

	res.Count = len(words)
	res.SpreadsheetUrl = "https://bytedance.sg.larkoffice.com/bitable/" + baseToken
	res.Message = fmt.Sprintf("成功导出 %d 个单词数据到多维表格", len(words))

	// 示例输出
	g.Log().Infof(ctx, "单词导出数据预览:")
	for i, row := range rows {
		if i < 3 {
			g.Log().Infof(ctx, "  %d. %s - %s", i+1, row[0], row[3])
		}
	}

	return res, nil
}

// SendReminder 发送学习提醒
func (s *Lark) SendReminder(ctx context.Context, req *v1.SendReminderReq) (res *v1.SendReminderRes, err error) {
	res = &v1.SendReminderRes{}

	// 1. 从 token 中获取用户 ID
	uid, err := users.New().GetUid(ctx)
	if err != nil {
		return nil, fmt.Errorf("获取用户信息失败: %w", err)
	}

	// 2. 查询用户信息
	var user *entity.Users
	err = dao.Users.Ctx(ctx).Where(dao.Users.Columns().Id, uid).Scan(&user)
	if err != nil || user == nil {
		return nil, fmt.Errorf("查询用户失败: %w", err)
	}

	// 3. 查询需要复习的单词 (熟练度 <= 3)
	var words []*entity.Words
	err = dao.Words.Ctx(ctx).
		Where(dao.Words.Columns().Uid, uid).
		WhereLTE(dao.Words.Columns().ProficiencyLevel, 3).
		Limit(10).
		Scan(&words)
	if err != nil {
		return nil, fmt.Errorf("查询单词失败: %w", err)
	}

	// 4. 构建提醒消息
	message := fmt.Sprintf("📚 【Star 单词学习】\n")
	message += fmt.Sprintf("👋 你好，%s\n", user.Username)
	if len(words) > 0 {
		message += fmt.Sprintf("\n📋 今天需要复习的单词 (%d个):\n", len(words))
		for i, word := range words {
			message += fmt.Sprintf("  %d. %s - %s (熟练度:%d)\n", i+1, word.Word, word.ChineseTranslation, word.ProficiencyLevel)
		}
	} else {
		message += "\n🎉 太棒了！你已经掌握了所有单词！\n"
	}
	message += "\n💪 继续加油学习！"

	// 5. 确定接收者 ID
	receiveID := req.ReceiveID
	if receiveID == "" {
		// 从配置文件获取默认接收者 ID
		config := gcfg.Instance()
		defaultID, _ := config.Get(ctx, "lark.defaultReceiveId", "")
		receiveID = defaultID.String()
		if receiveID == "" {
			g.Log().Infof(ctx, "请在请求中传入 receiveId 或在配置文件中设置默认值")
			res.Message = fmt.Sprintf("学习提醒准备成功，共 %d 个单词需要复习。请传入 receiveId 或设置默认值后重试", len(words))
			return res, nil
		}
		g.Log().Infof(ctx, "使用默认接收者: %s", receiveID)
	}

	// 6. 真正调用飞书 API 发送消息
	g.Log().Infof(ctx, "准备发送学习提醒给飞书用户: %s", receiveID)
	g.Log().Infof(ctx, "消息内容:\n%s", message)

	// 构建飞书消息格式（text 类型）
	msgContent := map[string]string{
		"text": message,
	}
	jsonContent, err := json.Marshal(msgContent)
	if err != nil {
		return nil, fmt.Errorf("构建消息内容失败: %w", err)
	}

	// 发送消息
	client := lark.NewClient()
	err = client.SendMessage(ctx, receiveID, "open_id", "text", string(jsonContent))
	if err != nil {
		g.Log().Errorf(ctx, "发送飞书消息失败: %v", err)
		return nil, fmt.Errorf("发送消息失败: %w", err)
	}

	res.Message = fmt.Sprintf("学习提醒已发送给飞书用户 %s，共 %d 个单词需要复习", receiveID, len(words))
	return res, nil
}
