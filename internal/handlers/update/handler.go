package update

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
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

	if err := h.groupService.Update(action, country); err != nil {
		c.Logger().Errorf("data wasn't added to database", errors.Unwrap(err))
		return c.String(http.StatusInternalServerError, "data wasn't added to database")
	}

	f, err := os.Open("counter.gif")
	if err != nil {
		c.Logger().Warnf("os.Open response file error:", err)
		return c.String(http.StatusInternalServerError, "response file error")
	}
	return c.Stream(http.StatusOK, "image/gif", f)
}
