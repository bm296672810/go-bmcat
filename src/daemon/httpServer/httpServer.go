package httpServer

import (
	"bmcat/src/daemon/config"
	lg "bmcat/src/daemon/log"
	"net/http"
	"time"
)

func InitHttpServer() {
	initHandler()
	server := http.Server{
		Addr:        config.GloabalConf.ListenPort,
		Handler:     &MyHandler{},
		ReadTimeout: 5 * time.Second,
	}
	lg.ILogger.Println("listen ", config.GloabalConf.ListenPort, "...")
	server.ListenAndServe()
}
