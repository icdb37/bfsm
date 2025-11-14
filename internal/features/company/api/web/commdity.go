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

// 商品服务接口
type commodityHandler struct {
	s service.CommodityServer
}

func (u *commodityHandler) search(c echo.Context) error {
	req := &coModel.SearchRequest[model.QueryCommodity]{}
	ctx := c.Request().Context()
	if err := c.Bind(req); err != nil {
		logx.Error("search commodity bind failed", "error", err)
		return c.JSON(http.StatusBadRequest, err)
	}
	req.Query.CompanyID = c.Param(field.CompanyID)
	resp, err := u.s.Search(ctx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, resp)
}

func (u *commodityHandler) get(c echo.Context) error {
	id := c.Param(field.ID)
	ctx := c.Request().Context()
	info, err := u.s.Get(ctx, id)
	if err != nil {
		logx.Error("get commodity failed", "error", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, info)
}

func (u *commodityHandler) create(c echo.Context) error {
	infos := []*model.EntireCommodity{}
	ctx := c.Request().Context()
	if err := c.Bind(&infos); err != nil {
		logx.Error("create commodity bind failed", "error", err)
		return c.JSON(http.StatusBadRequest, err)
	}
	for _, item := range infos {
		item.CompanyID = c.Param(field.CompanyID)
	}
	if err := u.s.Create(ctx, infos); err != nil {
		logx.Error("create commodity failed", "error", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	logx.Info("create commodity success")
	return nil
}

func (u *commodityHandler) update(c echo.Context) error {
	info := &model.EntireCommodity{}
	ctx := c.Request().Context()
	if err := c.Bind(info); err != nil {
		logx.Error("update commodity bind failed", "error", err)
		return c.JSON(http.StatusBadRequest, err)
	}
	info.ID = c.Param(field.ID)
	info.CompanyID = c.Param(field.CompanyID)
	if err := u.s.Update(ctx, info); err != nil {
		logx.Error("update commodity failed", "error", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	logx.Info("delete commodity success", field.ID, info.ID)
	return c.JSON(http.StatusOK, coModel.NewIDResponse(info.ID))
}
func (u *commodityHandler) delete(c echo.Context) error {
	companyID, id := c.Param(field.CompanyID), c.Param(field.ID)
	ctx := c.Request().Context()
	if err := u.s.Delete(ctx, id); err != nil {
		logx.Error("delete commodity failed", "error", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	logx.Info("delete commodity success", field.CompanyID, companyID, field.ID, id)
	return nil
}
