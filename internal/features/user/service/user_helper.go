package service

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"

	"github.com/icdb37/bfsm/internal/features/user/model"
)

func processCreate(_ context.Context, info *model.EntireUser) error {
	info.ID = uuid.NewString()
	info.CreatedAt = time.Now()
	info.UpdatedAt = info.CreatedAt
	return nil
}

func processCheck(_ context.Context, info *model.EntireUser) error {
	info.Normalize()
	if info.Name == "" {
		return errors.New("name is required")
	}
	return nil
}

func processUpdate(_ context.Context, info *model.EntireUser) error {
	info.UpdatedAt = time.Now()
	return nil
}
