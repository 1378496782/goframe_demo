package words

import (
	"context"

	v1 "star/api/words/v1"
	"star/internal/logic/words"
)

func (c *ControllerV1) WordsList(ctx context.Context, req *v1.WordsListReq) (res *v1.WordsListRes, err error) {
	output, err := c.Words.GetWordsList(ctx, &words.ListInput{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}
	res = &v1.WordsListRes{
		Words: output.Words,
		Total: uint(output.Total),
	}
	return
}
