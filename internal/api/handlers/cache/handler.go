package cache

import (
	"fmt"
	"lrucache/internal/config"
	"lrucache/internal/service/lru"

	"github.com/gin-gonic/gin"
)

type LRUHandler interface {
	Add(c *gin.Context)
	Get(c *gin.Context)
	Remove(c *gin.Context)
}

type LRUHandlerImpl struct {
	lru lru.LRUCache
}

func NewLRUHandler() LRUHandler {
	cacheSize := lru.DefaultSize

	if config.Cfg != nil {
		cacheSize = config.Cfg.CacheSize
	}

	fmt.Println("cache size:", cacheSize)

	return &LRUHandlerImpl{
		lru: lru.NewService(cacheSize),
	}
}
