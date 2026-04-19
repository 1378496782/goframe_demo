package words

import (
	"context"

	v1 "star/api/words/v1"
	"star/internal/logic/words"
)

func (c *ControllerV1) CreateWord(ctx context.Context, req *v1.CreateWordReq) (res *v1.CreateWordRes, err error) {
	uid, err := c.Users.GetUid(ctx)
	if err != nil {
		return nil, err
	}
	err = c.Words.CreateWord(ctx, &words.CreateInput{
		Uid:                uid,
		Word:               req.Word,
		Definition:         req.Definition,
		ExampleSentence:    req.ExampleSentence,
		ChineseTranslation: req.ChineseTranslation,
		Pronunciation:      req.Pronunciation,
		ProficiencyLevel:   req.ProficiencyLevel,
	})
	return nil, err
}
