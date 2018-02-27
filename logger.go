package main

import (
	"bytes"
	"fmt"
	"log"
)

// Logger used to construct logger proccess
type Logger struct {
	IsDebug bool
}

// Info used to log all INFO's messages
func (l *Logger) Info(msg ...interface{}) {
	logger, buf := createNew()

	msglog := fmt.Sprintf("INFO: %s", msg)
	logger.Print(msglog)
	fmt.Println(buf)
}

// Error used to catch any error messages
func (l *Logger) Error(msg ...interface{}) {
	logger, buf := createNew()

	msglog := fmt.Sprintf("Error: %s", msg)
	logger.Print(msglog)
	fmt.Println(buf)
}

// Debug used to log all DEBUG's message.
// This method should be active only if current process
// started with -debug (true).
func (l *Logger) Debug(msg ...interface{}) {
	if l.IsDebug {
		logger, buf := createNew()

		msglog := fmt.Sprintf("DEBUG: %s", msg)
		logger.Print(msglog)
		fmt.Println(buf)
	}
}

func logBuilder(isDebug bool) *Logger {
	builder := new(Logger)
	builder.IsDebug = isDebug
	return builder
}

func createNew() (*log.Logger, *bytes.Buffer) {
	var buf bytes.Buffer
	logger := log.New(&buf, "Log: ", log.Ldate|log.Ltime|log.LUTC)
	return logger, &buf
}
