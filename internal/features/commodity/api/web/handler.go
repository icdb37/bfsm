package web

import (
	"github.com/labstack/echo/v4"

	"github.com/icdb37/bfsm/internal/constx/featc"
	"github.com/icdb37/bfsm/internal/features/commodity/service"
	"github.com/icdb37/bfsm/internal/wire"
)

func Wire() {
	s1 := wire.ResolveName[service.TemplateServer](featc.CommodityTemplate)
	s2 := wire.ResolveName[service.CommodityServer](featc.CommodityCommodity)
	e := wire.Resolve[*echo.Echo]()
	registTemplate(e, s1)
	registCommodity(e, s2)
}

func registTemplate(e *echo.Echo, s service.TemplateServer) {
	u := &templateHandler{s: s}
	g := e.Group("/api/v1/commodity/template")
	{
		g.POST("/search", u.search)
		g.GET("/:id", u.get)
		g.POST("", u.create)
		g.PUT("/:id", u.update)
		g.DELETE("/:id", u.delete)
	}
}

func registCommodity(e *echo.Echo, s service.CommodityServer) {
	u := &commodityHandler{s: s}
	g := e.Group("/api/v1/commodity/commodity")
	{
		g.POST("/search", u.search)
		g.GET("/:id", u.get)
		g.POST("", u.create)
		g.PUT("/:id", u.update)
		g.DELETE("/:id", u.delete)
	}
}
