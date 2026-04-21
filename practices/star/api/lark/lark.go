// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package lark

import (
	"context"

	"star/api/lark/v1"
)

type ILarkV1 interface {
	ExportWords(ctx context.Context, req *v1.ExportWordsReq) (res *v1.ExportWordsRes, err error)
	SendReminder(ctx context.Context, req *v1.SendReminderReq) (res *v1.SendReminderRes, err error)
}
