package logger

import (
	"log"
	"os"
	"path"
	"time"
)

type Logger struct{}

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

var file *os.File
var err error

func NewLogger() *Logger {
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		_ = os.Mkdir("logs", os.ModeDir)
	}

	fileName := time.Now().Format("2006-01-02") + "-logs.txt"
	path := path.Join("logs", fileName)

	file, err = os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	return &Logger{}
}

func (s *Logger) Info(msg string) {
	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	InfoLogger.Println(msg)
}

func (s *Logger) Warning(msg string) {
	WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger.Println(msg)
}

func (s *Logger) Error(msg string, err error) {
	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger.Println(msg, err)
}
