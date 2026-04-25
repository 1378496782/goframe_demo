package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type CreateReq struct {
	g.Meta     `path:"/words/create" method:"post" sm:"创建单词" tags:"单词"`
	Word       string `json:"word" v:"required|length:1,100" dc:"单词"`
	Definition string `json:"definition" v:"required|length:1,300" dc:"单词定义"`
}

type CreateRes struct {
}

type DetailReq struct {
	g.Meta `path:"/words/{id}" method:"get" sm:"单词详情" tags:"单词"`
	Id     uint32 `json:"id" v:"required" dc:"单词ID"`
}

type DetailRes struct {
	Id                 uint32      `json:"id" dc:"单词ID"`
	Word               string      `json:"word" dc:"单词"`
	Definition         string      `json:"definition" dc:"单词定义"`
	ExampleSentence    string      `json:"exampleSentence" dc:"例句"`
	ChineseTranslation string      `json:"chineseTranslation" dc:"中文翻译"`
	Pronunciation      string      `json:"pronunciation" dc:"发音"`
	CreatedAt          *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt          *gtime.Time `json:"updatedAt" dc:"更新时间"`
}
