package users

import (
	"context"
	"star/internal/consts"
	"star/internal/dao"
	"star/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/golang-jwt/jwt/v5"

	"github.com/gogf/gf/v2/errors/gerror"
)

func (u *Users) Info(ctx context.Context) (user *entity.Users, err error) {
	tokenString := g.RequestFromCtx(ctx).Request.Header.Get("Authorization")
	if tokenString == "" {
		return nil, gerror.New("token is empty")
	}
	tokenClaims, _ := jwt.ParseWithClaims(tokenString, &jwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(consts.JwtKey), nil
	})
	if claims, ok := tokenClaims.Claims.(*jwtClaims); ok && tokenClaims.Valid {
		err = dao.Users.Ctx(ctx).WherePri(claims.Id).Scan(&user)
	}
	return
}
