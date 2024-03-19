package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Logger interface {
	Debug(msg string, args ...any)
	Error(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
}

type JLogger struct {
}

func NewJLogger() *JLogger {
	return &JLogger{}
}

func (j *JLogger) Debug(msg string, args ...any) {
	l := log.Debug()
	j.AttrsToZl(l, args...)
	l.Msg(msg)
}

func (j *JLogger) Error(msg string, args ...any) {
	l := log.Error()
	j.AttrsToZl(l, args...)
	l.Msg(msg)
}

func (j *JLogger) Info(msg string, args ...any) {
	l := log.Info()
	j.AttrsToZl(l, args...)
	l.Msg(msg)
}

func (j *JLogger) Warn(msg string, args ...any) {
	l := log.Warn()
	j.AttrsToZl(l, args...)
	l.Msg(msg)
}

func (j *JLogger) AttrsToZl(l *zerolog.Event, args ...any) {
	if len(args) == 0 {
		return
	}

	for i := 0; i < len(args); i += 2 {
		l.Any(args[i].(string), args[i+1])
	}

}
