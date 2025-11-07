package web

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/icdb37/bfsm/internal/features/purchase/model"
	"github.com/icdb37/bfsm/internal/features/purchase/service"
	"github.com/icdb37/bfsm/internal/infra/logx"
	coModel "github.com/icdb37/bfsm/internal/model"
)

// 商品服务接口
type goodsHandler struct {
	s service.GoodsServer
}

func (h *goodsHandler) search(c echo.Context) error {
	req := &coModel.SearchRequest[model.QueryPurchaseGoods]{}
	ctx := c.Request().Context()
	if err := c.Bind(req); err != nil {
		logx.Error("search purchase goods bind failed", "error", err)
		return c.JSON(http.StatusBadRequest, err)
	}
	resp, err := h.s.Search(ctx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, resp)
}
