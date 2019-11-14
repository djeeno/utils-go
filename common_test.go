package utils

import "errors"

const (
	testTmpRootDirPath = "/tmp/.go/src/github.com/djeeno/utils-go"
)

var (
	ErrorDummyErrorForTest = errors.New("dummy error for test")
)
