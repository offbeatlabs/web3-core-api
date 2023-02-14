package db

import (
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"time"
)

type SqliteConfig struct {
	Path     string              `mapstructure:"path"`
	LogLevel gormLogger.LogLevel `mapstructure:"log_level"`
}

func NewSqliteDB(databaseConfig SqliteConfig) (*gorm.DB, error) {
	dbLogger := gormLogger.New(
		log.StandardLogger(),
		gormLogger.Config{
			SlowThreshold:             time.Second,             // Slow SQL threshold
			LogLevel:                  databaseConfig.LogLevel, // Log level
			IgnoreRecordNotFoundError: true,                    // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,                   // Disable color
		},
	)
	db, err := gorm.Open(sqlite.Open(databaseConfig.Path), &gorm.Config{Logger: dbLogger})
	if err != nil {
		return nil, err
	}
	return db, nil
}
