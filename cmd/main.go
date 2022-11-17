package main

import (
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	a := app{}
	a.initConfig()
	a.initLogger()
	a.initValidator()
	a.initDB()
	a.initRepo()
	a.initExternal()
	a.initService()
	a.initTasks()
	a.initControllers()
	a.initServer()
}
