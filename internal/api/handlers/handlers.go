package handlers

import (
	"lrucache/internal/api/handlers/cache"

	"github.com/gin-gonic/gin"
)

type Handlers interface {
	Registry()
}

type handlers struct {
	router *gin.Engine
	cache  cache.LRUHandler
}

func NewHandlers(route *gin.Engine) Handlers {
	return &handlers{
		router: route,
		cache:  cache.NewLRUHandler(),
	}
}

func (h *handlers) Registry() {
	cacheGroup := h.router.Group("/cache")

	cacheGroup.POST("", h.cache.Add)
	cacheGroup.GET("", h.cache.Get)
	cacheGroup.DELETE("", h.cache.Remove)
}
