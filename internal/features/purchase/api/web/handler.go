package web

import (
	"github.com/labstack/echo/v4"

	"github.com/icdb37/bfsm/internal/constx/featc"
	"github.com/icdb37/bfsm/internal/features/purchase/service"
	"github.com/icdb37/bfsm/internal/wire"
)

func Wire() {
	s1 := wire.ResolveName[service.BatchServer](featc.PurchaseBatch)
	s2 := wire.ResolveName[service.GoodsServer](featc.PurchaseGoods)
	e := wire.Resolve[*echo.Echo]()
	registBatch(e, s1)
	registGoods(e, s2)
}

func registBatch(e *echo.Echo, s service.BatchServer) {
	u := &batchHandler{s: s}
	g := e.Group("/api/v1/purchase/batch")
	{
		g.POST("/search", u.search)
		g.GET("/:id", u.get)
		g.POST("", u.create)
		g.PUT("/:id", u.update)
		g.PATCH("/:id/status", u.updateStatus)
		g.DELETE("/:id", u.delete)
	}
}

func registGoods(e *echo.Echo, s service.GoodsServer) {
	u := &goodsHandler{s: s}
	g := e.Group("/api/v1/purchase/goods")
	{
		g.POST("/search", u.search)
		g.GET("/:id", u.get)
	}
}
