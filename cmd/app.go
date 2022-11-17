package main

import (
	"database/sql"
	httpClient "github.com/arhamj/offbeat-api/commons/http_client"
	"github.com/arhamj/offbeat-api/commons/logger"
	"github.com/arhamj/offbeat-api/config"
	"github.com/arhamj/offbeat-api/pkg/controller"
	"github.com/arhamj/offbeat-api/pkg/db"
	"github.com/arhamj/offbeat-api/pkg/external"
	"github.com/arhamj/offbeat-api/pkg/middleware"
	"github.com/arhamj/offbeat-api/pkg/repo"
	"github.com/arhamj/offbeat-api/pkg/service"
	"github.com/arhamj/offbeat-api/pkg/tasks"
	"github.com/go-co-op/gocron"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	defaultMiddleware "github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
	"time"
)

type app struct {
	logger    *logger.AppLogger
	validator *validator.Validate

	config config.Config

	db *sql.DB

	scheduler *gocron.Scheduler

	tokenRepo         repo.TokenRepo
	tokenPlatformRepo repo.TokenPlatformRepo

	coingeckoExternal external.CoingeckoGateway

	tokenService service.TokenService

	tokenController controller.TokenController
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
	a.db = sqliteDb
	a.logger.Info("successfully ran migrations")
}

func (a *app) initRepo() {
	a.tokenRepo = repo.NewTokenRepo(a.logger, a.db)
	a.tokenPlatformRepo = repo.NewTokenPlatformRepo(a.logger, a.db)
	a.logger.Info("successfully initialised repos")
}

func (a *app) initExternal() {
	a.coingeckoExternal = external.NewCoingeckoGateway(a.logger, httpClient.NewHttpClient(false))
	a.logger.Info("successfully initialised external gateways")
}

func (a *app) initService() {
	a.tokenService = service.NewTokenService(a.logger, &a.tokenRepo, &a.tokenPlatformRepo)
	a.logger.Info("successfully initialised services")
}

func (a *app) initTasks() {
	tokenListTask := tasks.NewTokenListTask(a.logger, a.coingeckoExternal, a.tokenService)
	tokenPriceTask := tasks.NewTokenPriceTask(a.logger, a.coingeckoExternal, a.tokenService)

	a.scheduler = gocron.NewScheduler(time.UTC)
	if a.config.FeatureFlags.EnableTokenSync {
		_, err := a.scheduler.Every(1).Days().Do(tokenListTask.Execute)
		if err != nil {
			a.logger.Fatal("failed to register token list task: ", err)
		}
	}
	if a.config.FeatureFlags.EnablePriceSync {
		_, err := a.scheduler.Every(5).Minutes().Do(tokenPriceTask.Execute)
		if err != nil {
			a.logger.Fatal("failed to register token price task: ", err)
		}
	}
	a.logger.Info("successfully initialised background tasks")
	a.scheduler.StartAsync()
}

func (a *app) initControllers() {
	a.tokenController = controller.NewTokenController(a.logger, a.tokenService)
}

func (a *app) initServer() {
	e := echo.New()
	e.Use(middleware.LoggingMiddleware(a.logger))
	e.Use(defaultMiddleware.Recover())

	// Register routes
	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"result": "pong",
		})
	})
	e.GET("/v1/token", a.tokenController.GetTokenDetails)
	a.logger.Fatal(e.Start(":1323"))
}
