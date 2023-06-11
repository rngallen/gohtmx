package logs

import (
	"log"
	"os"
)

var InfoLogger *log.Logger
var WarningLogger *log.Logger
var ErrorLogger *log.Logger

func AppLogger() {
	file, err := os.OpenFile("logs.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err.Error())
	}
	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime)
	WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
