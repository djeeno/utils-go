package utils

import (
	"log"
	"testing"
)

func TestLogT_IsDebug(t *testing.T) {
	t.Helper()
	testLog := Log

	testLog.IsDebug()
}

func TestLogT_SetDebug(t *testing.T) {
	t.Helper()
	testLog := Log

	testLog.SetDebug(true)
}

func TestLogT_Debugfln(t *testing.T) {
	t.Helper()
	testLog := Log

	testLog.SetDebug(true)
	testLog.Debugfln("test: %s", "TestLogT_Debugfln()")
}

func TestLogT_Printfln(t *testing.T) {
	t.Helper()
	testLog := Log

	testLog.Printfln("test: %s", "TestLogT_Printfln()")
}

func TestLogT_Fatalfln(t *testing.T) {
	t.Helper()
	testLog := Log

	testLog.fatallnFn = log.Println
	testLog.Fatalfln("test: %s", "TestLogT_Fatalfln()")
}
