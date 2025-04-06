package main

import (
	"backend/config"
	"backend/internal/database"
	"fmt"
	"log"
)

func main() {
	cfg := config.MustGetConfig()

	err := database.InitDB(&cfg.Database)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	fmt.Println("Moodread")
}
