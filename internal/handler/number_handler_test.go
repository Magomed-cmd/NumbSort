package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"numbsort/internal/handler/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestAddNumber_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	svc := mocks.NewNumberService(t)
	svc.On("AddAndList", mock.Anything, 2).Return([]int{1, 2}, nil)

	handler := NewNumberHandler(svc)
	r := gin.New()
	r.POST("/numbers", handler.AddNumber)

	body := bytes.NewBufferString(`{"value":2}`)
	req := httptest.NewRequest(http.MethodPost, "/numbers", body)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	require.Contains(t, w.Body.String(), `"numbers":[1,2]`)
	svc.AssertExpectations(t)
}

func TestAddNumber_BadRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)
	handler := NewNumberHandler(mocks.NewNumberService(t))
	r := gin.New()
	r.POST("/numbers", handler.AddNumber)

	req := httptest.NewRequest(http.MethodPost, "/numbers", bytes.NewBufferString(`{}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code)
}

func TestAddNumber_ZeroAllowed(t *testing.T) {
	gin.SetMode(gin.TestMode)
	svc := mocks.NewNumberService(t)
	svc.On("AddAndList", mock.Anything, 0).Return([]int{0, 1}, nil)

	handler := NewNumberHandler(svc)
	r := gin.New()
	r.POST("/numbers", handler.AddNumber)

	req := httptest.NewRequest(http.MethodPost, "/numbers", bytes.NewBufferString(`{"value":0}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	require.Contains(t, w.Body.String(), `"numbers":[0,1]`)
	svc.AssertExpectations(t)
}

func TestAddNumber_ServiceError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	svc := mocks.NewNumberService(t)
	svc.On("AddAndList", mock.Anything, 1).Return(nil, assertErr{})

	handler := NewNumberHandler(svc)
	r := gin.New()
	r.POST("/numbers", handler.AddNumber)

	req := httptest.NewRequest(http.MethodPost, "/numbers", bytes.NewBufferString(`{"value":1}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusInternalServerError, w.Code)
	svc.AssertExpectations(t)
}

type assertErr struct{}

func (assertErr) Error() string { return "boom" }
