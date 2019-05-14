package logger

import (
	"fmt"
	"log"
	"strings"
)

type Logger struct {
	*log.Logger
	level string
}

func New(level string, logger *log.Logger) *Logger {
	l := &Logger{logger, level}

	return l
}
func (l *Logger) isDebug() bool {
	if strings.ToLower(l.level) == "debug" {
		return true
	}
	return false
}

func (l *Logger) Debug(v ...interface{}) {
	if l.isDebug() {
		all := append([]interface{}{"[DEBUG]"}, v...)

		l.Output(2, fmt.Sprint(all...))
	}
}
func (l *Logger) Debugln(v ...interface{}) {
	if l.isDebug() {
		all := append([]interface{}{"[DEBUG]"}, v...)

		l.Output(2, fmt.Sprintln(all...))
	}
}
func (l *Logger) Debugf(format string, v ...interface{}) {
	if l.isDebug() {
		all := append([]interface{}{"[DEBUG]"}, v...)

		l.Output(2, fmt.Sprintf(format, all...))
	}
}
