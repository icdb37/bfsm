package web

import (
	"github.com/labstack/echo/v4"

	"github.com/icdb37/bfsm/internal/constx/featc"
	"github.com/icdb37/bfsm/internal/features/bill/service"
	"github.com/icdb37/bfsm/internal/wire"
)

func Wire() {
	s1 := wire.ResolveName[service.BatchServer](featc.BillBatch)
	s2 := wire.ResolveName[service.DealServer](featc.BillDeal)
	e := wire.Resolve[*echo.Echo]()
	registBatch(e, s1)
	registGoods(e, s2)
}

func registBatch(e *echo.Echo, s service.BatchServer) {
	u := &batchHandler{s: s}
	g := e.Group("/api/v1/bill/batch")
	{
		g.POST("/search", u.search)
		g.GET("/:id", u.get)
		g.POST("", u.create)
		g.PUT("/:id", u.update)
		g.PATCH("/:id/status", u.updateStatus)
		g.PATCH("/:id/amount", u.updateAmount)
	}
}

func registGoods(e *echo.Echo, s service.DealServer) {
	u := &dealHandler{s: s}
	g := e.Group("/api/v1/bill/deal")
	{
		g.POST("/search", u.search)
		g.PUT("/:id", u.update)
	}
}
