package Logger

import (
	"fmt"
	"log"
	"os"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
)

type ILogger interface {
	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string)
	SetLevel(level LogLevel)
}

type Logger struct {
	level  LogLevel
	logger *log.Logger
}

func NewLogger() ILogger {
	return &Logger{
		level:  INFO,
		logger: log.New(os.Stdout, "", log.LstdFlags),
	}
}

func (l *Logger) SetLevel(level LogLevel) {
	l.level = level
}

func (l *Logger) Debug(msg string) {
	if l.level <= DEBUG {
		l.logger.Printf("[DEBUG] %s", msg)
	}
}

func (l *Logger) Info(msg string) {
	if l.level <= INFO {
		l.logger.Printf("[INFO] %s", msg)
	}
}

func (l *Logger) Warn(msg string) {
	if l.level <= WARN {
		l.logger.Printf("[WARN] %s", msg)
	}
}

func (l *Logger) Error(msg string) {
	if l.level <= ERROR {
		l.logger.Printf("[ERROR] %s", msg)
	}
}

type MockLogger struct {
	Messages []string
}

func NewMockLogger() *MockLogger {
	return &MockLogger{
		Messages: make([]string, 0),
	}
}

func (ml *MockLogger) Debug(msg string) {
	ml.Messages = append(ml.Messages, fmt.Sprintf("[DEBUG] %s", msg))
}

func (ml *MockLogger) Info(msg string) {
	ml.Messages = append(ml.Messages, fmt.Sprintf("[INFO] %s", msg))
}

func (ml *MockLogger) Warn(msg string) {
	ml.Messages = append(ml.Messages, fmt.Sprintf("[WARN] %s", msg))
}

func (ml *MockLogger) Error(msg string) {
	ml.Messages = append(ml.Messages, fmt.Sprintf("[ERROR] %s", msg))
}

func (ml *MockLogger) SetLevel(level LogLevel) {
}



