package words

import (
	"context"

	v1 "star/api/words/v1"
	"star/internal/dao"
	"star/internal/model/entity"

	"github.com/gogf/gf/v2/errors/gerror"
)

type SetLevelInput struct {
	ID    int
	UID   int
	Level v1.ProficiencyLevel
}

func (w *Words) SetWordLevel(ctx context.Context, in *SetLevelInput) error {
	if in.Level < v1.ProficiencyLevel1 || in.Level > v1.ProficiencyLevel5 {
		return gerror.New("熟练度值不合法")
	}
	var (
		cls = dao.Words.Columns()
		orm = dao.Words.Ctx(ctx)
	)
	// 检查单词是否存在且属于该用户
	word := &entity.Words{}
	err := orm.Where(cls.Id, in.ID).Where(cls.Uid, in.UID).Scan(&word)
	if err != nil || word.Id == 0 {
		return gerror.New("单词不存在或无权限设置等级")
	}

	// 更新单词等级
	orm = orm.Where(cls.Uid, in.UID).WherePri(cls.Id, in.ID).Data(cls.ProficiencyLevel, in.Level)
	_, err = orm.Update()
	return err
}
