package words

import (
	"context"

	v1 "star/api/words/v1"
	"star/internal/logic/words"
)

func (c *ControllerV1) UpdateWord(ctx context.Context, req *v1.UpdateWordReq) (res *v1.UpdateWordRes, err error) {
	uid, err := c.Users.GetUid(ctx)
	if err != nil {
		return nil, err
	}

	err = c.Words.UpdateWord(ctx, words.UpdateInput{
		Id:                 req.Id,
		Uid:                uid,
		Word:               req.Word,
		Definition:         req.Definition,
		ExampleSentence:    req.ExampleSentence,
		ChineseTranslation: req.ChineseTranslation,
		Pronunciation:      req.Pronunciation,
		ProficiencyLevel:   req.ProficiencyLevel,
	})
	return
}
