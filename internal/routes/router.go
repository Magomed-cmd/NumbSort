package routes

import (
	"context"
	"net/http"
	"time"

	"numbsort/internal/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter(h *handler.NumberHandler) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery(), gin.Logger())

	r.POST("/numbers", h.AddNumber)
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	return r
}

type HTTPServer struct {
	srv *http.Server
}

func NewServer(addr string, handler http.Handler) *HTTPServer {
	return &HTTPServer{
		srv: &http.Server{
			Addr:    addr,
			Handler: handler,
		},
	}
}

func (s *HTTPServer) Start() error {
	return s.srv.ListenAndServe()
}

func (s *HTTPServer) Stop(ctx context.Context) error {
	shutdownCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return s.srv.Shutdown(shutdownCtx)
}
