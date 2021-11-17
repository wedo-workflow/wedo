package main

import (
	"flag"
	"os"
	"os/signal"
	"time"

	"github.com/onrik/logrus/filename"
	log "github.com/sirupsen/logrus"
	"github.com/wedo-workflow/wedo/cmd/apiserver/app"
	"github.com/wedo-workflow/wedo/cmd/apiserver/config"
)

var (
	configPath = flag.String("config", "", "server configuration file path or read it from env variable CONFIG_PATH")
)

func main() {
	time.Local = time.FixedZone("UTC+8", 8*60*60)

	flag.Parse()

	if *configPath == "" {
		*configPath = os.Getenv("CONFIG_PATH")
	}
	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatalf("load config file (%s) failed: %v", *configPath, err)
	}

	logLevel, err := log.ParseLevel(cfg.LogLevel)
	if err != nil {
		logLevel = log.TraceLevel
	}
	log.SetLevel(logLevel)
	log.AddHook(filename.NewHook())
	log.SetReportCaller(true)
	log.SetFormatter(&log.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "01-02T15:04:05.999999",
	})

	server, err := app.NewAPIServer(cfg)
	if err != nil {
		log.Fatalf("api server err %v", err)
	}
	go server.Run()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt)
	sig := <-sc
	// close called.
	server.Close()
	log.Warn("system exit by received signal: ", sig.String())
}
