package words

import (
	"context"

	"golang.org/x/sync/errgroup"
)

type CreateBatchInput struct {
	Words []CreateInput
}

// func (w *Words) CreateWordsBatch(ctx context.Context, in *CreateBatchInput) error {
// 	for _, word := range in.Words {
// 		err := w.CreateWord(ctx, &word)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

// CreateWordsBatch 并发创建单词批量
func (w *Words) CreateWordsBatch(ctx context.Context, in *CreateBatchInput) error {
	g, ctx := errgroup.WithContext(ctx)
	g.SetLimit(10) // 最多同时10个并发

	for _, word := range in.Words {
		word := word
		g.Go(func() error {
			return w.CreateWord(ctx, &word)
		})
	}

	return g.Wait()
}
