package getdata

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
	return handler.Upload
}

type groupInterface interface {
	GetData() (string, error)
}

func (h *Handle) Upload(c echo.Context) error {
	if c.Request().Method != http.MethodGet {
		return c.String(http.StatusBadRequest, "incorrect method")
	}

	result, err := h.groupService.GetData()
	if err != nil {
		c.Logger().Errorf("can't get data from database", errors.Unwrap(err))
		return c.String(http.StatusInternalServerError, "can't get data from database")
	}

	return c.String(http.StatusOK, result)
}
