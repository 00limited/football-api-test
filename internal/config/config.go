package config

import (
	"fmt"
	"os"

	"github.com/00limited/football-api/internal/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Config struct {
	Port       string
	DBDriver   string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
	DBDSN      string
	JWTSecret  string
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	cfg := &Config{
		Port:       getEnv("PORT", "8080"),
		DBDriver:   getEnv("DB_DRIVER", "postgres"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "postgres"),
		DBName:     getEnv("DB_NAME", "football_api"),
		DBSSLMode:  getEnv("DB_SSLMODE", "disable"),
		DBDSN:      os.Getenv("DB_DSN"),
		JWTSecret:  os.Getenv("JWT_SECRET"),
	}

	if cfg.JWTSecret == "" {
		return nil, fmt.Errorf("JWT_SECRET is required")
	}

	return cfg, nil
}

func OpenDatabase(cfg *Config) (*gorm.DB, error) {
	var (
		db  *gorm.DB
		err error
	)

	switch cfg.DBDriver {
	case "sqlite":
		dsn := cfg.DBDSN
		if dsn == "" {
			dsn = "football_api.db"
		}
		db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	default:
		dsn := cfg.DBDSN
		if dsn == "" {
			dsn = "host=" + cfg.DBHost +
				" port=" + cfg.DBPort +
				" user=" + cfg.DBUser +
				" password=" + cfg.DBPassword +
				" dbname=" + cfg.DBName +
				" sslmode=" + cfg.DBSSLMode
		}
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	}
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(
		&models.Admin{},
		&models.Team{},
		&models.Player{},
		&models.Match{},
		&models.MatchResult{},
		&models.Goal{},
	); err != nil {
		return nil, err
	}

	return db, nil
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
