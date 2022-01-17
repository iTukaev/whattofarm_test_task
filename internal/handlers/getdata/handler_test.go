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
	getCorrect bool
	expectedStatus int
	expectedBody   string
}

func (t *TestCases) GetData(timeBegin, timeEnd string) ([]byte, error) {
	if t.getCorrect {
		return []byte(fmt.Sprintf("time: %s, time: %s", timeBegin, timeEnd)), nil
	}
	return nil, fmt.Errorf("some error")
}

func TestHandle_Upload(t *testing.T)  {
	var testCases = []TestCases{
		{
			method:         http.MethodGet,
			path:           "/?time_begin=2022-01-17_11_%2b05&time_end=2022-01-17_14_%2b05",
			getCorrect:  true,
			expectedStatus: http.StatusOK,
			expectedBody:   "time: 2022-01-17_11_+05, time: 2022-01-17_14_+05",
		},
		{
			method:         http.MethodGet,
			path:           "/?time_begin=2022-01-17_11_%2b05&time_end=2022-01-17_14_%2b05",
			getCorrect:  false,
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "can't get data from database",
		},
		{
			method:         http.MethodPut,
			path:           "/?time_begin=2022-01-17_11_%2b05&time_end=2022-01-17_14_%2b05",
			getCorrect:  true,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "incorrect method",
		},
		{
			method:         http.MethodGet,
			path:           "/?time_begin=2022-01-17_11_%2b05",
			getCorrect:  true,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "not enough query parameters",
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