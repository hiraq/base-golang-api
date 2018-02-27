package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexsasharegan/dotenv"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var (
	logger              *Logger
	isDebug             bool
	enablePrintRecovery bool
	address             string
	timeout             int
	envFile             string
)

const (
	isDebugDefault                  = false
	isDebugUsage                    = "Enable debug mode"
	enablePrintRecoveryStackDefault = false
	enablePrintRecoveryUsage        = "Used to print all stack traces when panic happened"
	addressDefault                  = "127.0.0.1:8080"
	addressUsage                    = "Setup your running ip & port"
	timeoutDefault                  = 15
	timeoutUsage                    = "Set your write and timeout limit"
	envFileDefault                  = ""
	envFileUsage                    = "Set your dot env filpath"
)

func init() {
	flag.StringVar(&address, "address", addressDefault, addressUsage)
	flag.StringVar(&envFile, "env", envFileDefault, envFileUsage)
	flag.BoolVar(&isDebug, "debug", isDebugDefault, isDebugUsage)
	flag.BoolVar(&enablePrintRecovery, "enablePrintStack", enablePrintRecoveryStackDefault, enablePrintRecoveryUsage)
	flag.IntVar(&timeout, "timeout", timeoutDefault, timeoutUsage)
	flag.Parse()

	logger = logBuilder(isDebug)
	logger.Debug(fmt.Sprintf("Enable Print Recovery? %t", enablePrintRecovery))
	logger.Debug(fmt.Sprintf("Timeout: %v", timeout))
	logger.Info(fmt.Sprintf("Address: %v", address))

	// we need to stop all process if required env params cannot be loaded
	// only if env file is not empty
	// You can access all env variables using os.Getenv("YOURKEY")
	// ref: https://github.com/alexsasharegan/dotenv
	if envFile != "" {
		errLoadEnv := dotenv.Load(envFile)
		if errLoadEnv != nil {
			log.Fatalf("Error loading .env file: %v", errLoadEnv)
		}

		logger.Debug(fmt.Sprintf("Env file: %v", envFile))
	}
}

func main() {
	r := mux.NewRouter()

	// register routes and their controllers
	routers := RegisterRoutes(r)

	// register middlewares
	routers.Use(AccessLogMiddleware)
	routers.Use(ContentNegotiatorMiddleware)
	routers.Use(handlers.RecoveryHandler(handlers.PrintRecoveryStack(enablePrintRecovery)))

	srv := &http.Server{
		Handler:      routers,
		Addr:         address,
		WriteTimeout: time.Duration(timeout) * time.Second,
		ReadTimeout:  time.Duration(timeout) * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
