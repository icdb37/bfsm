package web

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/icdb37/bfsm/internal/constx/field"
	"github.com/icdb37/bfsm/internal/features/purchase/model"
	"github.com/icdb37/bfsm/internal/features/purchase/service"
	"github.com/icdb37/bfsm/internal/infra/logx"
	coModel "github.com/icdb37/bfsm/internal/model"
)

type batchHandler struct {
	s service.BatchServer
}

func (h *batchHandler) search(c echo.Context) error {
	req := &coModel.SearchRequest[model.QueryPurchase]{}
	ctx := c.Request().Context()
	if err := c.Bind(req); err != nil {
		logx.Error("search purchase batch bind failed", "error", err)
		return c.JSON(http.StatusBadRequest, err)
	}
	resp, err := h.s.Search(ctx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, resp)
}
func (h *batchHandler) get(c echo.Context) error {
	id := c.Param(field.ID)
	ctx := c.Request().Context()
	info, err := h.s.Get(ctx, id)
	if err != nil {
		logx.Error("get purchase batch failed", "error", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, info)
}
func (h *batchHandler) create(c echo.Context) error {
	info := &model.PurchaseBatch{}
	ctx := c.Request().Context()
	if err := c.Bind(info); err != nil {
		logx.Error("create purchase batch bind failed", "error", err)
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := h.s.Create(ctx, info); err != nil {
		logx.Error("create purchase batch failed", "error", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	logx.Info("create purchase batch success", field.ID, info.ID)
	return c.JSON(http.StatusOK, coModel.NewIDResponse(info.ID))
}

func (h *batchHandler) update(c echo.Context) error {
	info := &model.PurchaseBatch{}
	ctx := c.Request().Context()
	if err := c.Bind(info); err != nil {
		logx.Error("update purchase batch bind failed", "error", err)
		return err
	}
	info.ID = c.Param(field.ID)
	if err := h.s.Update(ctx, info); err != nil {
		logx.Error("update purchase batch failed", "error", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	logx.Info("update purchase batch success", field.ID, info.ID)
	return c.JSON(http.StatusOK, coModel.NewIDResponse(info.ID))
}
func (h *batchHandler) delete(c echo.Context) error {
	id := c.Param(field.ID)
	ctx := c.Request().Context()
	if err := h.s.Delete(ctx, id); err != nil {
		logx.Error("delete purchase batch failed", "error", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	logx.Info("delete purchase batch success", field.ID, id)
	return nil
}

func (h *batchHandler) updateStatus(c echo.Context) error {
	req := &coModel.UpdateStatus{}
	ctx := c.Request().Context()
	if err := c.Bind(req); err != nil {
		logx.Error("update purchase batch status bind failed", "error", err)
		return c.JSON(http.StatusBadRequest, err)
	}
	req.ID = c.Param(field.ID)
	if err := h.s.UpdateStatus(ctx, req); err != nil {
		logx.Error("update purchase batch status failed", "error", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	logx.Info("update purchase batch status success", field.ID, req.ID)
	return c.JSON(http.StatusOK, coModel.NewIDResponse(req.ID))
}
