package main

import (
	"ates/services/user-service/config"
	"ates/services/user-service/server"
	"ates/services/user-service/service"

	"io"
	"log"
	"os"
	"path/filepath"
)

func init() {
	config.Init()

	initLogging()
	log.SetPrefix("user-service: ")
	service.Init()
}

func main() {
	server.Start()
}

func initLogging() {
	logFolderPath := config.Get("log_path").String()
	logFilePath := filepath.Join(filepath.Clean(logFolderPath), "user-service.log")
	logFile, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	logDest := io.MultiWriter(logFile, os.Stdout)
	if err != nil {
		log.Println("could not open log folder path")
		return
	}
	log.SetOutput(logDest)
}
