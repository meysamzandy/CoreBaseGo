package persistence

import (
	"CoreBaseGo/internal/infrastructure/persistence/migrations"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"sync"
)

// DBInstance represents the singleton instance for the database connection
var DBInstance *gorm.DB
var once sync.Once

// GetInstance returns the singleton instance of the GORM DB connection
func GetInstance() (*gorm.DB, error) {
	dbType := viper.GetString("DB_TYPE")
	once.Do(func() {
		var err error
		if dbType == "mysql" {
			DBInstance, err = connectToMySqlDatabase()
		} else if dbType == "postgres" {
			DBInstance, err = connectToPostgreSQLDatabase()
		}
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}
	})
	if DBInstance == nil {
		return nil, errors.New("failed to create database instance")
	}
	return DBInstance, nil
}

// connectToPostgreSQLDatabase establishes the connection to the PostgreSQL database
func connectToPostgreSQLDatabase() (*gorm.DB, error) {
	dbHost := viper.GetString("DB_HOST")
	dbPort := viper.GetString("DB_PORT")
	dbUser := viper.GetString("DB_USER")
	dbPassword := viper.GetString("DB_PASS")
	dbName := viper.GetString("DB_NAME")
	dbSSLMode := "disable" // e.g., "disable", "require"

	// PostgreSQL DSN (Data Source Name) format
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName, dbSSLMode)

	// Create a new GORM configuration
	config := &gorm.Config{}

	// Open the connection using PostgreSQL
	db, err := gorm.Open(postgres.Open(dsn), config)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to PostgreSQL database: %w", err)
	}
	return db, nil
}

// connectToMySqlDatabase establishes the connection to the database
func connectToMySqlDatabase() (*gorm.DB, error) {
	dbHost := viper.GetString("DB_HOST")
	dbPort := viper.GetString("DB_PORT")
	dbUser := viper.GetString("DB_USER")
	dbPassword := viper.GetString("DB_PASS")
	dbName := viper.GetString("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
	// Create a new GORM configuration
	config := &gorm.Config{
		//TranslateError: true,
	}
	db, err := gorm.Open(mysql.Open(dsn), config)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	return db, nil
}

// DatabasePrepare prepares the database by migrating the provided models
func DatabasePrepare(models ...interface{}) error {
	db, err := GetInstance()
	if err != nil {
		return fmt.Errorf("failed to get database instance: %w", err)
	}

	if db == nil {
		return errors.New("database instance is nil")
	}

	// Log database connection details for debugging
	log.Printf("Database connected")

	// Auto-migrate the models
	if err := db.AutoMigrate(models...); err != nil {
		return fmt.Errorf("failed to auto-migrate models: %w", err)
	}
	return nil
}

// InitDatabase initializes the database by performing necessary preparation tasks.
func InitDatabase() {
	modelsToMigrate := migrations.Tables()

	if err := DatabasePrepare(modelsToMigrate...); err != nil {
		log.Fatalf("Error initializing the database: %v", err)
	}
}
