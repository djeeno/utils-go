package utils

import "testing"

func TestFmtT_Printfln(t *testing.T) {
	t.Helper()

	Fmt.Printfln("test: %s", "TestFmtT_Printfln()")
}

func TestFmtT_PrintflnStderr(t *testing.T) {
	t.Helper()

	Fmt.PrintflnStderr("test: %s", "TestFmtT_PrintflnStderr()")
}
