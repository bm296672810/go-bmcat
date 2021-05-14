package main

import (
	"bmcat/src/daemon/config"
	"bmcat/src/daemon/database"
	"bmcat/src/daemon/httpServer"
)

func main() {
	config.LoadConfig("./config/config.json")
	database.InitMainSqlite(config.GloabalConf.Path)
	httpServer.InitHttpServer()
}
