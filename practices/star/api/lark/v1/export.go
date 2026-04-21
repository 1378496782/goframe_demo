package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type ExportWordsReq struct {
	g.Meta           `path:"/lark/words/export" method:"post" tags:"飞书集成" summary:"导出单词到飞书"`
	SpreadsheetToken string `json:"spreadsheetToken" dc:"电子表格Token,不传则使用默认"`
}

type ExportWordsRes struct {
	Count          int    `json:"count" dc:"导出的单词数量"`
	SpreadsheetUrl string `json:"spreadsheetUrl" dc:"电子表格链接"`
	Message        string `json:"message"`
}

type SendReminderReq struct {
	g.Meta    `path:"/lark/reminder/send" method:"post" tags:"飞书集成" summary:"发送学习提醒"`
	ReceiveID string `json:"receiveId" dc:"接收者ID,不传则查询用户的飞书ID"`
}

type SendReminderRes struct {
	Message string `json:"message"`
}
