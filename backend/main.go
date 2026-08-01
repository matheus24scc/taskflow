/*
TaskFlow - Agile Team Productivity Suite
Backend API with WebSocket support for real-time Kanban boards
*/

package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

// Configure logger
var logger = zerolog.New(os.Stdout).With().Timestamp().Logger()

// Config
type Config struct {
	Port     string
	Database string
}

func getConfig() Config {
	return Config{
		Port:     getEnv("PORT", "8080"),
		Database: getEnv("DATABASE_URL", "postgres://user:pass@localhost:5432/taskflow"),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

// Task model
type Task struct {
	ID          string    `json:"id" bson:"_id"`
	BoardID     string    `json:"boardId" bson:"board_id"`
	Title       string    `json:"title" bson:"title"`
	Description string    `json:"description,omitempty" bson:"description,omitempty"`
	Status      string    `json:"status" bson:"status"`
	Priority    string    `json:"priority" bson:"priority"`
	AssigneeID  string    `json:"assigneeId,omitempty" bson:"assignee_id,omitempty"`
	CreatedAt   time.Time `json:"createdAt" bson:"created_at"`
	UpdatedAt   time.Time `json:"updatedAt" bson:"updated_at"`
}

// Board model
type Board struct {
	ID          string    `json:"id" bson:"_id"`
	Title       string    `json:"title" bson:"title"`
	Description string    `json:"description,omitempty" bson:"description,omitempty"`
	Columns     []Column  `json:"columns" bson:"columns"`
	CreatedBy   string    `json:"createdBy" bson:"created_by"`
	CreatedAt   time.Time `json:"createdAt" bson:"created_at"`
	UpdatedAt   time.Time `json:"updatedAt" bson:"updated_at"`
}

// Column model
type Column struct {
	ID    string `json:"id" bson:"_id"`
	Title string `json:"title" bson:"title"`
	Order int    `json:"order" bson:"order"`
}

// Database connection
var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("postgres", getConfig().Database)
	if err != nil {
		log.Fatal(err)
	}
}

// Health check handler
func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":    "healthy",
		"timestamp": time.Now().UTC(),
		"service":   "taskflow-backend",
		"version":   "1.0.0",
	})
}

// WebSocket upgrade but not fully implemented (requires Gorilla WebSocket)
func setupRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		// Health check
		api.GET("/health", healthCheck)

		// Boards
		api.GET("/boards", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "List boards - TODO: implement with MongoDB"})
		})
		api.POST("/boards", func(c *gin.Context) {
			c.JSON(http.StatusCreated, gin.H{"message": "Create board - TODO: implement"})
		})
		api.GET("/boards/:id", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Get board - TODO: implement"})
		})

		// Tasks
		api.GET("/tasks", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "List tasks - TODO: implement with WebSocket"})
		})
		api.POST("/tasks", func(c *gin.Context) {
			c.JSON(http.StatusCreated, gin.H{"message": "Create task - TODO: implement"})
		})
		api.GET("/tasks/:id", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Get task - TODO: implement"})
		})
		api.PUT("/tasks/:id", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Update task - TODO: implement"})
		})

		// WebSocket endpoint placeholder
		api.GET("/ws", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "WebSocket endpoint - TODO: implement real-time updates",
				"note": "Requires Gorilla WebSocket for full implementation",
			})
		})
	}
}

func main() {
	config := getConfig()

	// Initialize logger
	logger.Info().Msg("Starting TaskFlow Backend")

	// Initialize database
	initDB()

	// Set Gin mode
	gin.SetMode(gin.ReleaseMode)

	// Create router
	router := gin.Default()

	// Security middleware
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("X-Content-Type-Options", "nosniff")
		c.Writer.Header().Set("X-Frame-Options", " DENY")
		c.Writer.Header().Set("X-XSS-Protection", "1; mode=block")
		c.Next()
	})

	// Setup routes
	setupRoutes(router)

	// Start server
	srv := &http.Server{
		Addr:    ":" + config.Port,
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal().Err(err).Msg("Server failed")
		}
	}()

	logger.Info().Msgf("TaskFlow API server starting on port %s", config.Port)
	log.Printf("🚀 TaskFlow API listening on :%s", config.Port)

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	<-quit
	logger.Info().Msg("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal().Err(err).Msg("Server forced to shutdown")
	}

	logger.Info().Msg("Server exiting")
}
