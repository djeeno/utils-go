package utils

import "testing"

func TestFmtT_Printfln(t *testing.T) {
	Fmt.Printfln("test: %s", "TestFmtT_Printfln()")
}

func TestFmtT_PrintflnStderr(t *testing.T) {
	Fmt.Printfln("test: %s", "TestFmtT_PrintflnStderr()")
}
