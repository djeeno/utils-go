package utils

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"testing"
)

func TestHttpT_PostFormFile(t *testing.T) {
	t.Helper()

	testReadCloser := ioutil.NopCloser(bytes.NewReader([]byte("testReadCloser")))

	// http.Post / err != nil
	Http.httpPost = func(url, contentType string, body io.Reader) (resp *http.Response, err error) {
		return &http.Response{
			Body: testReadCloser,
		}, nil
	} // mock http.Post
	if _, err := Http.PostFormFile("", "", "", testReadCloser); err != nil {
		t.Errorf("TestHttpT_PostFormFile(): Http.httpPost(): err != nil")
	}

	// http.Post / err == nil
	Http.httpPost = func(url, contentType string, body io.Reader) (resp *http.Response, err error) {
		return &http.Response{
			Body: testReadCloser,
		}, errors.New("http.Post is mock")
	} // mock http.Post
	if _, err := Http.PostFormFile("", "", "", testReadCloser); err == nil {
		t.Errorf("TestHttpT_PostFormFile(): Http.httpPost(): err == nil")
	}

	// io.Copy / err == nil
	Http.ioCopy = func(dst io.Writer, src io.Reader) (written int64, err error) {
		return 0, errors.New("io.Copy is mock")
	} // mock io.Copy
	if _, err := Http.PostFormFile("", "", "", testReadCloser); err == nil {
		t.Errorf("TestHttpT_PostFormFile(): Http.ioCopy(): err == nil")
	}

	// multipart.CreateFormFile / err == nil
	Http.multipartCreateFormFile = func(multipartWriter *multipart.Writer, fieldname, filename string) (io.Writer, error) {
		return nil, errors.New("multipart.CreateFormFile is mock")
	} // mock multipart.CreateFormFile
	if _, err := Http.PostFormFile("", "", "", testReadCloser); err == nil {
		t.Errorf("TestHttpT_PostFormFile(): Http.multipartCreateFormFile(): err == nil")
	}
}
