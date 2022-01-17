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
	GetData(timeBegin, timeEnd string) ([]byte, error)
}

// Upload return MongoDB's document as a JSON string.
// Request method should be "GET".
func (h *Handle) Upload(c echo.Context) error {
	if c.Request().Method != http.MethodGet {
		return c.String(http.StatusBadRequest, "incorrect method")
	}

	timeBegin := c.QueryParams().Get("time_begin")
	timeEnd := c.QueryParams().Get("time_end")
	if timeBegin == "" || timeEnd == "" {
		c.Logger().Info("not enough query parameters")
		return c.String(http.StatusBadRequest, "not enough query parameters")
	}

	result, err := h.handleService.GetData(timeBegin, timeEnd)
	if err != nil {
		c.Logger().Errorf("can't get data from database", errors.Unwrap(err))
		return c.String(http.StatusInternalServerError, "can't get data from database")
	}

	return c.Blob(http.StatusOK, echo.MIMEApplicationJSON, result)
}
