package main

import (
	"fmt"
	"lrucache/internal/api/handlers"
	"lrucache/internal/config"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.InitConfig("APP")
	if err != nil {
		panic(err)
	}
	config.Cfg = cfg

	router := gin.Default()
	handlerService := handlers.NewHandlers(router)
	handlerService.Registry()
	_ = router.Run(fmt.Sprintf(":%d", 8080))
}
