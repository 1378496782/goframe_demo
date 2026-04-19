package v1

import "github.com/gogf/gf/v2/frame/g"

type DeleteReq struct {
	g.Meta `path:"/words/{id}" method:"delete" sm:"删除单词" tags:"单词"`
	Id     uint `json:"id" v:"required" dc:"单词ID"`
}

type DeleteRes struct {
}
