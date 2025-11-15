package web

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/icdb37/bfsm/internal/constx/field"
	"github.com/icdb37/bfsm/internal/features/company/model"
	"github.com/icdb37/bfsm/internal/features/company/service"
	"github.com/icdb37/bfsm/internal/infra/logx"
	coModel "github.com/icdb37/bfsm/internal/model"
)

type companyHandler struct {
	s service.CompanyServer
}

func (u *companyHandler) search(c echo.Context) error {
	req := &coModel.SearchRequest[model.QueryCompany]{}
	ctx := c.Request().Context()
	if err := c.Bind(req); err != nil {
		logx.Error("search company bind failed", "error", err)
		return c.JSON(http.StatusBadRequest, err)
	}
	if req.Query == nil {
		req.Query = &model.QueryCompany{}
	}
	resp, err := u.s.Search(ctx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, resp)
}

func (u *companyHandler) selectAll(c echo.Context) error {
	ctx := c.Request().Context()
	resp, err := u.s.SelectAll(ctx)
	if err != nil {
		logx.Error("select all company failed", "error", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, resp)
}

func (u *companyHandler) get(c echo.Context) error {
	id := c.Param(field.ID)
	ctx := c.Request().Context()
	info, err := u.s.Get(ctx, id)
	if err != nil {
		logx.Error("get company failed", "error", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, info)
}
func (u *companyHandler) create(c echo.Context) error {
	info := &model.EntireCompany{}
	ctx := c.Request().Context()
	if err := c.Bind(info); err != nil {
		logx.Error("create company bind failed", "error", err)
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := u.s.Create(ctx, info); err != nil {
		logx.Error("create company failed", "error", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	logx.Info("create company success", field.ID, info.ID)
	return c.JSON(http.StatusOK, coModel.NewIDResponse(info.ID))
}

func (u *companyHandler) update(c echo.Context) error {
	info := &model.EntireCompany{}
	ctx := c.Request().Context()
	if err := c.Bind(info); err != nil {
		logx.Error("update company bind failed", "error", err)
		return err
	}
	info.ID = c.Param(field.ID)
	if err := u.s.Update(ctx, info); err != nil {
		logx.Error("update company failed", "error", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	logx.Info("update company success", field.ID, info.ID)
	return c.JSON(http.StatusOK, coModel.NewIDResponse(info.ID))
}
func (u *companyHandler) delete(c echo.Context) error {
	id := c.Param(field.ID)
	ctx := c.Request().Context()
	if err := u.s.Delete(ctx, id); err != nil {
		logx.Error("delete company failed", "error", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	logx.Info("delete company success", field.ID, id)
	return nil
}
