package users

import (
	"context"
	v1 "star/api/users/v1"
	"star/internal/consts"
	"star/internal/dao"
	"star/internal/model/entity"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/gogf/gf/v2/errors/gerror"
)

type jwtClaims struct {
	Id       uint
	Username string
	jwt.RegisteredClaims
}

func (u *Users) Login(ctx context.Context, req *v1.LoginReq) (token string, err error) {
	var user entity.Users
	err = dao.Users.Ctx(ctx).Where("username", req.Username).Scan(&user)
	if err != nil {
		return "", gerror.New("用户名或密码错误")
	}
	if user.Id == 0 {
		return "", gerror.New("用户不存在")
	}
	if user.Password != u.encryptPassword(req.Password) {
		return "", gerror.New("密码错误")
	}

	// 生成 token
	uc := &jwtClaims{
		Id:       user.Id,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(6 * time.Hour)),
		},
	}
	token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, uc).SignedString([]byte(consts.JwtKey))
	if err != nil {
		return "", gerror.New("生成 token失败")
	}

	return token, nil
}
