package tracks

import (
	"context"
	"github.com/fuadvi/music-catalog/internal/middleware"
	"github.com/fuadvi/music-catalog/internal/models/spotify"
	"github.com/gin-gonic/gin"
)

//go:generate mockgen -source=handler.go -destination=handler_mock_test.go -package=tracks
type service interface {
	Search(ctx context.Context, query string, pageSize, pageIndex int) (*spotify.SearchResponse, error)
}

type Handler struct {
	*gin.Engine
	Service service
}

func NewHandler(api *gin.Engine, service service) *Handler {
	return &Handler{
		Engine:  api,
		Service: service,
	}
}

func (h *Handler) RegisterRoute() {
	route := h.Group("/tracks")
	route.Use(middleware.AuthMiddleware())
	route.GET("search", h.Search)
}
