package words

import (
	"context"
	"star/internal/dao"
	"star/internal/model/entity"
)

type ListInput struct {
	Page int `json:"page" v:"min:1" dc:"页码，默认1"`
	Size int `json:"size" v:"between:1,100" dc:"每页数量，默认10"`
}

type ListOutput struct {
	Words []entity.Words `json:"words"`
	Total int            `json:"total"`
}

func (w *Words) GetWordsList(ctx context.Context, in *ListInput) (output ListOutput, err error) {
	// 对于查询初始值的处理
	if in.Page == 0 {
		in.Page = 1
	}
	if in.Size == 0 {
		in.Size = 15
	}

	output.Words = []entity.Words{}
	err = dao.Words.Ctx(ctx).Page(in.Page, in.Size).ScanAndCount(&output.Words, &output.Total, true)
	return
}
