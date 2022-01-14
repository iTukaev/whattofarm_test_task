package update

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)


type Handle struct {
	groupService groupInterface
}

func NewHandler(service groupInterface) func(c echo.Context) error {
	handler := &Handle {
		groupService: service,
	}
	return handler.Update
}


type groupInterface interface {
	Update(action, country string) error
}

func (h *Handle) Update(c echo.Context) error {
	if c.Request().Method != http.MethodGet {
		return c.String(http.StatusBadRequest, "incorrect method")
	}

	action := c.QueryParams().Get("action")
	country := c.QueryParams().Get("country")
	if action == "" || country == "" {
		c.Logger().Info("not enough query parameters")
		return c.String(http.StatusBadRequest, "not enough params in request ")
	}

	if err := h.groupService.Update(action, country); err != nil {
		c.Logger().Errorf("data wasn't added to database", errors.Unwrap(err))
		return c.String(http.StatusInternalServerError, "data wasn't added to database")
	}
	c.Response().Header().Set(echo.HeaderContentType, "image/gif")
	c.Response().WriteHeader(http.StatusOK)
	return c.File("counter.gif")
}
