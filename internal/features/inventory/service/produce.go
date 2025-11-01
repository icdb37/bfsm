package service

import (
	"context"

	"github.com/icdb37/bfsm/internal/infra/store"
	coModel "github.com/icdb37/bfsm/internal/model"
)

type produceImpl struct {
	repo store.Tabler
}

func (c *produceImpl) Search(ctx context.Context, req *coModel.SearchRequest[coModel.QueryBatch]) (resp *coModel.SearchResponse[coModel.EntireBatch], err error) {
	return nil, nil
}

func (c *produceImpl) Get(ctx context.Context, id string) (resp []*coModel.EntireBatch, err error) {
	return nil, nil
}
func (c *produceImpl) Create(ctx context.Context, info *coModel.EntireBatch) error {
	return nil
}
func (c *produceImpl) Update(ctx context.Context, info *coModel.EntireBatch) error {
	return nil
}

func (c *produceImpl) Delete(ctx context.Context, id string) error {
	return nil
}
