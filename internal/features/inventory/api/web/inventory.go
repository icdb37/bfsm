package web

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/icdb37/bfsm/internal/constx/field"
	"github.com/icdb37/bfsm/internal/features/inventory/model"
	"github.com/icdb37/bfsm/internal/features/inventory/service"
	"github.com/icdb37/bfsm/internal/infra/logx"
	coModel "github.com/icdb37/bfsm/internal/model"
	coService "github.com/icdb37/bfsm/internal/service"
)

// 商品服务接口
type inventoryHandler struct {
	s service.InventoryInventory
	p coService.InventoryProducer
	c coService.InventoryConsumer
}

func (u *inventoryHandler) search(c echo.Context) error {
	req := &coModel.SearchRequest[model.QueryCommodity]{}
	ctx := c.Request().Context()
	if err := c.Bind(req); err != nil {
		logx.Error("search commodity bind failed", "error", err)
		return c.JSON(http.StatusBadRequest, err)
	}
	resp, err := u.s.Search(ctx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, resp)
}

func (u *inventoryHandler) get(c echo.Context) error {
	id := c.Param(field.ID)
	ctx := c.Request().Context()
	info, err := u.s.Get(ctx, id)
	if err != nil {
		logx.Error("get commodity failed", "error", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, info)
}

func (u *inventoryHandler) produce(c echo.Context) error {
	info := &coModel.EntireBatch{}
	ctx := c.Request().Context()
	if err := c.Bind(info); err != nil {
		logx.Error("produce commodity bind failed", "error", err)
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := u.p.Produce(ctx, info); err != nil {
		logx.Error("produce commodity failed", "error", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	logx.Info("produce commodity success", field.ID, info.ID)
	return c.JSON(http.StatusOK, coModel.NewIDResponse(info.ID))
}

func (u *inventoryHandler) consume(c echo.Context) error {
	info := &coModel.EntireBatch{}
	ctx := c.Request().Context()
	if err := c.Bind(info); err != nil {
		logx.Error("consume commodity bind failed", "error", err)
		return c.JSON(http.StatusBadRequest, err)
	}
	info.ID = c.Param(field.ID)
	if err := u.c.Consume(ctx, info); err != nil {
		logx.Error("consume commodity failed", "error", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	logx.Info("consume commodity success", field.ID, info.ID)
	return c.JSON(http.StatusOK, coModel.NewIDResponse(info.ID))
}
