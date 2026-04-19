package words

import (
	"context"
	v1 "star/api/words/v1"
	"star/internal/dao"
	"star/internal/model/do"

	"github.com/gogf/gf/v2/errors/gerror"
)

type UpdateInput struct {
	Id                 uint                `json:"id" v:"required"`
	Uid                uint                `json:"uid" v:"required"`
	Word               string              `json:"word" v:"required|length:1,100" dc:"单词"`
	Definition         string              `json:"definition" v:"required|length:1,300" dc:"单词定义"`
	ExampleSentence    string              `json:"example_sentence" v:"required|length:1,300" dc:"例句"`
	ChineseTranslation string              `json:"chinese_translation" v:"required|length:1,300" dc:"中文翻译"`
	Pronunciation      string              `json:"pronunciation" v:"required|length:1,100" dc:"发音"`
	ProficiencyLevel   v1.ProficiencyLevel `json:"proficiency_level" v:"required|between:1,5" dc:"熟练度，1最低，5最高"`
}

func (w *Words) UpdateWord(ctx context.Context, in UpdateInput) error {
	var cls = dao.Words.Columns()
	count, err := dao.Words.Ctx(ctx).
		Where(cls.Uid, in.Uid).
		Where(cls.Word, in.Word).
		WhereNot(cls.Id, in.Id).Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return gerror.New("单词已存在")
	}

	_, err = dao.Words.Ctx(ctx).Data(do.Words{
		Word:               in.Word,
		Definition:         in.Definition,
		ExampleSentence:    in.ExampleSentence,
		ChineseTranslation: in.ChineseTranslation,
		Pronunciation:      in.Pronunciation,
		ProficiencyLevel:   in.ProficiencyLevel,
	}).Where(cls.Id, in.Id).Where(cls.Uid, in.Uid).Update()
	return err
}
