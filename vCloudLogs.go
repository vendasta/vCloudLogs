package vCloudLogs

// This module wraps the go.log module for cloud logging.

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	Tag            string
	TraceLogger    *log.Logger
	DebugLogger    *log.Logger
	InfoLogger     *log.Logger
	WarningLogger  *log.Logger
	ErrorLogger    *log.Logger
	CriticalLogger *log.Logger
)

const loggerFlags = log.Lshortfile // date and time are handled by cloud logs when we send them there

// This handler is used for toggling on and off the Trace and Debug level logs.
// Pass debug=on, or debug=off. trace=on or trace=off as url args
// example: ?debug=on&trace=off
func LoggingOnOffHandler(w http.ResponseWriter, r *http.Request) {
	Info("Toggling Debug and Trace Logs")

	responseString := "Changing Logging: "

	// toggle trace
	trace := r.URL.Query().Get("trace")
	trace = strings.ToLower(trace)
	if trace == "on" {
		Info("Toggling Trace Logs on")
		responseString += "Trace: on "
		TraceLogger = log.New(os.Stdout, Tag+" TRACE: ", loggerFlags)
	} else if trace == "off" {
		Info("Toggling Trace Logs off")
		responseString += "Trace: off "
		TraceLogger = log.New(ioutil.Discard, Tag+" TRACE: ", loggerFlags)
	}

	// toggle debug
	debug := r.URL.Query().Get("debug")
	debug = strings.ToLower(debug)
	if debug == "on" {
		Info("Toggling Debug Logs on")
		responseString += "Debug: on "
		DebugLogger = log.New(os.Stdout, Tag+" DEBUG: ", loggerFlags)
	} else if debug == "off" {
		Info("Toggling Debug Logs off ")
		responseString += "Debug: off "
		DebugLogger = log.New(ioutil.Discard, Tag+" DEBUG: ", loggerFlags)
	}

	w.Write([]byte(responseString))
}

// These should log to os.Stdout, os.Stderr. ioutil.Discard to go to google cloud
func InitLoggers(
	tag string,
	traceHandle io.Writer,
	debugHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer,
	criticalHandle io.Writer) {

	Tag = tag
	TraceLogger = log.New(traceHandle, tag+" TRACE: ", loggerFlags)
	DebugLogger = log.New(debugHandle, tag+" DEBUG: ", loggerFlags)
	InfoLogger = log.New(infoHandle, tag+" INFO: ", loggerFlags)
	WarningLogger = log.New(warningHandle, tag+" WARNING: ", loggerFlags)
	ErrorLogger = log.New(errorHandle, tag+" ERROR: ", loggerFlags)
	CriticalLogger = log.New(errorHandle, tag+" CRITICAL: ", loggerFlags)
}

func writeLog(logger *log.Logger, format string, v ...interface{}) {
	const callDepth = 3 // get the file name of calling file.
	logger.Output(callDepth, fmt.Sprintf(format, v...))
}

func Trace(v ...interface{}) {
	writeLog(TraceLogger, "%+v", v...)
}

func Tracef(format string, v ...interface{}) {
	writeLog(TraceLogger, format, v...)
}

func Debug(v ...interface{}) {
	writeLog(DebugLogger, "%+v", v...)
}
func Debugf(format string, v ...interface{}) {
	writeLog(DebugLogger, format, v...)
}

func Info(v ...interface{}) {
	writeLog(InfoLogger, "%+v", v...)
}

func Infof(format string, v ...interface{}) {
	writeLog(InfoLogger, format, v...)
}

func Warning(v ...interface{}) {
	writeLog(WarningLogger, "%+v", v...)
}

func Warningf(format string, v ...interface{}) {
	writeLog(WarningLogger, format, v...)
}

func Error(v ...interface{}) {
	writeLog(ErrorLogger, "%+v", v...)
}

func Errorf(format string, v ...interface{}) {
	writeLog(ErrorLogger, format, v...)
}

func Critical(v ...interface{}) {
	writeLog(CriticalLogger, "%+v", v...)
}

func Criticalf(format string, v ...interface{}) {
	writeLog(CriticalLogger, format, v...)
}
