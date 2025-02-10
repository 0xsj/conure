package utils

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"sync"
)


type LogLevel int

const (
	LevelDebug LogLevel = iota
	LevelInfo
	LevelWarn
	LevelError
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
)

type Logger struct {
	mu	sync.Mutex
	out	io.Writer
	level LogLevel
	logger *log.Logger
	color bool
}

func NewLogger (out io.Writer, level LogLevel, color bool) *Logger {
	return &Logger{
		out: out,
		level: level,
		logger: log.New(out, "", log.LstdFlags|log.Lshortfile),
		color: color,
	}
}

func (l *Logger) Debug(message string, fields map[string]interface{}) {
	l.log(LevelDebug, message, fields)
}

func (l *Logger) Info(message string, fields map[string]interface{}) {
	l.log(LevelInfo, message, fields)
}

func (l *Logger) Warn(message string, fields map[string]interface{}) {
	l.log(LevelInfo, message, fields)
}

func (l *Logger) Error(message string, fields map[string]interface{}) {
	l.log(LevelInfo, message, fields)
}

func (l *Logger) log(level LogLevel, message string, fields map[string]interface{}) {
	if level < l.level {
		return
	}

	logEntry := map[string]interface{} {
		"level":	levelToString(level),
		"message": message,
	}
	for key, value := range fields {
		logEntry[key] = value
	}

	logJSON, err := json.Marshal(logEntry)
	if err != nil {
		l.logger.Printf("Failed to marshal log entry: %v", err)
		return
	}

	logLevel := levelToString(level)
	if l.color {
		logLevel = colorizeLevel(logLevel, level)
	}

	l.mu.Lock()
	defer l.mu.Unlock()
	l.logger.Printf("%s %s", logLevel, string(logJSON))
}



func levelToString(level LogLevel) string {
	switch level {
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelWarn:
		return "WARN"
	case LevelError:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}


func colorizeLevel(level string, logLevel LogLevel) string {
	switch logLevel {
	case LevelDebug:
		return colorCyan + level + colorReset
	case LevelInfo:
		return colorGreen + level + colorReset
	case LevelWarn:
		return colorYellow + level + colorReset
	case LevelError:
		return colorRed + level + colorReset
	default:
		return level
	}
}

var DefaultLogger = NewLogger(os.Stdout, LevelInfo, true)