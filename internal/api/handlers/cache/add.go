package cache

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type AddRequest struct {
	Key   string `form:"key" binding:"required"`
	Value string `form:"value" binding:"required"`
}

type AddResponse struct {
	OK bool `json:"ok"`
}

func (l *LRUHandlerImpl) Add(c *gin.Context) {
	var req AddRequest
	if err := c.ShouldBindWith(&req, binding.Query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ok := l.lru.Add(req.Key, req.Value)

	c.JSON(http.StatusOK, AddResponse{
		OK: ok,
	})
}
