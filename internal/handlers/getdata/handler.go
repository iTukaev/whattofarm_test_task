package getdata

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)


type Handle struct {
	handleService handleInterface
}

func NewHandler(service handleInterface) func(c echo.Context) error {
	handler := &Handle {
		handleService: service,
	}
	return handler.Upload
}

type handleInterface interface {
	GetData() (string, error)
}

// Upload return MongoDB's document as a JSON string.
// Request method should be "GET".
// Return 500, if GetData finished with error
func (h *Handle) Upload(c echo.Context) error {
	if c.Request().Method != http.MethodGet {
		return c.String(http.StatusBadRequest, "incorrect method")
	}

	result, err := h.handleService.GetData()
	if err != nil {
		c.Logger().Errorf("can't get data from database", errors.Unwrap(err))
		return c.String(http.StatusInternalServerError, "can't get data from database")
	}

	return c.String(http.StatusOK, result)
}
