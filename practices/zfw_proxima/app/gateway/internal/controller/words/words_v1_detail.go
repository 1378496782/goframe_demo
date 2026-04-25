package words

import (
	"context"

	words "zfw_proxima/app/word/api/words/v1"

	"github.com/gogf/gf/v2/errors/gerror"

	v1 "zfw_proxima/app/gateway/api/words/v1"
)

func (c *ControllerV1) Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error) {
	word, err := c.WordsClient.Get(ctx, &words.GetReq{
		Id: uint32(req.Id),
	})
	if err != nil {
		return nil, err
	}
	if word == nil {
		return nil, gerror.New("word not found")
	}
	return &v1.DetailRes{
		Id:         uint32(word.Words.Id),
		Word:       word.Words.Word,
		Definition: word.Words.Definition,
	}, nil
}
