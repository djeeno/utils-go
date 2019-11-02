package utils

import (
	"bytes"
	"mime/multipart"
	"testing"
)

func TestMultipartT_CreateFormFile(t *testing.T) {
	t.Helper()

	requestBody := &bytes.Buffer{}
	multipartWriter := multipart.NewWriter(requestBody)

	if _, err := Multipart.CreateFormFile(multipartWriter, "", ""); err != nil {
		t.Errorf("TestMultipartT_CreateFormFile(): Multipart.CreateFormFile(): err != nil")
	}

}
