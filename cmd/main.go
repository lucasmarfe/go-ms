package main

import (
	"github.com/gin-gonic/gin"
	"go-proj/internal/domain/user"
	"log"
)

func main() {
	//User
	repo := user.NewInMemoryRepo()
	service := user.NewService(repo)
	handler := user.NewHandler(service)
	//Server
	ginServer := gin.Default()
	//Custom handlers
	handler.RegisterRoutes(ginServer)
	//Run server
	if err := ginServer.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
