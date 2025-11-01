package service

import (
	"context"

	"github.com/icdb37/bfsm/internal/infra/store"
	coModel "github.com/icdb37/bfsm/internal/model"
)

type consumeImpl struct {
	repo store.Tabler
}

func (c *consumeImpl) Search(ctx context.Context, req *coModel.SearchRequest[coModel.QueryBatch]) (resp *coModel.SearchResponse[coModel.EntireBatch], err error) {
	return nil, nil
}

func (c *consumeImpl) Get(ctx context.Context, id string) (resp []*coModel.EntireBatch, err error) {
	return nil, nil
}
func (c *consumeImpl) Create(ctx context.Context, info *coModel.EntireBatch) error {
	return nil
}
func (c *consumeImpl) Update(ctx context.Context, info *coModel.EntireBatch) error {
	return nil
}

func (c *consumeImpl) Delete(ctx context.Context, id string) error {
	return nil
}
