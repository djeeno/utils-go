package utils

import (
	"io"
	"mime/multipart"
)

var Multipart = multipartT{}

type multipartT struct{}

func (multipartT) CreateFormFile(multipartWriter *multipart.Writer, fieldname, filename string) (io.Writer, error) {
	return multipartWriter.CreateFormFile(fieldname, filename)
}
