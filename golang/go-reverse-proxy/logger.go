package main

import (
	"flag"
	"log"
	"os"
)

var (
	Log *log.Logger = Loggerx()
)

func Loggerx() *log.Logger {
	LOG_FILE_LOCATION := os.Getenv("LOG_FILE_LOCATION")
	if LOG_FILE_LOCATION == "" {
		LOG_FILE_LOCATION = "../logs/" + APP_NAME + ".log"
	} else {
		LOG_FILE_LOCATION = LOG_FILE_LOCATION + APP_NAME + ".log"
	}
	flag.Parse()
	if _, err := os.Stat(LOG_FILE_LOCATION); os.IsNotExist(err) {
		file, err1 := os.Create(LOG_FILE_LOCATION)
		if err1 != nil {
			panic(err1)
		}
		return log.New(file, "", log.Ldate|log.Ltime|log.Lshortfile)
	} else {
		file, err := os.OpenFile(LOG_FILE_LOCATION, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			panic(err)
		}
		return log.New(file, "", log.Ldate|log.Ltime|log.Lshortfile)
	}
}
