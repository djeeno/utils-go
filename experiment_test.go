package utils

import (
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"testing"
)

func TestHttpT_PostFormFile(t *testing.T) {
	t.Helper()
	testHTTP := Experiment.HTTP

	testReadCloser := ioutil.NopCloser(bytes.NewReader([]byte("testReadCloser")))

	t.Run("normal/http.Post", func(t *testing.T) {
		testHTTP.httpPost = func(url, contentType string, body io.Reader) (resp *http.Response, err error) {
			return &http.Response{
				Body: testReadCloser,
			}, nil
		} // mock http.Post
		if _, err := testHTTP.PostFormFile("", "", "", testReadCloser); err != nil {
			t.Errorf("TestHttpT_PostFormFile(): testHTTP.PostFormFile(): err != nil")
		}
	})

	t.Run("non-normal/http.Post", func(t *testing.T) {
		testHTTP.httpPost = func(url, contentType string, body io.Reader) (resp *http.Response, err error) {
			return &http.Response{
				Body: testReadCloser,
			}, ErrorDummyErrorForTest
		} // mock http.Post
		if _, err := testHTTP.PostFormFile("", "", "", testReadCloser); err == nil {
			t.Errorf("TestHttpT_PostFormFile(): testHTTP.PostFormFile(): err == nil")
		}
	})

	t.Run("non-normal/io.Copy", func(t *testing.T) {
		testHTTP.ioCopy = func(dst io.Writer, src io.Reader) (written int64, err error) {
			return 0, ErrorDummyErrorForTest
		} // mock io.Copy
		if _, err := testHTTP.PostFormFile("", "", "", testReadCloser); err == nil {
			t.Errorf("TestHttpT_PostFormFile(): testHTTP.PostFormFile(): err == nil")
		}
	})

	t.Run("non-normal/multipart.CreateFormFile", func(t *testing.T) {
		testHTTP.multipartCreateFormFile = func(multipartWriter *multipart.Writer, fieldname, filename string) (io.Writer, error) {
			return nil, ErrorDummyErrorForTest
		} // mock multipart.CreateFormFile
		if _, err := testHTTP.PostFormFile("", "", "", testReadCloser); err == nil {
			t.Errorf("TestHttpT_PostFormFile(): testHTTP.PostFormFile(): err == nil")
		}
	})
}
