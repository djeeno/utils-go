package utils

import (
	"os"
	"strconv"
	"time"

	"golang.org/x/xerrors"
)

// EnvUtils is an empty structure that is prepared only for creating methods.
type EnvUtils struct{}

// Env is an entity that allows the methods of IntSliceUtils to be executed from outside the package without initializing IntSliceUtils.
var Env *EnvUtils

// GetOrDefaultString returns the value of the environment variable `env` if it is set, or `defaultValue` if it is not set.
func (*EnvUtils) GetOrDefaultString(env, defaultValue string) (value string) {
	valueString := os.Getenv(env)

	if valueString == "" {
		return defaultValue
	}

	return valueString
}

// GetOrDefaultBool returns the value of the environment variable `env` if it is set, or `defaultValue` if it is not set.
func (*EnvUtils) GetOrDefaultBool(env string, defaultValue bool) (value bool) {
	valueString := os.Getenv(env)

	v, err := strconv.ParseBool(valueString)
	if err != nil {
		return defaultValue
	}

	return v
}

// GetOrDefaultInt returns the value of the environment variable `env` if it is set, or `defaultValue` if it is not set.
func (*EnvUtils) GetOrDefaultInt(env string, defaultValue int) (value int) {
	valueString := os.Getenv(env)

	v, err := strconv.Atoi(valueString)
	if err != nil {
		return defaultValue
	}

	return v
}

// GetOrDefaultSecond returns the value of the environment variable `env` if it is set, or `defaultValue` if it is not set.
func (*EnvUtils) GetOrDefaultSecond(env string, defaultValue time.Duration) (value time.Duration) {
	v := Env.GetOrDefaultInt(env, -1)

	if v < 0 {
		return defaultValue
	}

	return time.Duration(v) * time.Second
}

// GetString returns the value of the environment variable `env` if it is set, or the error if it is not set.
func (*EnvUtils) GetString(env string) (value string, err error) {
	valueString := os.Getenv(env)

	if valueString == "" {
		return "", xerrors.Errorf("%s: %w", env, ErrEnvironmentVariableIsNotSetOrEmpty)
	}

	return valueString, nil
}

// GetBool returns the value of the environment variable `env` if it is set, or the error if it is not set or invalid.
func (*EnvUtils) GetBool(env string) (value bool, err error) {
	valueString := os.Getenv(env)

	if valueString == "" {
		return false, xerrors.Errorf("%s: %w", env, ErrEnvironmentVariableIsNotSetOrEmpty)
	}

	v, err := strconv.ParseBool(valueString)
	if err != nil {
		return false, xerrors.Errorf("%s: strconv.ParseBool: %w", env, err)
	}

	return v, nil
}

// GetInt returns the value of the environment variable `env` if it is set, or the error if it is not set or invalid.
func (*EnvUtils) GetInt(env string) (value int, err error) {
	valueString := os.Getenv(env)

	if valueString == "" {
		return 0, xerrors.Errorf("%s: %w", env, ErrEnvironmentVariableIsNotSetOrEmpty)
	}

	v, err := strconv.Atoi(valueString)
	if err != nil {
		return 0, xerrors.Errorf("%s: strconv.Atoi: %w", env, err)
	}

	return v, nil
}

// GetSecond returns the value of the environment variable `env` if it is set, or the error if it is not set or invalid.
func (*EnvUtils) GetSecond(env string) (value time.Duration, err error) {
	valueString := os.Getenv(env)

	if valueString == "" {
		return 0, xerrors.Errorf("%s: %w", env, ErrEnvironmentVariableIsNotSetOrEmpty)
	}

	v, err := strconv.Atoi(valueString)
	if err != nil {
		return 0, xerrors.Errorf("%s: strconv.Atoi: %w", env, err)
	}

	return time.Duration(v) * time.Second, nil
}
