package utils

import (
	"log"
	"os"
)

var OS = osT{
	logFatallnFunc: log.Fatalln,
}

type osT struct {
	logFatallnFunc func(...interface{})
}

// GetEnvOrDefault finding a value corresponding to the passed environment variable key,
// it returns found value, otherwise it returns default value.
func (osT) GetEnvOrDefault(key, defaultValue string) (value string) {
	value = os.Getenv(key)
	if value == "" {
		log.Println(key, "is empty. use default value:", defaultValue)
		value = defaultValue
	}
	return value
}

// GetEnvOrFatal finding a value corresponding to the passed environment variable key,
// it returns found value, otherwise it terminates abnormally.
func (o osT) GetEnvOrFatal(key string) (value string) {
	value = os.Getenv(key)
	if value == "" {
		o.logFatallnFunc("%s is empty.", key)
	}
	return value
}
