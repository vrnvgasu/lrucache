package cache

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type GetRequest struct {
	Key string `form:"key" binding:"required"`
}

type GetResponse struct {
	Value string `json:"value"`
	OK    bool   `json:"ok"`
}

func (l *LRUHandlerImpl) Get(c *gin.Context) {
	var req GetRequest
	if err := c.ShouldBindWith(&req, binding.Query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	value, ok := l.lru.Get(req.Key)

	c.JSON(http.StatusOK, GetResponse{
		Value: value,
		OK:    ok,
	})
}
