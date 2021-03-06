package utils

import (
	"bytes"
	"mime/multipart"
	"testing"
)

func TestMultipart(t *testing.T) {
	_ = Multipart()
}

func TestMultipartT_CreateFormFile(t *testing.T) {
	t.Helper()
	testMultipart := _multipart

	requestBody := &bytes.Buffer{}
	multipartWriter := multipart.NewWriter(requestBody)

	if _, err := testMultipart.CreateFormFile(multipartWriter, "", ""); err != nil {
		t.Errorf("TestMultipartT_CreateFormFile(): testMultipart.CreateFormFile(): err != nil")
	}
}
