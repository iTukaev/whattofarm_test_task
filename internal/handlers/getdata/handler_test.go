package getdata

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)


type TestCases struct {
	method string
	path string
	findAndUpdate bool
	expectedStatus int
	expectedBody   string
}

func (t *TestCases) GetData() (string, error) {
	if t.findAndUpdate {
		return fmt.Sprintf("method: %s, path: %s", t.method, t.path), nil
	}
	return "", fmt.Errorf("some error")
}

func TestHandle_Upload(t *testing.T)  {
	var testCases = []TestCases{
		{
			method:         http.MethodGet,
			path:           "/",
			findAndUpdate:  true,
			expectedStatus: http.StatusOK,
			expectedBody:   "method: GET, path: /",
		},
		{
			method:         http.MethodGet,
			path:           "/",
			findAndUpdate:  false,
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "can't get data from database",
		},
		{
			method:         http.MethodPut,
			path:           "/",
			findAndUpdate:  true,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "incorrect method",
		},
	}

	e := echo.New()
	e.Logger.SetLevel(log.OFF)
	for i, _ := range testCases {
		req := httptest.NewRequest(testCases[i].method, testCases[i].path, nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h := NewHandler(&testCases[i])
		if assert.NoError(t, h(c)) {
			assert.Equal(t, testCases[i].expectedStatus, rec.Code)
			assert.Equal(t, testCases[i].expectedBody, rec.Body.String())
		}
	}
}