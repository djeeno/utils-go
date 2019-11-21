package utils

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

func JSON() *jsonT {
	return &_json
}

var _json = jsonT{
	ioutilReadAllFn: ioutil.ReadAll,
}

type jsonT struct {
	ioutilReadAllFn func(r io.Reader) ([]byte, error)
}

func (j *jsonT) Unmarshal(r io.Reader, structPointer interface{}) error {
	buf, err := j.ioutilReadAllFn(r)
	if err != nil {
		return err
	}

	return json.Unmarshal(buf, structPointer)
}
