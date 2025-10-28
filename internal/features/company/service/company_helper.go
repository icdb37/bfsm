package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/icdb37/bfsm/internal/features/company/model"
	"github.com/icdb37/bfsm/internal/infra/cfpx"
)

func processCompanyCreate(_ context.Context, info *model.EntireCompany) error {
	if err := cfpx.Process(info); err != nil {
		return err
	}
	info.CreatedAt = info.UpdatedAt
	info.ID = uuid.NewString()
	return nil
}

func processCompanyUpdate(_ context.Context, info *model.EntireCompany) error {
	if err := cfpx.Process(info); err != nil {
		return err
	}
	return nil
}
