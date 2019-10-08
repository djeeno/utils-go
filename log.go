package utils

import (
	"fmt"
	"log"
)

var Log = logT{
	LogFatallnFunc: log.Fatalln,
}

type logT struct {
	LogFatallnFunc func(v ...interface{})
}

// Printfln calls Output to print to the standard logger *with new line*.
// Arguments are handled in the manner of fmt.Print.
func (logT) Printfln(format string, v ...interface{}) {
	log.Println(fmt.Sprintf(format, v...))
}

// Fatalfln is equivalent to Printfln() followed by a call to os.Exit(1).
func (l logT) Fatalfln(format string, v ...interface{}) {
	l.LogFatallnFunc(fmt.Sprintf(format, v...))
}
