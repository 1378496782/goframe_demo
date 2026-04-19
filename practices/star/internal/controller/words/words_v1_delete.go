package words

import (
	"context"

	v1 "star/api/words/v1"
)

func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	uid, err := c.Users.GetUid(ctx)
	if err != nil {
		return nil, err
	}
	err = c.Words.DeleteWord(ctx, uid, req.Id)
	return
}
