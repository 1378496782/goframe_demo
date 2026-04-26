package user

import (
	"context"
	v1 "my-user-grpc-service/app/user/api/user/v1"
	"my-user-grpc-service/app/user/internal/logic/user"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
)

type Controller struct {
	v1.UnimplementedUserServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterUserServer(s.Server, &Controller{})
}

func (*Controller) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	users, err := user.GetList(ctx)
	return &v1.GetListRes{
		Users: users,
	}, err
}

func (*Controller) GetOne(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error) {
	user, err := user.GetById(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &v1.GetOneRes{
		User: user,
	}, nil
}

func (*Controller) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	err = user.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return &v1.CreateRes{}, nil
}

func (*Controller) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	err = user.Delete(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &v1.DeleteRes{}, nil
}
