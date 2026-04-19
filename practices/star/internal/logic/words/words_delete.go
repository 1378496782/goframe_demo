package words

import (
	"context"
	"star/internal/dao"
	"star/internal/model/entity"

	"github.com/gogf/gf/v2/errors/gerror"
)

func (w *Words) DeleteWord(ctx context.Context, uid, id uint) error {
	var (
		cls = dao.Words.Columns()
		orm = dao.Words.Ctx(ctx)
	)

	// 检查单词是否存在且属于该用户
	var word entity.Words
	err := orm.Where(cls.Uid, uid).WherePri(id).Scan(&word)
	if err != nil {
		return gerror.New("单词不存在或无权限删除")
	}

	// 删除单词
	_, err = orm.Where(cls.Uid, uid).WherePri(id).Delete()
	return err
}
