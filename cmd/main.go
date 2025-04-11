package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib" // driver
	"go-proj/internal/config"
	"go-proj/internal/domain/user"
	"log"
)

func main() {
	//Config load
	cfg := config.LoadConfigFromPath("./configs")
	//Database setup
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.DBName,
		cfg.Database.SSLMode,
	)
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}
	//User
	repo := user.NewPostgresRepo(db)
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
