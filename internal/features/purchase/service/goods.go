package service

import (
	"context"

	"github.com/icdb37/bfsm/internal/features/purchase/model"
	"github.com/icdb37/bfsm/internal/infra/logx"
	"github.com/icdb37/bfsm/internal/infra/store"
	coModel "github.com/icdb37/bfsm/internal/model"
)

type goodsImpl struct {
	repo store.Tabler
}

// Search 搜索商品
func (g *goodsImpl) Search(ctx context.Context, req *coModel.SearchRequest[model.QueryGoods]) (resp *coModel.SearchResponse[model.PurchaseGoods], err error) {
	qf := store.Unmarshal(req.Query)
	resp = &coModel.SearchResponse[model.PurchaseGoods]{}
	pf := req.GetPage()
	if resp.Total, err = g.repo.Search(ctx, qf, pf, &(resp.Datas)); err != nil {
		logx.Error("search goods failed", "error", err)
		return nil, err
	}
	return resp, nil
}
