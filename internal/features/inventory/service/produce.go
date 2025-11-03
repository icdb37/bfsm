package service

import (
	"context"

	"github.com/icdb37/bfsm/internal/infra/store"
	coModel "github.com/icdb37/bfsm/internal/model"
)

type produceImpl struct {
	repo store.Tabler
}

func (p *produceImpl) Search(ctx context.Context, req *coModel.SearchRequest[coModel.QueryBatch]) (resp *coModel.SearchResponse[coModel.EntireBatch], err error) {
	return nil, nil
}

func (p *produceImpl) Get(ctx context.Context, id string) (resp []*coModel.EntireBatch, err error) {
	return nil, nil
}
func (p *produceImpl) Create(ctx context.Context, info *coModel.EntireBatch) error {
	return nil
}
func (p *produceImpl) Update(ctx context.Context, info *coModel.EntireBatch) error {
	return nil
}

func (p *produceImpl) Delete(ctx context.Context, id string) error {
	return nil
}
