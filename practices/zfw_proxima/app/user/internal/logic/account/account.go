package account

import (
	"context"
	"zfw_proxima/app/user/internal/model/entity"

	"github.com/gogf/gf/v2/os/gtime"
)

func Register(ctx context.Context) (id int, err error) {
	return 1, nil
}

func Login(ctx context.Context) (token string, err error) {
	return "I'm token", nil
}

func Info(ctx context.Context, token string) (user *entity.Users, err error) {
	return &entity.Users{
		Id:        1,
		Username:  "zfw",
		Password:  "123456",
		Email:     "zfw@bytedance.com",
		CreatedAt: gtime.New("2024-12-05 22:00:00"),
		UpdatedAt: gtime.New("2024-12-05 22:00:00"),
	}, nil
}
