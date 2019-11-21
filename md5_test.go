package utils

import (
	"io"
	"log"
	"os"
	"testing"
)

func TestMD5(t *testing.T) {
	_ = MD5()
}

func TestMd5T_SumToString(t *testing.T) {
	t.Helper()
	testMD5 := _md5

	testFile, err := os.OpenFile("/etc/hosts", os.O_RDONLY, 0644)
	if err != nil {
		t.Fatalf("TestMd5T_SumToString(): os.OpenFile(): %v", err)
	}

	t.Run("normal/MD5.SumToString", func(t *testing.T) {
		md5s, err := testMD5.SumToString(testFile)
		if err != nil {
			t.Errorf("TestMd5T_SumToString(): testMD5.SumToString(): err != nil: %v", err)
		}
		log.Printf("%s: %s\n", testFile.Name(), md5s)
	})

	t.Run("non-normal/ioutil.ReadAll", func(t *testing.T) {
		testMD5.ioutilReadAllFn = func(r io.Reader) (bytes []byte, e error) {
			return nil, ErrorDummyErrorForTest
		}
		if _, err := testMD5.SumToString(testFile); err == nil {
			t.Errorf("TestMd5T_SumToString(): testMD5.SumToString(): err == nil: %v", err)
		}
	})
}
