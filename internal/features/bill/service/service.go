package service

import (
	"context"

	"github.com/icdb37/bfsm/internal/constx/featc"
	"github.com/icdb37/bfsm/internal/features/bill/model"
	"github.com/icdb37/bfsm/internal/infra/logx"
	"github.com/icdb37/bfsm/internal/infra/store"
	coModel "github.com/icdb37/bfsm/internal/model"
	coService "github.com/icdb37/bfsm/internal/service"
	"github.com/icdb37/bfsm/internal/wire"
)

type DealServer interface {
	Search(ctx context.Context, req *coModel.SearchRequest[model.QueryBillDeal]) (resp *coModel.SearchResponse[model.BillDeal], err error)
	Update(ctx context.Context, newInfo *model.BillDeal) error
}

type BatchServer interface {
	Search(ctx context.Context, req *coModel.SearchRequest[model.QueryBillBatch]) (resp *coModel.SearchResponse[model.SimpleBillBatch], err error)
	Get(ctx context.Context, id string) (info *model.BillBatch, err error)
	Create(ctx context.Context, info *model.BillBatch) error
	Update(ctx context.Context, newInfo *model.BillBatch) error
	SyncAmount(ctx context.Context, id string) error
	UpdateAmount(ctx context.Context, info *coModel.UpdateAmount) error
	Approve(ctx context.Context, param *coModel.UpdateStatus) error
	Complete(ctx context.Context, param *coModel.UpdateStatus) error
	Cancel(ctx context.Context, param *coModel.UpdateStatus) error
	Close(ctx context.Context, param *coModel.UpdateStatus) error
}

func Provide() {
	repoBatch, err := store.NewTable(&model.BillBatch{})
	if err != nil {
		logx.Fatal("create bill batch repo failed", "error", err)
	}
	repoDeal, err := store.NewTable(&model.BillDeal{})
	if err != nil {
		logx.Fatal("create bill deal repo failed", "error", err)
	}
	bs := &batchImpl{repoBatch: repoBatch, repoDeal: repoDeal}
	wire.ProvideName(featc.BillSave, func() coService.BillSaver {
		return bs
	})
	wire.ProvideName(featc.BillBatch, func() BatchServer {
		return bs
	})
	wire.ProvideName(featc.BillDeal, func() DealServer {
		return &dealImpl{repoDeal: repoDeal}
	})
}
