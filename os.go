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
		o.logFatallnFunc(key, "is empty.")
	}
	return value
}

// Exists returns true if anything exists in the path.
// ref. https://qiita.com/hnakamur/items/848097aad846d40ae84b
func (osT) Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
