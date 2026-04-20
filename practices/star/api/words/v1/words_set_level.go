package v1

import "github.com/gogf/gf/v2/frame/g"

type SetLevelReq struct {
	g.Meta `path:"/words/{id}/level" method:"patch" sm:"设置单词等级" tags:"单词"`
	Id     int              `json:"id" v:"required"`
	Level  ProficiencyLevel `json:"level" v:"required|between:1,5"`
}

type SetLevelRes struct {
}
