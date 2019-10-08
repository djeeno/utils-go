package utils

import (
	"log"
	"testing"
)

func TestLogT_Printfln(t *testing.T) {
	Log.Printfln("test: TestLogT_Printfln()")
}

func TestLogT_Fatalfln(t *testing.T) {
	Log := logT{LogFatallnFunc: log.Println,}
	Log.Fatalfln("test: TestLogT_Fatalfln()")
}
