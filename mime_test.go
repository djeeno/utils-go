package utils_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/newtstat/utils-go"
)

type TestErrorReader struct {
	e string
}

func (t *TestErrorReader) Read(p []byte) (n int, err error) {
	return 0, fmt.Errorf("%s: %#v", t.e, p)
}

func TestDetectContentType(t *testing.T) {
	t.Run("DetectContentType/success", func(t *testing.T) {
		const expect = "text/html; charset=utf-8"
		actual, err := utils.MIME.DetectContentType(strings.NewReader("<!DOCTYPE html>"))
		if err != nil {
			t.Error(err)
		}
		if expect != actual {
			t.Error()
		}
	})

	t.Run("DetectContentType/error", func(t *testing.T) {
		r := &TestErrorReader{e: "error"}
		if _, err := utils.MIME.DetectContentType(r); err == nil {
			t.Error()
		}
	})
}
