package utils

import (
	"fmt"
)

var Fmt = fmtT{}

type fmtT struct{}

// Printfln formats according to a format specifier and writes to standard output *with new line*.
// It returns the number of bytes written and any write error encountered.
func (fmtT) Printfln(format string, a ...interface{}) (n int, err error) {
	return fmt.Println(fmt.Sprintf(format, a...))
}
