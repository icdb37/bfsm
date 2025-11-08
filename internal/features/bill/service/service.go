package service

import (
	"github.com/icdb37/bfsm/internal/constx/featc"
	"github.com/icdb37/bfsm/internal/features/bill/model"
	"github.com/icdb37/bfsm/internal/infra/logx"
	"github.com/icdb37/bfsm/internal/infra/store"
	coService "github.com/icdb37/bfsm/internal/service"
	"github.com/icdb37/bfsm/internal/wire"
)

func Provide() {
	wire.ProvideName(featc.BillBatch, func() coService.BillSaver {
		repo, err := store.NewTable(&model.BillBatch{})
		if err != nil {
			logx.Fatal("create bill batch repo failed", "error", err)
		}
		return &batchImpl{repoBatch: repo}
	})
}
