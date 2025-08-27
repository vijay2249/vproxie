package utils

import (
	"log"
	"os"
)

var (
	WarnLogger *log.Logger
	InfoLogger *log.Logger
	DebugLogger *log.Logger
	ErrorLogger *log.Logger
)

// initiatie the looger variables
// Stdin, Stdout, and Stderr are open Files pointing to the standard input, standard output, and standard error file descriptors.
func init(){
	necessaryFlags := log.Ldate|log.LUTC|log.Lshortfile
	log.SetFlags(necessaryFlags)
	WarnLogger = log.New(os.Stdout, "WARN: ", necessaryFlags)
	InfoLogger = log.New(os.Stdout, "INFO: ", necessaryFlags)
	DebugLogger = log.New(os.Stdout, "DEBUG: ", necessaryFlags)
	ErrorLogger = log.New(os.Stdout, "ERROR: ", necessaryFlags)
}


//Get logger config details -> and get a logger variable
func GetLogger(){}

// <Time> - <endpoint> - <response> - <request id>
func PrintMessage(){}

// Print message to terminal or file as metioned in config file
func RedirectLogger(){}

//Get log level from config
func LogLevel() {}