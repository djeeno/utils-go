package utils

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

var JSON = jsonT{
	ioutilReadAllFunc: ioutil.ReadAll,
}

type jsonT struct {
	ioutilReadAllFunc func(r io.Reader) ([]byte, error)
}

func (j *jsonT) Unmarshal(r io.Reader, structPointer interface{}) error {
	buf, err := j.ioutilReadAllFunc(r)
	if err != nil {
		return err
	}

	return json.Unmarshal(buf, structPointer)
}
