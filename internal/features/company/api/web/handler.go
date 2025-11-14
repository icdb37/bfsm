package web

import (
	"github.com/labstack/echo/v4"

	"github.com/icdb37/bfsm/internal/constx/featc"
	"github.com/icdb37/bfsm/internal/features/company/service"
	"github.com/icdb37/bfsm/internal/wire"
)

func Wire() {
	s1 := wire.ResolveName[service.CompanyServer](featc.CompanyCompany)
	s2 := wire.ResolveName[service.CommodityServer](featc.CompanyCommodity)
	e := wire.Resolve[*echo.Echo]()
	registCompany(e, s1)
	registCommodity(e, s2)
}

func registCompany(e *echo.Echo, s service.CompanyServer) {
	u := &companyHandler{s: s}
	g := e.Group("/api/v1/company")
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
	g := e.Group("/api/v1/company/:company_id/commodity")
	{
		g.POST("/search", u.search)
		g.GET("/:id", u.get)
		g.POST("", u.create)
		g.PUT("/:id", u.update)
		g.DELETE("/:id", u.delete)
	}
}
