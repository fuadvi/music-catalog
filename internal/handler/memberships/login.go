package memberships

import (
	"github.com/fuadvi/music-catalog/internal/models/memberships"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) Login(c *gin.Context) {
	var request memberships.LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accessToken, err := h.Service.Login(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": accessToken})
}
