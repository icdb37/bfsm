package web

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/icdb37/bfsm/internal/constx/enum"
	"github.com/icdb37/bfsm/internal/constx/field"
	"github.com/icdb37/bfsm/internal/features/bill/model"
	"github.com/icdb37/bfsm/internal/features/bill/service"
	"github.com/icdb37/bfsm/internal/infra/errx"
	"github.com/icdb37/bfsm/internal/infra/logx"
	coModel "github.com/icdb37/bfsm/internal/model"
)

type batchHandler struct {
	s service.BatchServer
}

func (h *batchHandler) search(c echo.Context) error {
	req := &coModel.SearchRequest[model.QueryBillBatch]{}
	ctx := c.Request().Context()
	if err := c.Bind(req); err != nil {
		logx.Error("search bill batch bind failed", "error", err)
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
		logx.Error("get bill batch failed", "error", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, info)
}
func (h *batchHandler) create(c echo.Context) error {
	info := &model.BillBatch{}
	ctx := c.Request().Context()
	if err := c.Bind(info); err != nil {
		logx.Error("create bill batch bind failed", "error", err)
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := h.s.Create(ctx, info); err != nil {
		logx.Error("create bill batch failed", "error", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	logx.Info("create bill batch success", field.ID, info.ID)
	return c.JSON(http.StatusOK, coModel.NewIDResponse(info.ID))
}

func (h *batchHandler) update(c echo.Context) error {
	info := &model.BillBatch{}
	ctx := c.Request().Context()
	if err := c.Bind(info); err != nil {
		logx.Error("update bill batch bind failed", "error", err)
		return err
	}
	info.ID = c.Param(field.ID)
	if err := h.s.Update(ctx, info); err != nil {
		logx.Error("update bill batch failed", "error", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	logx.Info("update bill batch success", field.ID, info.ID)
	return c.JSON(http.StatusOK, coModel.NewIDResponse(info.ID))
}

func (h *batchHandler) updateStatus(c echo.Context) error {
	id := c.Param(field.ID)
	strStatus := c.Param(field.Status)
	status, err := strconv.Atoi(strStatus)
	if err != nil {
		logx.Error("update bill batch status parse failed", "status", strStatus, "error", err)
		return c.JSON(http.StatusBadRequest, errx.NewParam(field.Status, strStatus))
	}
	ctx := c.Request().Context()
	switch enum.StatusCode(status) {
	case enum.StatusCodeApproved:
		err = h.s.Approve(ctx, id)
	case enum.StatusCodeCompleted:
		err = h.s.Complete(ctx, id)
	case enum.StatusCodeCanceled:
		err = h.s.Cancel(ctx, id)
	case enum.StatusCodeClosed:
		err = h.s.Close(ctx, id)
	default:
		err = errx.NewParam(field.Status, strStatus)
	}
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, coModel.NewIDResponse(id))
}

func (h *batchHandler) updateAmount(c echo.Context) error {
	id := c.Param(field.ID)
	info := &coModel.UpdateAmount{}
	ctx := c.Request().Context()
	if err := c.Bind(info); err != nil {
		logx.Error("update bill batch amount bind failed", "error", err)
		return c.JSON(http.StatusBadRequest, err)
	}
	info.ID = id
	if err := h.s.UpdateAmount(ctx, info); err != nil {
		logx.Error("update bill batch amount failed", "error", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, coModel.NewIDResponse(id))
}
