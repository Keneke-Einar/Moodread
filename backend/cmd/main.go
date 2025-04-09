package main

import (
	"backend/config"
	"backend/internal/database"
	"backend/pkg/logger"
)

func main() {
	// Initialize logger
	logger.InitLogger(false) // Use false for development, true for production
	log := logger.GetLogger()
	defer log.Sync()

	log.Info("Starting Moodread application...")

	cfg := config.MustGetConfig()
	log.Infow("Loaded configuration", "config", cfg)

	err := database.InitDB(&cfg.Database)
	if err != nil {
		log.Fatalw("Failed to initialize database", "error", err)
	}

	log.Info("Moodread application started successfully")
}
