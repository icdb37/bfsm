package web

import (
	"net/http"

	"github.com/icdb37/bfsm/internal/features/user/model"
	"github.com/icdb37/bfsm/internal/features/user/service"
	"github.com/icdb37/bfsm/internal/infra/logx"
	coModel "github.com/icdb37/bfsm/internal/model"
	"github.com/labstack/echo/v4"
)

type userImpl struct {
	s service.Server
}

func Init(e *echo.Echo) error {
	svc, err := service.New()
	if err != nil {
		return err
	}
	u := &userImpl{s: svc}
	g := e.Group("/api/v1/users")
	g.POST("", u.create)
	g.POST("/search", u.search)
	g.PUT("/:id", u.update)
	g.DELETE("/:id", u.delete)
	g.GET("/:id", u.get)
	return nil
}

func (u *userImpl) create(c echo.Context) error {
	info := &model.EntireUser{}
	ctx := c.Request().Context()
	if err := c.Bind(info); err != nil {
		logx.Error("create user bind failed", "error", err)
		return err
	}
	if err := u.s.CreateUser(ctx, info); err != nil {
		logx.Error("create user failed", "error", err)
		return err
	}
	return c.JSON(http.StatusOK, coModel.NewIDResponse(info.ID))
}
func (u *userImpl) search(c echo.Context) error {
	req := &model.SearchRequest{}
	ctx := c.Request().Context()
	if err := c.Bind(req); err != nil {
		logx.Error("search user bind failed", "error", err)
		return err
	}
	resp, err := u.s.SearchUser(ctx, req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, resp)
}

func (u *userImpl) update(c echo.Context) error {
	info := &model.EntireUser{}
	ctx := c.Request().Context()
	if err := c.Bind(info); err != nil {
		logx.Error("update user bind failed", "error", err)
		return err
	}
	if err := u.s.UpdateUser(ctx, info); err != nil {
		logx.Error("update user failed", "error", err)
		return err
	}
	return c.JSON(http.StatusOK, coModel.NewIDResponse(info.ID))
}
func (u *userImpl) delete(c echo.Context) error {
	id := c.Param("id")
	ctx := c.Request().Context()
	if err := u.s.DeleteUser(ctx, id); err != nil {
		logx.Error("delete user failed", "error", err)
		return err
	}
	return nil
}

func (u *userImpl) get(c echo.Context) error {
	id := c.Param("id")
	ctx := c.Request().Context()
	user, err := u.s.GetUser(ctx, id)
	if err != nil {
		logx.Error("get user failed", "error", err)
		return err
	}
	return c.JSON(http.StatusOK, user)
}
