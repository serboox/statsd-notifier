package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/serboox/statsd-notifier/src/requests"
	log "github.com/sirupsen/logrus"
)

type config struct {
	Host      string
	Port      int
	DebugMode bool
}

func main() {
	cnf := config{
		Host:      "127.0.0.1",
		Port:      8077,
		DebugMode: true,
	}

	ginMode := gin.ReleaseMode
	if cnf.DebugMode {
		ginMode = gin.DebugMode
	}

	gin.SetMode(ginMode)
	log.SetLevel(log.DebugLevel)

	router := requests.SetupRouter()

	log.Infof("HTTP-Server: Start in %s:%d ", cnf.Host, cnf.Port)
	log.Infof("HTTP-Server: PID %d", os.Getpid())

	err := http.ListenAndServe(fmt.Sprintf("%s:%d", cnf.Host, cnf.Port), router)

	if err != nil {
		log.Fatalf("Fail HTTP server start: %v", err)
	}
}
