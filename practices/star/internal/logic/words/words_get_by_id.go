package words

import (
	"context"
	"star/internal/dao"
	"star/internal/model/entity"
)

func (w *Words) GetWordByID(ctx context.Context, id int) (res *entity.Words, err error) {
	res = &entity.Words{}
	err = dao.Words.Ctx(ctx).WherePri(id).Scan(&res)
	return
}
