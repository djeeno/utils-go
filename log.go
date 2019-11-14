package utils

import (
	"fmt"
	"log"
)

var Log = logT{
	debug:     false,
	fatallnFn: log.Fatalln,
}

type logT struct {
	debug     bool
	fatallnFn func(v ...interface{})
}

func (l *logT) IsDebug() bool {
	return l.debug
}

func (l *logT) SetDebug(flag bool) {
	l.debug = flag
}

func (l *logT) Debugfln(format string, a ...interface{}) {
	if l.debug {
		l.Printfln(fmt.Sprintf(format, a...))
	}
}

// Printfln calls Output to print to the standard logger *with new line*.
// Arguments are handled in the manner of fmt.Print.
func (*logT) Printfln(format string, a ...interface{}) {
	log.Println(fmt.Sprintf(format, a...))
}

// Fatalfln is equivalent to Printfln() followed by a call to os.Exit(1).
func (l *logT) Fatalfln(format string, a ...interface{}) {
	l.fatallnFn(fmt.Sprintf(format, a...))
}
