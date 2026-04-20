package words

import (
	"context"

	v1 "star/api/words/v1"
	"star/internal/logic/words"

	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *ControllerV1) SetLevel(ctx context.Context, req *v1.SetLevelReq) (res *v1.SetLevelRes, err error) {
	uid, err := c.Users.GetUid(ctx)
	if err != nil {
		return
	}
	_, err = c.Words.GetWordByID(ctx, req.Id)
	if err != nil {
		return nil, gerror.New("单词不存在")
	}
	err = c.Words.SetWordLevel(ctx, &words.SetLevelInput{
		ID:    req.Id,
		UID:   int(uid),
		Level: req.Level,
	})
	return
}
