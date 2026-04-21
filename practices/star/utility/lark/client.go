package lark

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
)

// Client 飞书 API 客户端
type Client struct {
	AppID             string
	AppSecret         string
	TenantAccessToken string
	TokenExpireTime   time.Time
}

// NewClient 创建飞书客户端
func NewClient() *Client {
	config := gcfg.Instance()
	appID, _ := config.Get(context.Background(), "lark.appId", "")
	appSecret, _ := config.Get(context.Background(), "lark.appSecret", "")

	return &Client{
		AppID:     appID.String(),
		AppSecret: appSecret.String(),
	}
}

// getTenantAccessToken 获取租户访问令牌
func (c *Client) getTenantAccessToken(ctx context.Context) (string, error) {
	if c.TenantAccessToken != "" && time.Now().Before(c.TokenExpireTime) {
		return c.TenantAccessToken, nil
	}

	url := "https://open.feishu.cn/open-apis/auth/v3/tenant_access_token/internal"
	payload := map[string]string{
		"app_id":     c.AppID,
		"app_secret": c.AppSecret,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var result struct {
		Code              int    `json:"code"`
		Msg               string `json:"msg"`
		TenantAccessToken string `json:"tenant_access_token"`
		Expire            int    `json:"expire"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}

	if result.Code != 0 {
		return "", fmt.Errorf("获取 token 失败: %s", result.Msg)
	}

	c.TenantAccessToken = result.TenantAccessToken
	c.TokenExpireTime = time.Now().Add(time.Duration(result.Expire) * time.Second)

	g.Log().Infof(ctx, "成功获取飞书 tenant_access_token")
	return c.TenantAccessToken, nil
}

// AppendSpreadsheetRows 追加数据到电子表格（多维表格）
func (c *Client) AppendSpreadsheetRows(ctx context.Context, spreadsheetToken, sheetName string, headers []string, rows [][]string) error {
	token, err := c.getTenantAccessToken(ctx)
	if err != nil {
		return err
	}

	// 1. 先获取表列表，找到正确的表 ID
	tableUrl := fmt.Sprintf("https://open.feishu.cn/open-apis/bitable/v1/apps/%s/tables", spreadsheetToken)
	req, err := http.NewRequestWithContext(ctx, "GET", tableUrl, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+token)

	httpClient := &http.Client{Timeout: 30 * time.Second}
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	g.Log().Infof(ctx, "获取表列表响应: %s", string(body))

	var tableList struct {
		Code int `json:"code"`
		Data struct {
			Items []struct {
				TableId string `json:"table_id"`
				Name    string `json:"name"`
			} `json:"items"`
		} `json:"data"`
		Msg string `json:"msg"`
	}

	if err := json.Unmarshal(body, &tableList); err != nil {
		g.Log().Errorf(ctx, "解析表列表失败: %v", err)
		return fmt.Errorf("解析表列表失败: %w", err)
	}

	if tableList.Code != 0 {
		return fmt.Errorf("获取表列表失败: code=%d, msg=%s", tableList.Code, tableList.Msg)
	}

	// 找到目标表 ID
	tableId := ""
	for _, table := range tableList.Data.Items {
		g.Log().Infof(ctx, "找到表: name=%s, id=%s", table.Name, table.TableId)
		if table.Name == sheetName {
			tableId = table.TableId
			break
		}
	}

	if tableId == "" {
		return fmt.Errorf("未找到表名 %s", sheetName)
	}

	g.Log().Infof(ctx, "使用表 ID: %s", tableId)

	// 2. 构建多维表格记录
	var records []map[string]interface{}
	for _, row := range rows {
		record := make(map[string]interface{})
		for i, header := range headers {
			if i < len(row) {
				record[header] = row[i]
			}
		}
		records = append(records, map[string]interface{}{
			"fields": record,
		})
	}

	url := fmt.Sprintf("https://open.feishu.cn/open-apis/bitable/v1/apps/%s/tables/%s/records/batch_create",
		spreadsheetToken, tableId)

	payload := map[string]interface{}{
		"records": records,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	g.Log().Infof(ctx, "准备写入 %d 条记录到多维表格 %s, 表名: %s, 表 ID: %s", len(rows), spreadsheetToken, sheetName, tableId)
	g.Log().Infof(ctx, "请求数据: %s", string(jsonData))

	req, err = http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err = httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	g.Log().Infof(ctx, "多维表格 API 响应状态: %d, 响应: %s", resp.StatusCode, string(body))

	var result struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data any    `json:"data"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		g.Log().Errorf(ctx, "解析响应失败: %v, 响应内容: %s", err, string(body))
		return fmt.Errorf("解析响应失败: %w", err)
	}

	if result.Code != 0 {
		return fmt.Errorf("追加数据失败: code=%d, msg=%s", result.Code, result.Msg)
	}

	g.Log().Infof(ctx, "成功追加 %d 条记录到多维表格", len(rows))
	return nil
}

// 辅助函数：将 []string 转换为 []interface{}
func interfaceToSlice(strs []string) []interface{} {
	result := make([]interface{}, len(strs))
	for i, s := range strs {
		result[i] = s
	}
	return result
}

// SendMessage 发送消息到用户
func (c *Client) SendMessage(ctx context.Context, receiveID, receiveIDType, msgType, content string) error {
	token, err := c.getTenantAccessToken(ctx)
	if err != nil {
		return err
	}

	url := "https://open.feishu.cn/open-apis/im/v1/messages?receive_id_type=" + receiveIDType

	payload := map[string]interface{}{
		"receive_id": receiveID,
		"msg_type":   msgType,
		"content":    content,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var result struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data any    `json:"data"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		g.Log().Errorf(ctx, "解析响应失败: %v, 响应内容: %s", err, string(body))
		return err
	}

	g.Log().Infof(ctx, "飞书 API 响应: code=%d, msg=%s", result.Code, result.Msg)

	if result.Code != 0 {
		return fmt.Errorf("发送消息失败: code=%d, msg=%s", result.Code, result.Msg)
	}

	g.Log().Infof(ctx, "成功发送消息到 %s", receiveID)
	return nil
}
