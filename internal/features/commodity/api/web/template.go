package web

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/icdb37/bfsm/internal/constx/field"
	"github.com/icdb37/bfsm/internal/features/commodity/model"
	"github.com/icdb37/bfsm/internal/features/commodity/service"
	"github.com/icdb37/bfsm/internal/infra/logx"
	coModel "github.com/icdb37/bfsm/internal/model"
)

type templateHandler struct {
	s service.TemplateServer
}

func (u *templateHandler) search(c echo.Context) error {
	req := &coModel.SearchRequest[model.QueryTemplate]{}
	ctx := c.Request().Context()
	if err := c.Bind(req); err != nil {
		logx.Error("search template bind failed", "error", err)
		return c.JSON(http.StatusBadRequest, err)
	}
	resp, err := u.s.Search(ctx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, resp)
}
func (u *templateHandler) get(c echo.Context) error {
	id := c.Param(field.ID)
	ctx := c.Request().Context()
	info, err := u.s.Get(ctx, id)
	if err != nil {
		logx.Error("get template failed", "error", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, info)
}
func (u *templateHandler) create(c echo.Context) error {
	info := &model.EntireTemplate{}
	ctx := c.Request().Context()
	if err := c.Bind(info); err != nil {
		logx.Error("create template bind failed", "error", err)
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := u.s.Create(ctx, info); err != nil {
		logx.Error("create template failed", "error", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	logx.Info("create template success", field.ID, info.ID)
	return c.JSON(http.StatusOK, coModel.NewIDResponse(info.ID))
}

func (u *templateHandler) update(c echo.Context) error {
	info := &model.EntireTemplate{}
	ctx := c.Request().Context()
	if err := c.Bind(info); err != nil {
		logx.Error("update template bind failed", "error", err)
		return err
	}
	info.ID = c.Param(field.ID)
	if err := u.s.Update(ctx, info); err != nil {
		logx.Error("update template failed", "error", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	logx.Info("update template success", field.ID, info.ID)
	return c.JSON(http.StatusOK, coModel.NewIDResponse(info.ID))
}
func (u *templateHandler) delete(c echo.Context) error {
	id := c.Param(field.ID)
	ctx := c.Request().Context()
	if err := u.s.Delete(ctx, id); err != nil {
		logx.Error("delete template failed", "error", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	logx.Info("delete template success", field.ID, id)
	return nil
}
