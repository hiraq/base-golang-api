package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var (
	logger              *Logger
	isDebug             bool
	enablePrintRecovery bool
	address             string
	timeout             int
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
)

func init() {
	flag.StringVar(&address, "address", addressDefault, addressUsage)
	flag.BoolVar(&isDebug, "debug", isDebugDefault, isDebugUsage)
	flag.BoolVar(&enablePrintRecovery, "enablePrintStack", enablePrintRecoveryStackDefault, enablePrintRecoveryUsage)
	flag.IntVar(&timeout, "timeout", timeoutDefault, timeoutUsage)
	flag.Parse()

	logger = logBuilder(isDebug)
	logger.Debug(fmt.Sprintf("Enable Print Recovery? %t", enablePrintRecovery))
	logger.Debug(fmt.Sprintf("Timeout: %v", timeout))
	logger.Info(fmt.Sprintf("Address: %v", address))
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
