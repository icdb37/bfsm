package web

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/icdb37/bfsm/internal/constx/enum"
	"github.com/icdb37/bfsm/internal/constx/field"
	"github.com/icdb37/bfsm/internal/features/inventory/model"
	"github.com/icdb37/bfsm/internal/features/inventory/service"
	"github.com/icdb37/bfsm/internal/infra/logx"
	coModel "github.com/icdb37/bfsm/internal/model"
	coService "github.com/icdb37/bfsm/internal/service"
)

// 商品服务接口
type inventoryHandler struct {
	i service.InventoryInventory
	s coService.InventorySaver
}

func (h *inventoryHandler) searchLast(c echo.Context) error {
	req := &coModel.SearchRequest[model.QueryLastCommodity]{}
	ctx := c.Request().Context()
	if err := c.Bind(req); err != nil {
		logx.Error("search commodity bind failed", "error", err)
		return c.JSON(http.StatusBadRequest, err)
	}
	resp, err := h.i.SearchLast(ctx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *inventoryHandler) searchFull(c echo.Context) error {
	ctx := c.Request().Context()
	req := &coModel.SearchRequest[model.QueryFullGoods]{}
	if err := c.Bind(req); err != nil {
		logx.Error("search commodity bind failed", "error", err)
		return c.JSON(http.StatusBadRequest, err)
	}
	info, err := h.i.SearchFull(ctx, req)
	if err != nil {
		logx.Error("get commodity failed", "error", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, info)
}

func (h *inventoryHandler) produce(c echo.Context) error {
	info := &coModel.BatchGoods{}
	ctx := c.Request().Context()
	if err := c.Bind(info); err != nil {
		logx.Error("produce commodity bind failed", "error", err)
		return c.JSON(http.StatusBadRequest, err)
	}
	info.SourceCode = enum.SourceCodeInventoryCreateProduce
	if err := h.s.Save(ctx, info); err != nil {
		logx.Error("produce commodity failed", "error", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	logx.Info("produce commodity success", field.ID, info.BatchID)
	return c.JSON(http.StatusOK, coModel.NewIDResponse(info.BatchID))
}

func (h *inventoryHandler) consume(c echo.Context) error {
	info := &coModel.BatchGoods{}
	ctx := c.Request().Context()
	if err := c.Bind(info); err != nil {
		logx.Error("consume commodity bind failed", "error", err)
		return c.JSON(http.StatusBadRequest, err)
	}
	info.SourceCode = enum.SourceCodeInventoryCreateConsume
	if err := h.s.Save(ctx, info); err != nil {
		logx.Error("consume commodity failed", "error", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	logx.Info("consume commodity success", field.ID, info.BatchID)
	return c.JSON(http.StatusOK, coModel.NewIDResponse(info.BatchID))
}
func (h *inventoryHandler) updateLast(c echo.Context) error {
	ctx := c.Request().Context()
	req := &model.LastCommodity{}
	if err := c.Bind(req); err != nil {
		logx.Error("update full commodity bind failed", "error", err)
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := h.i.UpdateLast(ctx, req); err != nil {
		logx.Error("update last commodity failed", "error", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	logx.Info("update last commodity success", field.ID, req.ID)
	return c.JSON(http.StatusOK, coModel.NewIDResponse(req.ID))
}

func (h *inventoryHandler) updateFull(c echo.Context) error {
	ctx := c.Request().Context()
	req := &model.FullGoods{}
	if err := c.Bind(req); err != nil {
		logx.Error("update full commodity bind failed", "error", err)
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := h.i.UpdateFull(ctx, req); err != nil {
		logx.Error("update full commodity failed", "error", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	logx.Info("update full commodity success", field.ID, req.ID)
	return c.JSON(http.StatusOK, coModel.NewIDResponse(req.ID))
}
