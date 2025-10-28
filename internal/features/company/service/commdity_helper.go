package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/icdb37/bfsm/internal/features/company/model"
	"github.com/icdb37/bfsm/internal/infra/cfpx"
)

func processCommodityCreate(_ context.Context, info *model.EntireCommodity) error {
	if err := cfpx.Process(info); err != nil {
		return err
	}
	info.CreatedAt = info.UpdatedAt
	info.ID = uuid.NewString()
	return nil
}

func processCommodityUpdate(_ context.Context, info *model.EntireCommodity) error {
	if err := cfpx.Process(info); err != nil {
		return err
	}
	return nil
}
