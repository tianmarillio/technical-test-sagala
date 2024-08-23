package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tianmarillio/technical-test-sagala/src/config"
	"github.com/tianmarillio/technical-test-sagala/src/database"
	"github.com/tianmarillio/technical-test-sagala/src/routes"
)

func main() {
	// Load env variables
	config.LoadEnv()
	cfg := config.GetEnv()

	// Initialize database connection
	db := database.InitDB(cfg)

	// Initialize gin router
	r := gin.Default()
	routes.RegisterRoutes(r, db)

	// Start gin server
	portStr := fmt.Sprintf(":%d", cfg.Port)
	if err := r.Run(portStr); err != nil {
		log.Fatal("Failed to start server")
	}
}
