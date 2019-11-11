package utils

import (
	"log"
	"testing"
)

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
