package web

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/icdb37/bfsm/internal/constx/field"
	"github.com/icdb37/bfsm/internal/features/bill/model"
	"github.com/icdb37/bfsm/internal/features/bill/service"
	"github.com/icdb37/bfsm/internal/infra/logx"
	coModel "github.com/icdb37/bfsm/internal/model"
)

// 商品服务接口
type dealHandler struct {
	s service.DealServer
}

func (h *dealHandler) search(c echo.Context) error {
	req := &coModel.SearchRequest[model.QueryBillDeal]{}
	ctx := c.Request().Context()
	if err := c.Bind(req); err != nil {
		logx.Error("search bill deal bind failed", "error", err)
		return c.JSON(http.StatusBadRequest, err)
	}
	resp, err := h.s.Search(ctx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *dealHandler) update(c echo.Context) error {
	info := &model.BillDeal{}
	ctx := c.Request().Context()
	if err := c.Bind(info); err != nil {
		logx.Error("update bill deal bind failed", "error", err)
		return err
	}
	info.ID = c.Param(field.ID)
	if err := h.s.Update(ctx, info); err != nil {
		logx.Error("update bill deal failed", "error", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	logx.Info("update bill deal success", field.ID, info.ID)
	return c.JSON(http.StatusOK, coModel.NewIDResponse(info.ID))
}
