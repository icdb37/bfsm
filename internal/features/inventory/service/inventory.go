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
	repoLast store.Tabler
	repoFull store.Tabler
}

func (i *inventoryImpl) Search(ctx context.Context, req *coModel.SearchRequest[model.QueryCommodity]) (resp *coModel.SearchResponse[model.LastCommodity], err error) {
	qf := store.Unmarshal(req.Query)
	resp = &coModel.SearchResponse[model.LastCommodity]{}
	pf := req.GetPage()
	if resp.Total, err = i.repoLast.Search(ctx, qf, pf, &(resp.Datas)); err != nil {
		logx.Error("search last commodity failed", "error", err)
		return nil, err
	}
	return resp, nil
}

func (i *inventoryImpl) Get(ctx context.Context, id string) (resp []*model.FullCommodity, err error) {
	resp = []*model.FullCommodity{}
	if err := i.repoFull.Query(ctx, store.NewFilter().Eq(field.ID, id), &resp); err != nil {
		logx.Error("get full commodity failed", "id", id, "error", err)
		return nil, err
	}
	return resp, nil
}

// Produce 增加库存
func (i *inventoryImpl) Produce(ctx context.Context, info *coModel.EntireBatch) error {
	return i.save(ctx, info)
}

// Consume 减少库存
func (i *inventoryImpl) Consume(ctx context.Context, info *coModel.EntireBatch) error {
	return i.save(ctx, info)
}
