package main

import (
	"context"
	httpClient "github.com/arhamj/go-commons/pkg/http_client"
	"github.com/go-co-op/gocron"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	defaultMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/offbeatlabs/web3-core-api/config"
	"github.com/offbeatlabs/web3-core-api/docs"
	"github.com/offbeatlabs/web3-core-api/pkg/controller"
	"github.com/offbeatlabs/web3-core-api/pkg/db"
	"github.com/offbeatlabs/web3-core-api/pkg/external"
	"github.com/offbeatlabs/web3-core-api/pkg/middleware"
	"github.com/offbeatlabs/web3-core-api/pkg/models"
	"github.com/offbeatlabs/web3-core-api/pkg/repo"
	"github.com/offbeatlabs/web3-core-api/pkg/service"
	"github.com/offbeatlabs/web3-core-api/pkg/tasks"
	"github.com/offbeatlabs/web3-core-api/pkg/util"
	log "github.com/sirupsen/logrus"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type app struct {
	validator *validator.Validate

	config config.Config

	db *gorm.DB

	scheduler *gocron.Scheduler

	tokenRepo         repo.TokenRepo
	tokenPlatformRepo repo.TokenPlatformRepo

	coingeckoExternal external.CoingeckoGateway

	tokenService service.TokenService

	tokenController controller.TokenController

	echoServer *echo.Echo
}

func (a *app) initConfig() {
	cfg, err := config.NewConfig("./config/config.json")
	if err != nil {
		log.Fatal("init config failed: ", err)
	}
	a.config = cfg
}

func (a *app) initValidator() {
	a.validator = validator.New()
}

func (a *app) initDB() {
	if a.config.FeatureFlags.UseSqlite {
		sqliteDb, err := db.NewSqliteDB(a.config.SqliteConfig)
		if err != nil {
			log.Fatal("init sqlite db failed: ", err)
		}
		log.Info("successfully initialised sqlite database")

		if a.config.HelperFlags.RunMigrations {
			err = sqliteDb.AutoMigrate(&models.Token{}, &models.TokenPlatform{})
			if err != nil {
				log.WithField("err", err).Fatal("running db migrations for sqlite db failed")
			}
			log.Info("successfully ran migrations")
		}
		a.db = sqliteDb
	} else {
		postgresDb, err := db.NewPostgresDB(a.config.PostgresConfig)
		if err != nil {
			log.Fatal("init postgres db failed: ", err)
		}
		log.Info("successfully initialised postgres database")

		if a.config.HelperFlags.RunMigrations {
			err = postgresDb.AutoMigrate(&models.Token{}, &models.TokenPlatform{})
			if err != nil {
				log.WithField("err", err).Fatal("running db migrations for postgres db failed")
			}
			log.Info("successfully ran migrations")
		}
		a.db = postgresDb
	}
	log.Info("successfully ran migrations")
}

func (a *app) initRepo() {
	a.tokenRepo = repo.NewTokenRepo(a.db)
	a.tokenPlatformRepo = repo.NewTokenPlatformRepo(a.db)
	log.Info("successfully initialised repos")
}

func (a *app) initExternal() {
	a.coingeckoExternal = external.NewCoingeckoGateway(httpClient.NewHttpClient(false))
	log.Info("successfully initialised external gateways")
}

func (a *app) initService() {
	a.tokenService = service.NewTokenService(&a.tokenRepo, &a.tokenPlatformRepo)
	log.Info("successfully initialised services")
}

func (a *app) initTasks() {
	tokenListTask := tasks.NewTokenListTask(a.coingeckoExternal, a.tokenService)
	tokenPriceTask := tasks.NewTokenPriceTask(a.coingeckoExternal, a.tokenService)

	a.scheduler = gocron.NewScheduler(time.UTC)
	if a.config.FeatureFlags.EnableTokenSync {
		_, err := a.scheduler.Every(1).Days().Do(tokenListTask.Execute)
		if err != nil {
			log.Fatal("failed to register token list task: ", err)
		}
	}
	if a.config.FeatureFlags.EnablePriceSync {
		_, err := a.scheduler.Every(5).Minutes().Do(tokenPriceTask.Execute)
		if err != nil {
			log.Fatal("failed to register token price task: ", err)
		}
	}
	log.Info("successfully initialised background tasks")
	a.scheduler.StartAsync()
}

func (a *app) initControllers() {
	a.tokenController = controller.NewTokenController(a.tokenService)
}

func (a *app) initServer() {
	e := echo.New()
	e.HideBanner = true
	e.Validator = util.NewRequestValidation(a.validator)
	e.Use(middleware.LoggingMiddleware())
	e.Use(defaultMiddleware.Recover())
	e.Use(defaultMiddleware.CORS())

	// Register routes
	e.GET("/v1/tokens", a.tokenController.GetTokenDetails)
	e.GET("/v1/tokens/multi", a.tokenController.MultiGetTokenDetails)

	a.echoServer = e

	a.registerAdminRoutes()

	log.Info("starting app server...")
	log.Fatal(e.Start(":1324"))
}

func (a *app) registerAdminRoutes() {
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Title = "Offbeat Web3 Token details and price API"
	docs.SwaggerInfo.Description = "API documentation"
	docs.SwaggerInfo.Host = a.config.ServerConfig.BaseUrlForSwagger

	// Register routes
	a.echoServer.GET("/admin/swagger/*", echoSwagger.WrapHandler)
	a.echoServer.GET("/admin/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"result": "pong",
		})
	})
	a.echoServer.GET("/admin/metrics", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"result": "pong",
		})
	})
}

func (a *app) shutdown(ctx context.Context) error {
	sqlDb, err := a.db.DB()
	if err != nil {
		return err
	}
	_ = sqlDb.Close()

	err = a.echoServer.Shutdown(ctx)
	if err != nil {
		return err
	}

	return nil
}
