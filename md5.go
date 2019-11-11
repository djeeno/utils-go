package utils

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"io/ioutil"
)

var MD5 = md5T{
	ioutilReadAllFn: ioutil.ReadAll,
}

type md5T struct {
	ioutilReadAllFn func(r io.Reader) ([]byte, error)
}

func (m *md5T) Sum(file io.Reader) (md5Hex [16]byte, err error) {
	buf, err := m.ioutilReadAllFn(file)
	if err != nil {
		return [16]byte{}, err
	}
	return md5.Sum(buf), nil
}

func (m *md5T) SumToString(file io.Reader) (md5HexString string, err error) {
	md5Hex, err := m.Sum(file)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(md5Hex[:]), nil
}
