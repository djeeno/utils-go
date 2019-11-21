package utils

import (
	"io"
	"mime/multipart"
)

func Multipart() *multipartT {
	return &_multipart
}

var _multipart = multipartT{}

type multipartT struct{}

func (multipartT) CreateFormFile(multipartWriter *multipart.Writer, fieldName, filename string) (io.Writer, error) {
	return multipartWriter.CreateFormFile(fieldName, filename)
}
