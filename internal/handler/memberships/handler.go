package memberships

import (
	"github.com/fuadvi/music-catalog/internal/models/memberships"
	"github.com/gin-gonic/gin"
)

//go:generate mockgen -source=handler.go -destination=handler_mock_test.go -package=memberships
type service interface {
	SignUp(request memberships.SignUpRequest) error
	Login(request memberships.LoginRequest) (string, error)
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
	route := h.Group("/memberships")
	route.POST("/sign-up", h.SignUp)
	route.POST("/login", h.Login)
}
