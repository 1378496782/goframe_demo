// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package words

import (
	"context"

	"star/api/words/v1"
)

type IWordsV1 interface {
	CreateWord(ctx context.Context, req *v1.CreateWordReq) (res *v1.CreateWordRes, err error)
	WordsCreateBatch(ctx context.Context, req *v1.WordsCreateBatchReq) (res *v1.WordsCreateBatchRes, err error)
	Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error)
	WordsList(ctx context.Context, req *v1.WordsListReq) (res *v1.WordsListRes, err error)
	UpdateWord(ctx context.Context, req *v1.UpdateWordReq) (res *v1.UpdateWordRes, err error)
}
