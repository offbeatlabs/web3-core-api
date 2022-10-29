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

func (a *app) initConfig() {
	cfg, err := config.NewConfig("./config/config.json")
	if err != nil {
		log.Fatal("init config failed: ", err)
	}
	a.config = cfg
}

func (a *app) initLogger() {
	a.logger = logger.NewAppLogger(&a.config.LogConfig)
	a.logger.InitLogger()
	a.logger.Info("successfully initialised logger")
}

func (a *app) initValidator() {
	a.validator = validator.New()
}

func (a *app) initDB() {
	sqliteDb, err := db.NewDB(a.config.SqliteConfig.Path)
	if err != nil {
		a.logger.Fatal("init db failed: ", err)
	}
	a.logger.Info("successfully initialised sqlite database")

	if a.config.HelperFlags.RunMigrations {
		err = db.RunMigrationScripts(sqliteDb)
		if err != nil {
			a.logger.Fatal("running db migrations failed: ", err)
		}
	}
	a.logger.Info("successfully ran migrations")
}
