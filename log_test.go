package utils

import (
	"log"
	"testing"
)

func TestLogT_Printfln(t *testing.T) {
	Log.Printfln("test: %s", "TestLogT_Printfln()")
}

func TestLogT_Fatalfln(t *testing.T) {
	Log := logT{fatallnFunc: log.Println,}
	Log.Fatalfln("test: %s", "TestLogT_Fatalfln()")
}
