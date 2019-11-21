package utils

import "testing"

func TestFmt(t *testing.T) {
	_ = Fmt()
}

func TestFmtT_Printfln(t *testing.T) {
	t.Helper()

	_fmt.Printfln("test: %s", "TestFmtT_Printfln()")
}

func TestFmtT_PrintflnStderr(t *testing.T) {
	t.Helper()

	_fmt.PrintflnStderr("test: %s", "TestFmtT_PrintflnStderr()")
}
