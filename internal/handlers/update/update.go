package update

import (
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

	action := c.QueryParams().Get("actions")
	country := c.QueryParams().Get("countries")
	if action == "" || country == "" {
		return c.String(http.StatusBadRequest, "not enough params in request ")
	}

	if err := h.groupService.Update(action, country); err != nil {
		return c.String(http.StatusInternalServerError, "data wasn't added to database")
	}

	return c.File("counter.gif")
}
