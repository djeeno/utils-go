package utils

import "errors"

const (
	testTmpRootDirPath = "/tmp/.golang/github.com/djeeno/utils"
)

var (
	ErrorDummyErrorForTest = errors.New("dummy error for test")
)
