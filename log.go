package utils

import (
	"fmt"
	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
)

func Log() *logT {
	return &_log
}

var _log = logT{
	hostname:          OS().Hostname(),
	zerologlogFatalFn: zlog.Fatal,
	zerologlogPanicFn: zlog.Panic,
}

type logT struct {
	hostname          string
	zerologlogFatalFn func() *zerolog.Event
	zerologlogPanicFn func() *zerolog.Event
}

func (l *logT) SetDebugLevel(level zerolog.Level) {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
}

func (l *logT) NoLevel(logType, format string, a ...interface{}) {
	zlog.Log().
		Str("hostname", OS().Hostname()).
		Str("type", logType).
		Msg(fmt.Sprintf(format, a...))
}

func (l *logT) Trace(logType, format string, a ...interface{}) {
	zlog.Trace().
		Str("hostname", l.hostname).
		Str("type", logType).
		Msg(fmt.Sprintf(format, a...))
}

func (l *logT) Debug(logType, format string, a ...interface{}) {
	zlog.Debug().
		Str("hostname", l.hostname).
		Str("type", logType).
		Msg(fmt.Sprintf(format, a...))
}

func (l *logT) Info(logType, format string, a ...interface{}) {
	zlog.Info().
		Str("hostname", l.hostname).
		Str("type", logType).
		Msg(fmt.Sprintf(format, a...))
}

func (l *logT) Warn(logType, format string, a ...interface{}) {
	zlog.Warn().
		Str("hostname", l.hostname).
		Str("type", logType).
		Msg(fmt.Sprintf(format, a...))
}

func (l *logT) Error(logType, format string, a ...interface{}) {
	zlog.Error().
		Str("hostname", l.hostname).
		Str("type", logType).
		Msg(fmt.Sprintf(format, a...))
}

func (l *logT) Fatal(logType, format string, a ...interface{}) {
	l.zerologlogFatalFn().
		Str("hostname", l.hostname).
		Str("type", logType).
		Msg(fmt.Sprintf(format, a...))
}

func (l *logT) Panic(logType, format string, a ...interface{}) {
	l.zerologlogPanicFn().
		Str("hostname", l.hostname).
		Str("type", logType).
		Msg(fmt.Sprintf(format, a...))
}
