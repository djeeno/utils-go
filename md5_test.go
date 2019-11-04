package utils

import (
	"io"
	"os"
	"testing"
)

func TestMd5T_SumToString(t *testing.T) {
	file, err := os.OpenFile("/etc/hosts", os.O_RDONLY, 0644)
	if err != nil {
		t.Errorf("TestMd5T_SumToString(): os.OpenFile(): %v", err)
	}

	md5s, err := MD5.SumToString(file)
	if err != nil {
		t.Errorf("TestMd5T_SumToString(): MD5.SumToString(): %v", err)
	}
	Log.Printfln("%s: %s", file.Name(), md5s)

	MD5.ioutilReadAllFunc = func(r io.Reader) (bytes []byte, e error) {
		return nil, ErrDummy
	}
	if _, err := MD5.SumToString(file); err == nil {
		t.Errorf("TestMd5T_SumToString(): MD5.SumToString(): %v", err)
	}
}
