package words

import (
	"context"

	v1 "star/api/words/v1"
	"star/internal/logic/words"
)

func (c *ControllerV1) wordsCreateBatchReq2CreateBatchInput(ctx context.Context, req *v1.WordsCreateBatchReq) (in words.CreateBatchInput, err error) {
	uid, err := c.Users.GetUid(ctx)
	if err != nil {
		return words.CreateBatchInput{}, err
	}
	in = words.CreateBatchInput{
		Words: make([]words.CreateInput, 0, len(req.Words)),
	}
	for _, word := range req.Words {
		in.Words = append(in.Words, words.CreateInput{
			Uid:                uid,
			ProficiencyLevel:   word.ProficiencyLevel,
			Word:               word.Word,
			Definition:         word.Definition,
			ExampleSentence:    word.ExampleSentence,
			ChineseTranslation: word.ChineseTranslation,
			Pronunciation:      word.Pronunciation,
		})
	}
	return
}

func (c *ControllerV1) WordsCreateBatch(ctx context.Context, req *v1.WordsCreateBatchReq) (res *v1.WordsCreateBatchRes, err error) {
	in, err := c.wordsCreateBatchReq2CreateBatchInput(ctx, req)
	if err != nil {
		return nil, err
	}
	err = c.Words.CreateWordsBatch(ctx, &in)
	if err != nil {
		return nil, err
	}
	return
}
