package database

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Connection *gorm.DB

func loadDataSourceName() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))
}

func ConnectDB() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(loadDataSourceName()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get generic database object: %w", err)
	}

	maxIdle := getEnvAsInt("DB_MAX_IDLE_CONNS", 10)
	maxOpen := getEnvAsInt("DB_MAX_OPEN_CONNS", 100)
	lifetime := getEnvAsDuration("DB_CONN_MAX_LIFETIME", time.Hour)

	sqlDB.SetMaxIdleConns(maxIdle)
	sqlDB.SetMaxOpenConns(maxOpen)
	sqlDB.SetConnMaxLifetime(lifetime)
	
	Connection = db
	return Connection, nil
}

func getEnvAsInt(name string, defaultVal int) int {
	if val, ok := os.LookupEnv(name); ok {
		var v int
		fmt.Sscanf(val, "%d", &v)
		return v
	}
	return defaultVal
}

func getEnvAsDuration(name string, defaultVal time.Duration) time.Duration {
	if val, ok := os.LookupEnv(name); ok {
		if d, err := time.ParseDuration(val); err == nil {
			return d
		}
	}
	return defaultVal
}