package main

import (
	"io/ioutil"
	"log"
	"os"
)

const (
	normalText  = "\033[0m"
	boldText    = "\033[1m"
	blackText   = "\033[30;1m"
	redText     = "\033[31;1m"
	greenText   = "\033[32;1m"
	yellowText  = "\033[33;1m"
	blueText    = "\033[34;1m"
	magentaText = "\033[35;1m"
	cyanText    = "\033[36;1m"
	whiteText   = "\033[37;1m"
)

var (
	infoTag    = greenText + "INFO " + normalText
	warningTag = yellowText + "WARNING " + normalText
	debugTag   = blueText + "DEBUG " + normalText
	fatalTag   = redText + "FATAL " + normalText

	logger *log.Logger
)

func setLogFile(filepath string) {
	var fileWriter *os.File
	var err error
	if filepath == "" {
		fileWriter = os.Stdout
	} else {
		fileWriter, err = os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("failure when trying to log to file: %s", err)
		}
	}

	logger = log.New(fileWriter, "", log.LstdFlags)
}

func turnOffLogging() {
	log.SetOutput(ioutil.Discard)
}

func info(message string)                            { logger.Print(infoTag + message) }
func warning(message string)                         { logger.Print(warningTag + message) }
func debug(message string)                           { logger.Print(debugTag + message) }
func fatal(message string)                           { logger.Fatal(fatalTag + message) }
func infof(message string, values ...interface{})    { logger.Printf(infoTag+message, values...) }
func warningf(message string, values ...interface{}) { logger.Printf(warningTag+message, values...) }
func debugf(message string, values ...interface{})   { logger.Printf(debugTag+message, values...) }
func fatalf(message string, values ...interface{})   { logger.Fatalf(fatalTag+message, values...) }
