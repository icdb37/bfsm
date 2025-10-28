package echox

import (
	"context"

	"github.com/icdb37/bfsm/internal/infra/config"
	"github.com/icdb37/bfsm/internal/wire"
	"github.com/labstack/echo/v4"
)

type service struct {
	e *echo.Echo
}

func (s *service) Start(_ context.Context) error {
	endpoint := config.GetEndpoint()
	go s.e.Start(endpoint)
	return nil
}
func (s *service) Stop(_ context.Context) error {
	return s.e.Close()
}

func Provide() {
	e := echo.New()
	wire.Provide(func() *echo.Echo {
		return e
	})
	wire.Provide(func() wire.Starter {
		return &service{e: e}
	})
}
