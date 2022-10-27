package main

import (
	"database/sql"
	"github.com/arhamj/offbeat-api/commons/logger"
	"github.com/arhamj/offbeat-api/config"
	"github.com/arhamj/offbeat-api/pkg/db"
	"github.com/go-playground/validator"
	"log"
)

type app struct {
	logger    *logger.AppLogger
	validator *validator.Validate

	config config.Config

	db *sql.DB
}

func (a *app) initValidator() {
	a.validator = validator.New()
}

func (a *app) initLogger() {
	a.logger = logger.NewAppLogger(&a.config.LogConfig)
}

func (a *app) initConfig() {
	cfg, err := config.NewConfig("./config/config.json")
	if err != nil {
		a.logger.Panic("init config failed", err)
	}
	a.config = cfg
}

func (a *app) initDB() {
	sqliteDb, err := db.NewDB(a.config.SqliteConfig.Path)
	if err != nil {
		log.Fatal(err)
	}

	err = db.RunMigrationScripts(sqliteDb)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("successfully migrated DB..")
}
