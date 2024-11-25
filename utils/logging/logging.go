// SPDX-License-Identifier: Apache-2.0
// Copyright Pionix GmbH and Contributors to EVerest

package logging

import (
	"fmt"
	"time"
)

type Logger struct{}

type logLevel int

const (
	logLevelTrace logLevel = iota
	logLevelDebug
	logLevelInfo
	logLevelError
)

func (log_level logLevel) String() string {
	switch log_level {
	case logLevelTrace:
		return "TRACE"
	case logLevelDebug:
		return "DEBUG"
	case logLevelInfo:
		return "INFO"
	case logLevelError:
		return "ERROR"
	default:
		fmt.Print("Unknown log level")
		return "UNKNOWN"
	}
}

func currentTimestamp() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func log(log_level logLevel, args ...interface{}) {
	value := fmt.Sprintln(args...)
	fmt.Printf("%s %s %s", currentTimestamp(), log_level.String(), value)
}

func logf(log_level logLevel, format string, args ...interface{}) {
	value := fmt.Sprintf(format, args...)
	fmt.Printf("%s %s %s\n", currentTimestamp(), log_level.String(), value)
}

func (logger *Logger) Trace(args ...interface{}) {
	Trace(args...)
}

func (logger *Logger) Tracef(format string, args ...interface{}) {
	Tracef(format, args...)
}

func (logger *Logger) Debug(args ...interface{}) {
	Debug(args...)
}

func (logger *Logger) Debugf(format string, args ...interface{}) {
	Debugf(format, args...)
}

func (logger *Logger) Info(args ...interface{}) {
	Info(args...)
}

func (logger *Logger) Infof(format string, args ...interface{}) {
	Infof(format, args...)
}

func (logger *Logger) Error(args ...interface{}) {
	Error(args...)
}

func (logger *Logger) Errorf(format string, args ...interface{}) {
	Errorf(format, args...)
}

func Trace(args ...interface{}) {
	log(logLevelTrace, args...)
}

func Tracef(format string, args ...interface{}) {
	logf(logLevelTrace, format, args...)
}

func Debug(args ...interface{}) {
	log(logLevelDebug, args...)
}

func Debugf(format string, args ...interface{}) {
	logf(logLevelDebug, format, args...)
}

func Info(args ...interface{}) {
	log(logLevelInfo, args...)
}

func Infof(format string, args ...interface{}) {
	logf(logLevelInfo, format, args...)
}

func Error(args ...interface{}) {
	log(logLevelError, args...)
}

func Errorf(format string, args ...interface{}) {
	logf(logLevelError, format, args...)
}
