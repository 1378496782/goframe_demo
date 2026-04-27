// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure of table user for DAO operations like Where/Data.
type User struct {
	g.Meta   `orm:"table:user, do:true"`
	Id       any         // User ID
	Passport any         // User Passport
	Password any         // User Password
	Nickname any         // User Nickname
	CreateAt *gtime.Time // Created Time
	UpdateAt *gtime.Time // Updated Time
	DeleteAt *gtime.Time // Deleted Time
}
