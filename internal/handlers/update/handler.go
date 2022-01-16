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

// Update increment total, action and country counters.
// Request method should be "GET".
// Return 500, if service.Update finished with error.
// Return 400 "not enough query parameters", if parameters were incorrect
func (h *Handle) Update(c echo.Context) error {
	if c.Request().Method != http.MethodGet {
		return c.String(http.StatusBadRequest, "incorrect method")
	}

	action := c.QueryParams().Get("action")
	country := c.QueryParams().Get("country")
	if action == "" || country == "" {
		c.Logger().Info("not enough query parameters")
		return c.String(http.StatusBadRequest, "not enough query parameters")
	}

	_ = h.groupService.Update(action, country)
	c.Response().Header().Set(echo.HeaderContentType, "image/gif")
	c.Response().WriteHeader(http.StatusOK)
	return c.File("counter.gif")
}
