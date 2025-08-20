package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// PingResponse represents the response structure for ping endpoint
type PingResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

// HealthResponse represents the response structure for health check
type HealthResponse struct {
	Service string `json:"service"`
	Status  string `json:"status"`
	Version string `json:"version"`
}

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default values")
	}

	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Set gin mode based on environment
	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Create Gin router
	router := gin.Default()

	// Add middleware for CORS (Cross-Origin Resource Sharing)
	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		
		c.Next()
	})

	// Basic ping endpoint - health check
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, PingResponse{
			Message: "pong",
			Status:  "healthy",
		})
	})

	// Enhanced health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, HealthResponse{
			Service: "Go API with Gin",
			Status:  "running",
			Version: "1.0.0",
		})
	})

	// Endpoint demonstrating path parameters
	router.GET("/user/:id", func(c *gin.Context) {
		userID := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"user_id": userID,
			"message": "User retrieved successfully",
		})
	})

	// Endpoint demonstrating query parameters
	router.GET("/search", func(c *gin.Context) {
		// Get query parameters
		query := c.Query("q")           // Required query parameter
		limit := c.DefaultQuery("limit", "10") // Optional with default
		page := c.DefaultQuery("page", "1")    // Optional with default

		// Validate required parameter
		if query == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Query parameter 'q' is required",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"query":   query,
			"limit":   limit,
			"page":    page,
			"results": []string{}, // Placeholder for actual search results
		})
	})

	// Endpoint combining both path and query parameters
	router.GET("/user/:id/posts", func(c *gin.Context) {
		userID := c.Param("id")
		category := c.DefaultQuery("category", "all")
		sort := c.DefaultQuery("sort", "date")

		c.JSON(http.StatusOK, gin.H{
			"user_id":  userID,
			"category": category,
			"sort":     sort,
			"posts":    []string{}, // Placeholder for actual posts
		})
	})

	// Start server
	log.Printf("Server starting on port %s", port)
	log.Printf("Health check available at: http://localhost:%s/ping", port)
	
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}