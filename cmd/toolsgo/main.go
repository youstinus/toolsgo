package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
	"github.com/youstinus/toolsgo/api/server/router"
	"github.com/youstinus/toolsgo/pkg/config"
	"github.com/youstinus/toolsgo/pkg/services"
)

const (
	emptyString             = ""
	titleDescription        = "example contains example http api"
	fieldJSON               = "json"
	appTitle                = "example"
	fieldHelp               = "help"
	titleHelpDescription    = "prints usage"
	fieldConf               = "conf"
	titleConfDescription    = "configuration file path. File Consists of variables (AMQP, External API keys, HTTP API data)"
	fieldVersion            = "version"
	titleVersionDescription = "prints binary version information"
	titleStartExample       = "starting example"
	titleFailedDatabase     = "failed to connect to database"
	titleDaemonExited       = "daemon exited"
)

func startExample(confFilePath string) {
	appVersion := config.NewVersion(appTitle)
	// Parse configuration file
	cfg := config.ReadYAMLfile(confFilePath)

	// Configure logging from config (Logrus)
	logLevel, err := logrus.ParseLevel(cfg.LogLevel)
	if err != nil {
		logLevel = logrus.InfoLevel
	}

	log := logrus.New()
	log.SetReportCaller(true)
	// Sets Log Formatter based on config
	switch cfg.LogFormatter {
	case fieldJSON:
		log.SetFormatter(&logrus.JSONFormatter{})
	default:
		log.SetFormatter(&logrus.TextFormatter{})
	}

	log.SetLevel(logLevel)
	log.WithField(fieldVersion, appVersion).Info(titleStartExample)

	// Configure services
	servicesIf := services.Init(&cfg.Services, log)

	// Configure controllers with services for router
	controllers := router.InitControllers(servicesIf)

	// Configure HTTP server (REST API)
	server := MakeHTTPServer(&cfg.HTTP, log, controllers)
	server.ServeHTTP()

	// Create signal channel
	signalChan := make(chan os.Signal, 1)
	// Prep signal channel
	signal.Notify(signalChan, syscall.SIGTERM, os.Interrupt)
	// Continuous loop that waits for termination signal
	for {
		switch <-signalChan {
		case os.Interrupt, syscall.SIGTERM:
			// server close not used with waitgroup.
			// Lets close server before closing other services.
			server.Close()
			log.Info(titleDaemonExited)

			return
		}
	}
}

func main() {
	var flagHelp = flag.Bool(fieldHelp, false, titleHelpDescription)

	var flagVersion = flag.Bool(fieldVersion, false, titleVersionDescription)

	var confPath = flag.String(fieldConf, emptyString, titleConfDescription)

	flag.Parse()

	if *flagHelp {
		fmt.Println(titleDescription)
		flag.PrintDefaults()
		os.Exit(0)
	}

	if *flagVersion {
		config.PrintAppVersionInfo()
		os.Exit(0)
	}

	if *confPath == emptyString {
		flag.PrintDefaults()
		os.Exit(1)
	}

	startExample(*confPath)
}
