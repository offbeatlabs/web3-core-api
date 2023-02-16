package db

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"time"
)

type PostgresConfig struct {
	ConnectionString   string              `mapstructure:"connection_string"`
	MaxIdleConnections int                 `mapstructure:"max_idle_connections"`
	MaxOpenConnections int                 `mapstructure:"max_open_connections"`
	LogLevel           gormLogger.LogLevel `mapstructure:"log_level"`
}

func NewPostgresDB(databaseConfig PostgresConfig) (*gorm.DB, error) {
	dbLogger := gormLogger.New(
		log.StandardLogger(),
		gormLogger.Config{
			SlowThreshold:             time.Second,             // Slow SQL threshold
			LogLevel:                  databaseConfig.LogLevel, // Log level
			IgnoreRecordNotFoundError: true,                    // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,                   // Disable color
		},
	)
	var db *gorm.DB
	var err error
	db, err = gorm.Open(postgres.Open(databaseConfig.ConnectionString), &gorm.Config{Logger: dbLogger})
	if err != nil {
		log.Fatalf("error in connecting to database %v", err)
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("error in connecting to database %v", err)
		return nil, err
	}
	sqlDB.SetMaxIdleConns(databaseConfig.MaxIdleConnections)
	sqlDB.SetMaxOpenConns(databaseConfig.MaxOpenConnections)
	return db, nil
}
