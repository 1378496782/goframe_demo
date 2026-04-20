package v1

import (
	"star/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type GetByIDReq struct {
	g.Meta `path:"/words/{id}" method:"GET" summary:"根据ID获取单词" tags:"单词"`
	Id     int `json:"id" v:"required" dc:"单词ID"`
}

type GetByIDRes struct {
	Word *entity.Words `json:"word" dc:"单词"`
}
