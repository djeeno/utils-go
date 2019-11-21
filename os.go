package utils

import (
	"log"
	"os"
)

func OS() *osT {
	return &_os
}

var _os = osT{
	logFatalfFn: log.Fatalf,
}

type osT struct {
	logFatalfFn func(format string, v ...interface{})
	hostname    string
}

// GetEnvOrDefault finding a value corresponding to the passed environment variable key,
// it returns found value, otherwise it returns default value.
func (osT) GetEnvOrDefault(key, defaultValue string) (value string) {
	value = os.Getenv(key)
	if value == "" {
		log.Printf("%s is empty. use default value: %s\n", key, defaultValue)
		value = defaultValue
	}
	return value
}

// GetEnvOrFatal finding a value corresponding to the passed environment variable key,
// it returns found value, otherwise it terminates abnormally.
func (o osT) GetEnvOrFatal(key string) (value string) {
	value = os.Getenv(key)
	if value == "" {
		o.logFatalfFn("%s is empty.\n", key)
	}
	return value
}

func (o *osT) Hostname() string {
	if o.hostname == "" {
		o.hostname, _ = os.Hostname()
	}
	return o.hostname
}

// Exists returns true if anything exists in the path.
// ref. https://qiita.com/hnakamur/items/848097aad846d40ae84b
func (osT) Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
