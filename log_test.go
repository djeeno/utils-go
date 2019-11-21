package utils

import (
	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
	"testing"
)

func TestLog(t *testing.T) {
	_ = Log()
}

func TestLogT_SetDebugLevel(t *testing.T) {
	t.Helper()
	testLog := _log

	testLog.SetDebugLevel(zerolog.TraceLevel)
}

func TestLogT_NoLevel(t *testing.T) {
	t.Helper()
	testLog := _log

	testLog.NoLevel("test", "sample log")
}
func TestLogT_Trace(t *testing.T) {
	t.Helper()
	testLog := _log

	testLog.Trace("test", "sample log")
}

func TestLogT_Debug(t *testing.T) {
	t.Helper()
	testLog := _log

	testLog.Debug("test", "sample log")
}

func TestLogT_Info(t *testing.T) {
	t.Helper()
	testLog := _log

	testLog.Info("test", "sample log")
}

func TestLogT_Warn(t *testing.T) {
	t.Helper()
	testLog := _log

	testLog.Warn("test", "sample log")
}

func TestLogT_Error(t *testing.T) {
	t.Helper()
	testLog := _log

	testLog.Error("test", "sample log")
}

func TestLogT_Fatal(t *testing.T) {
	t.Helper()
	testLog := _log
	testLog.zerologlogFatalFn = zlog.Log

	testLog.Fatal("test", "sample log")
}

func TestLogT_Panic(t *testing.T) {
	t.Helper()
	testLog := _log
	testLog.zerologlogPanicFn = zlog.Log

	testLog.Panic("test", "sample log")
}
