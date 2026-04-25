package words

import (
	"context"
	"zfw_proxima/app/word/api/pbentity"
	v1 "zfw_proxima/app/word/api/words/v1"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"google.golang.org/protobuf/types/known/timestamppb"

	"zfw_proxima/app/word/internal/logic/words"
)

type Controller struct {
	v1.UnimplementedWordsServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterWordsServer(s.Server, &Controller{})
}

func (*Controller) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	id, err := words.Create(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.CreateRes{Id: uint32(id)}, nil
}

func (*Controller) Get(ctx context.Context, req *v1.GetReq) (res *v1.GetRes, err error) {
	words, err := words.Get(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.GetRes{
		Words: &pbentity.Words{
			Id:                 uint32(words.Id),
			Uid:                uint32(words.Uid),
			Word:               words.Word,
			Definition:         words.Definition,
			ExampleSentence:    words.ExampleSentence,
			ChineseTranslation: words.ChineseTranslation,
			Pronunciation:      words.Pronunciation,
			CreatedAt:          timestamppb.New(words.CreatedAt.Time),
			UpdatedAt:          timestamppb.New(words.UpdatedAt.Time),
		},
	}, nil
}
