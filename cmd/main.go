package main

import (
	"context"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"time"
)

func main() {
	a := app{}
	a.initConfig()
	a.initValidator()
	a.initDB()
	a.initRepo()
	a.initExternal()
	a.initService()
	a.initTasks()
	a.initControllers()
	a.initServer()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := a.shutdown(ctx); err != nil {
		log.WithError(err).Error("error shutting the app down")
	}
}
