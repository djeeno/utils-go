package utils

import (
	"log"
	"testing"
)

func TestLogT_Printfln(t *testing.T) {
	Log.Printfln("test: %s", "TestLogT_Printfln()")
}

func TestLogT_Fatalfln(t *testing.T) {
	Log := logT{logFatallnFunc: log.Println,}
	Log.Fatalfln("test: %s", "TestLogT_Fatalfln()")
}
