package applog

import (
	"strings"

	"github.com/charmbracelet/log"
)

type LogLevel int

const (
	DebugLevel LogLevel = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)

const defaultLevel = InfoLevel

var tradeTypeMappings = map[LogLevel]log.Level{
	DebugLevel: log.DebugLevel,
	InfoLevel:  log.InfoLevel,
	WarnLevel:  log.WarnLevel,
	ErrorLevel: log.ErrorLevel,
	FatalLevel: log.FatalLevel,
}

func SetLevel(level LogLevel) {
	log.SetLevel(tradeTypeMappings[level])
}

func LevelFromString(levelString string) LogLevel {
	level := defaultLevel
	switch strings.ToLower(levelString) {
	case "debug":
		level = DebugLevel
	case "info":
		level = InfoLevel
	case "warn":
		level = WarnLevel
	case "error":
		level = ErrorLevel
	case "fatal":
		level = FatalLevel
	}
	return level
}

func Debug(message string) {
	log.Debug(message)
}

func Info(message string) {
	log.Info(message)
}

func Warn(message string) {
	log.Warn(message)
}

func Error(message string) {
	log.Error(message)
}

func Fatal(message string) {
	log.Fatal(message)
}
