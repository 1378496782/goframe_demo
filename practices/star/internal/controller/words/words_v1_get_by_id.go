package words

import (
	"context"

	v1 "star/api/words/v1"
)

func (c *ControllerV1) GetByID(ctx context.Context, req *v1.GetByIDReq) (res *v1.GetByIDRes, err error) {
	id := req.Id
	word, err := c.Words.GetWordByID(ctx, id)
	if err != nil {
		return
	}
	res = &v1.GetByIDRes{
		Word: word,
	}
	return
}
