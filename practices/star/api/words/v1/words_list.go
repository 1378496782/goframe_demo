package v1

import (
	"star/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type WordsListReq struct {
	g.Meta `path:"/words" method:"get" sm:"分页获取单词列表" tags:"单词"`
	Page   int `json:"page" v:"min:1" dc:"页码，默认1"`
	Size   int `json:"size" v:"between:1,100" dc:"每页数量，默认10"`
}

type WordsListRes struct {
	Total uint           `json:"total"`
	Words []entity.Words `json:"words"`
}
