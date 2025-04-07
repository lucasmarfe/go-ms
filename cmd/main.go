package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-proj/internal/config"
	"go-proj/internal/domain/user"
	"log"
)

func main() {
	//Config
	cfg := config.LoadConfig()
	//User
	repo := user.NewInMemoryRepo()
	service := user.NewService(repo)
	handler := user.NewHandler(service)
	//Server
	ginServer := gin.Default()
	//Custom handlers
	handler.RegisterRoutes(ginServer)
	//Run server
	if err := ginServer.Run(fmt.Sprintf(":%d", cfg.Server.Port)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
