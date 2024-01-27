package logger

import (
	"log"
	"os"
)

type Logger struct{}

var logger *log.Logger

func Init(prefix string) {
	logger = log.New(os.Stdout, prefix, log.Ldate|log.Ltime|log.Lshortfile)
}

func Info(message string) {
	log.Println("[ℹ️ INFO] -", message)
}

func Error(message string) {
	log.Println("[❌ ERROR] -", message)
}

func Warning(message string) {
	log.Println("[⚠️ WARNING] -", message)
}

func Fatal(message string) {
	log.Fatal("[🚨 FALTAL] - " + message)
}
