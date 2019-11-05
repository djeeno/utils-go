package utils

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
)

var Experiment = experimentT{
	HTTP: experimentHTTP,
}

type experimentT struct {
	HTTP httpT
}

var experimentHTTP = httpT{
	httpPost:                http.Post,
	multipartCreateFormFile: Multipart.CreateFormFile,
	ioCopy:                  io.Copy,
}

type httpT struct {
	httpPost                func(url, contentType string, body io.Reader) (resp *http.Response, err error)
	multipartCreateFormFile func(multipartWriter *multipart.Writer, fieldname, filename string) (io.Writer, error)
	ioCopy                  func(dst io.Writer, src io.Reader) (written int64, err error)
}

func (h *httpT) PostFormFile(url, formFieldName, formFilename string, formFile io.Reader) (*http.Response, error) {
	requestBody := &bytes.Buffer{}
	multipartWriter := multipart.NewWriter(requestBody)

	// multipart.CreateFormFile
	formFileWriter, err := h.multipartCreateFormFile(multipartWriter, formFieldName, formFilename)
	if err != nil {
		return nil, err
	}

	// io.Copy
	if _, err := h.ioCopy(formFileWriter, formFile); err != nil {
		return nil, err
	}

	contentType := multipartWriter.FormDataContentType()

	// http.Post
	return h.httpPost(url, contentType, requestBody)
}
