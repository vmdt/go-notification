package logger

import (
	log "github.com/sirupsen/logrus"
)

type ILogger interface {
	getLevel() log.Level
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Panic(args ...interface{})
	Panicf(format string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
}

var (
	Logger ILogger
	loggerLvMap = map[string]log.Level{
		"debug": log.DebugLevel,
		"info": log.InfoLevel,
		"warn": log.WarnLevel,
		"error": log.ErrorLevel,
		"panic": log.PanicLevel,
		"fatal": log.FatalLevel,
	}
)

type LoggerConfig struct {
	LogLevel string `mapstructure:"level"`
}

type appLogger struct {
	level string
	logger *log.Logger
}

func InitLogger(cfg *LoggerConfig) ILogger {
	l := &appLogger{ level: cfg.LogLevel }
	l.logger = log.New()

	logLogger := l.getLevel()
	log.SetFormatter(&log.JSONFormatter{})

	log.SetLevel(logLogger)

	return l
}

func (l *appLogger) getLevel() log.Level {
	return loggerLvMap[l.level]
}

func (l *appLogger) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

func (l *appLogger) Debugf(format string, args ...interface{}) {
	l.logger.Debugf(format, args...)
}

func (l *appLogger) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l *appLogger) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

func (l *appLogger) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}

func (l *appLogger) Warnf(format string, args ...interface{}) {
	l.logger.Warnf(format, args...)
}

func (l *appLogger) Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l *appLogger) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

func (l *appLogger) Panic(args ...interface{}) {
	l.logger.Panic(args...)
}

func (l *appLogger) Panicf(format string, args ...interface{}) {
	l.logger.Panicf(format, args...)
}

func (l *appLogger) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}

func (l *appLogger) Fatalf(format string, args ...interface{}) {
	l.logger.Fatalf(format, args...)
}
