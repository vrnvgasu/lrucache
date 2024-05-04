package cache

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type RemoveRequest struct {
	Key string `form:"key" binding:"required"`
}

type RemoveResponse struct {
	OK bool `json:"ok"`
}

func (l *LRUHandlerImpl) Remove(c *gin.Context) {
	var req RemoveRequest
	if err := c.ShouldBindWith(&req, binding.Query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ok := l.lru.Remove(req.Key)

	c.JSON(http.StatusOK, RemoveResponse{
		OK: ok,
	})
}
