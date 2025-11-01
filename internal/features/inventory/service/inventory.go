package service

import (
	"context"

	"github.com/icdb37/bfsm/internal/constx/field"
	"github.com/icdb37/bfsm/internal/features/inventory/model"
	"github.com/icdb37/bfsm/internal/infra/logx"
	"github.com/icdb37/bfsm/internal/infra/store"
	coModel "github.com/icdb37/bfsm/internal/model"
)

type inventoryImpl struct {
	repo store.Tabler
}

func (i *inventoryImpl) Search(ctx context.Context, req *coModel.SearchRequest[model.QueryCommodity]) (resp *coModel.SearchResponse[model.EntireCommodity], err error) {
	return nil, nil
}

func (i *inventoryImpl) Get(ctx context.Context, id string) (resp []*model.EntireCommodity, err error) {
	resp = []*model.EntireCommodity{}
	if err := i.repo.Query(ctx, store.NewFilter().Eq(field.ID, id), &resp); err != nil {
		logx.Error("get commodity failed", "id", id, "error", err)
		return nil, err
	}
	return resp, nil
}
