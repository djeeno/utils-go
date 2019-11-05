package utils

import (
	"io"
	"os"
	"testing"
)

var testMD5 = MD5

func TestMd5T_SumToString(t *testing.T) {
	file, err := os.OpenFile("/etc/hosts", os.O_RDONLY, 0644)
	if err != nil {
		t.Errorf("TestMd5T_SumToString(): os.OpenFile(): %v", err)
	}

	md5s, err := testMD5.SumToString(file)
	if err != nil {
		t.Errorf("TestMd5T_SumToString(): testMD5.SumToString(): err != nil: %v", err)
	}
	Log.Printfln("%s: %s", file.Name(), md5s)

	testMD5.ioutilReadAllFunc = func(r io.Reader) (bytes []byte, e error) {
		return nil, ErrorDummyErrorForTest
	}
	if _, err := testMD5.SumToString(file); err == nil {
		t.Errorf("TestMd5T_SumToString(): testMD5.SumToString(): err == nil: %v", err)
	}
}
