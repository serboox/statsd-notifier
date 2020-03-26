package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/serboox/statsd-notifier/src/configs"
	"github.com/serboox/statsd-notifier/src/version"

	"github.com/gin-gonic/gin"
	"github.com/serboox/statsd-notifier/src/requests"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Repository: ", version.Repository)
	log.Info("Release: ", version.Release)
	log.Info("Commit: ", version.Commit)
	log.Info("BuildTime: ", version.BuildTime)
	log.Infof("HTTP-Server: PID %d", os.Getpid())

	cnf := configs.NewConfig()
	cnf.ParseFromFile()

	ctx := configs.NewContext(cnf)
	defer ctx.StatsD.Close()

	ginMode := gin.ReleaseMode
	logMode := log.InfoLevel

	if cnf.Server.DebugMode {
		ginMode = gin.DebugMode
		logMode = log.DebugLevel
	}

	gin.SetMode(ginMode)
	log.SetLevel(logMode)

	router := requests.SetupRouter(ctx)

	log.Infof("HTTP-Server: Start in %s:%d ", cnf.Server.Host, cnf.Server.Port)
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", cnf.Server.Host, cnf.Server.Port), router)

	if err != nil {
		log.Fatalf("Fail HTTP server start: %v", err)
	}
}
