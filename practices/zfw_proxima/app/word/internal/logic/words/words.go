package words

import (
	"context"
	"zfw_proxima/app/word/internal/model/entity"

	"github.com/gogf/gf/v2/os/gtime"
)

func Create(ctx context.Context) (id int, err error) {
	return 1, nil
}

func Get(ctx context.Context) (words *entity.Words, err error) {
	return &entity.Words{
		Id:                 1,
		Uid:                1,
		Word:               "hello",
		Definition:         "used as a greeting when you meet somebody.",
		ExampleSentence:    "Hello, I am oldme!",
		ChineseTranslation: "你好",
		Pronunciation:      "həˈləʊ",
		CreatedAt:          gtime.New("2024-12-05 22:00:00"),
		UpdatedAt:          gtime.New("2024-12-05 22:00:00"),
	}, nil
}
