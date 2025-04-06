package database

import (
	"backend/config"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

const (
	maxOpenConns    = 25
	maxIdleConns    = 25
	connMaxLifetime = time.Hour
)

var DbConn *PostgresDB

type PostgresDB struct {
	db *gorm.DB
}

func InitDB(dbCfg *config.DatabaseConfig) error {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		dbCfg.Host,
		dbCfg.Port,
		dbCfg.User,
		dbCfg.Password,
		dbCfg.Name,
	)

	fmt.Printf("connecting to dsn: %s", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to open database connection: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get SQL DB from gorm DB: %w", err)
	}
	sqlDB.SetMaxOpenConns(maxOpenConns)
	sqlDB.SetMaxIdleConns(maxIdleConns)
	sqlDB.SetConnMaxLifetime(connMaxLifetime)

	log.Println("Database connected and migrations applied successfully.")
	return nil
}

func GetDB() (*gorm.DB, error) {
	if DbConn.db == nil {
		return nil, fmt.Errorf("database not initialized")
	}
	return DbConn.db, nil
}
