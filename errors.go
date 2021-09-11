package utils

import "errors"

var (
	// ErrShouldNotBeNil should not be nil
	ErrShouldNotBeNil = errors.New("should not be nil")

	// ErrStringIsEmpty string is empty
	ErrStringIsEmpty = errors.New("string is empty")

	// ErrSliceIsNil slice is nil
	ErrSliceIsNil = errors.New("slice is nil")

	// ErrEnvironmentVariableIsNotSetOrEmpty environment variable is not set or empty
	ErrEnvironmentVariableIsNotSetOrEmpty = errors.New("environment variable is not set or empty")
)
