package update

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

func (t *TestCases) Update(action, country string) error {
	if t.findAndUpdate {
		return nil
	}
	return fmt.Errorf("some error")
}

func TestHandle_Update(t *testing.T)  {
	var testCases = []TestCases{
		{
			method:         http.MethodGet,
			path:           "/counter.gif?action=view&country=ru",
			findAndUpdate:  true,
			expectedStatus: http.StatusOK,
			expectedBody:   "GIF89a\x01\x00\x01\x00\x00\x00\x00!" +
				"\xf9\x04\x01\x00\x00\x00\x01,\x00\x00\x00\x00\x01\x00\x01\x00\x00\x01\x00",
		},
		{
			method:         http.MethodGet,
			path:           "/counter.gif?action=view&country=ru",
			findAndUpdate:  false,
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "data wasn't added to database",
		},
		{
			method:         http.MethodPut,
			path:           "/counter.gif?action=view&country=ru",
			findAndUpdate:  true,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "incorrect method",
		},
		{
			method:         http.MethodGet,
			path:           "/counter.gif?country=ru",
			findAndUpdate:  false,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "not enough query parameters",
		},
	}

	e := echo.New()
	e.Logger.SetLevel(log.OFF)
	for i := range testCases {
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