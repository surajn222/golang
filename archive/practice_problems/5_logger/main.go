package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	fmt.Print("HW")
	// There are different types of logger

	// Inbuilt logger
	inbuilt_logger()

	// Inbuilt logger to file
	inbuilt_logger_to_file()

	// Inbuilt custom logger
	inbuilt_custom_logger()

	// logrus simple logger
	logrus_simple_logging()

	// logrus json logger
	logrus_json_logger()

	// logrus log level
	logrus_log_level()

	// Complex logger
}

func logrus_log_level() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.DebugLevel)
	logrus.Info("Info log")
	logrus.Warn("Warn log")
	logrus.Debug("Debug")
}

func logrus_json_logger() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.WithFields(
		logrus.Fields{
			"foo": "foo",
			"bar": "bar",
		},
	).Info("Something happened")

}

func logrus_simple_logging() {
	logrus.Println("Hello world!")
}

func inbuilt_custom_logger() {

	// Does not work
	var Infologger *log.Logger
	//var InfoLogger *log.Logger

	fmt.Println("Inbuilt custom logger")
	//file, _ := os.Create("testlog2.log")
	file, _ := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	Infologger = log.New(file, "INFO", log.Ldate|log.Ltime|log.Lshortfile)

	//Errorlogger := log.New(file, "ERROR", log.Ldate|log.Ltime|log.Lshortfile)
	//Warninglogger := log.New(file, "WARNING", log.Ldate|log.Ltime|log.Lshortfile)

	Infologger.Println("Info log")
	//Errorlogger.Println("Error log")
	//Warninglogger.Println("Warn log")
}

func inbuilt_logger() {
	fmt.Println("Inbuilt logger")
	log.Print("Inbuilt tLogger")
}

func inbuilt_logger_to_file() {

	// Create a file
	file, _ := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	// Set it in the logger
	log.SetOutput(file)

}
