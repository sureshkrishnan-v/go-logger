package gologger

import (
	"fmt"
	"log"
	"os"
	"time"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
)

var levelStrings = []string{"DEBUG", "INFO", "WARN", "ERROR"}

type Logger struct {
	logLevel LogLevel
	prefix   string
	file     *os.File
}

func NewLogger(prefix string, level LogLevel, filepath string) (*Logger, error) {
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	return &Logger{
		logLevel: level,
		prefix:   prefix,
		file:     file,
	}, nil
}

func (l *Logger) logInternal(level LogLevel, format string, args ...any) {
	if level < l.logLevel {
		return
	}

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	levelStr := levelStrings[level]
	message := fmt.Sprintf(format, args...)
	logLine := fmt.Sprintf("%s [%s] [%s] %s\n", timestamp, levelStr, l.prefix, message)
	fmt.Print(logLine)
	l.file.WriteString(logLine)
}

// Log functions
func (l *Logger) Debug(msg string, args ...interface{}) { l.logInternal(DEBUG, msg, args...) }
func (l *Logger) Info(msg string, args ...interface{})  { l.logInternal(INFO, msg, args...) }
func (l *Logger) Warn(msg string, args ...interface{})  { l.logInternal(WARN, msg, args...) }
func (l *Logger) Error(msg string, args ...interface{}) { l.logInternal(ERROR, msg, args...) }

// Close the logger
func (l *Logger) Close() {
	if err := l.file.Close(); err != nil {
		log.Println("Error closing log file:", err)
	}
}
