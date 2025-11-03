package web

import (
	"github.com/labstack/echo/v4"

	"github.com/icdb37/bfsm/internal/constx/featc"
	"github.com/icdb37/bfsm/internal/features/inventory/service"
	coService "github.com/icdb37/bfsm/internal/service"
	"github.com/icdb37/bfsm/internal/wire"
)

func Wire() {
	s := wire.ResolveName[service.InventoryInventory](featc.InventoryInventory)
	p := wire.ResolveName[coService.InventoryProducer](featc.InventoryProduce)
	c := wire.ResolveName[coService.InventoryConsumer](featc.InventoryConsume)
	e := wire.Resolve[*echo.Echo]()
	u := &inventoryHandler{s: s, p: p, c: c}
	g := e.Group("/api/v1/inventory/inventory")
	{
		g.POST("/search", u.search)
		g.GET("/:id", u.get)
	}
	e.POST("/api/v1/inventory/produce", u.produce)
	e.POST("/api/v1/inventory/consume", u.consume)
}
