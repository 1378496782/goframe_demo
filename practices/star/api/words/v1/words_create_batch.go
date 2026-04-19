package v1

import "github.com/gogf/gf/v2/frame/g"

type WordsCreateBatchReq struct {
	g.Meta `path:"/word/create/batch" method:"post" sm:"批量创建单词" tags:"单词"`
	Words  []CreateWordReq `json:"words" v:"required"`
}

type WordsCreateBatchRes struct {
}
