package lark

import (
	"context"

	v1 "star/api/lark/v1"
)

func (c *ControllerV1) ExportWords(ctx context.Context, req *v1.ExportWordsReq) (res *v1.ExportWordsRes, err error) {
	return c.Lark.ExportWords(ctx, req)
}

func (c *ControllerV1) SendReminder(ctx context.Context, req *v1.SendReminderReq) (res *v1.SendReminderRes, err error) {
	return c.Lark.SendReminder(ctx, req)
}
