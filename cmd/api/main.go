package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"

	_ "github.com/wow/nigger/docs/swagger"
	"github.com/wow/nigger/internal/config"
	"github.com/wow/nigger/internal/handler"
	"github.com/wow/nigger/internal/middleware"
)

// @title           Go Backend API
// @version         1.0.0
// @description     A production-ready Go backend API with JWT auth, structured logging, and Swagger docs.

// @contact.name   API Support
// @contact.email  support@example.com

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Enter your bearer token in the format: Bearer {token}

const Version = "1.0.0"

func main() {
	// Load config
	cfg := config.Load()

	// Initialize logger
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}
	defer logger.Sync()

	// Set Gin mode
	gin.SetMode(cfg.Server.Mode)

	// Create router
	r := gin.New()

	// Global middleware
	r.Use(middleware.RequestID())
	r.Use(middleware.Logger(logger))
	r.Use(middleware.Recovery(logger))
	r.Use(middleware.Auth(cfg.JWT.Secret))

	// Public routes (no auth required)
	r.GET("/health", handler.Health(Version))
	r.GET("/api/docs", handler.GetDocs)

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API v1 routes
	v1 := r.Group("/api/v1")
	{
		// Auth routes
		auth := v1.Group("/auth")
		{
			auth.POST("/login", handler.Login(cfg.JWT.Secret))
			auth.GET("/me", handler.GetMe)
		}

		// Test routes
		v1.GET("/ping", handler.Ping)
		v1.POST("/echo", handler.Echo)
	}

	// Start server
	addr := fmt.Sprintf(":%s", cfg.Server.Port)
	logger.Info("starting server",
		zap.String("address", addr),
		zap.String("version", Version),
	)

	if err := r.Run(addr); err != nil {
		logger.Fatal("failed to start server", zap.Error(err))
	}
}
