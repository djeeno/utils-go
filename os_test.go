package utils

import (
	"log"
	"os"
	"testing"
)

var testOS = OS

func TestOsT_GetEnvOrDefault(t *testing.T) {
	// Default
	if err := os.Setenv("TEST_ENV", ""); err != nil {
		t.Errorf("TestEnvT_GetOrDefault(): os.Setenv(): %v", err)
	}
	if "nullString" != testOS.GetEnvOrDefault("TEST_ENV", "nullString") {
		t.Errorf("TestEnvT_GetOrDefault(): testOS.GetEnvOrDefault()")
	}

	// Get
	if err := os.Setenv("TEST_ENV", "assert"); err != nil {
		t.Errorf("TestEnvT_GetOrDefault(): os.Setenv(): %v", err)
	}
	if "assert" != testOS.GetEnvOrDefault("TEST_ENV", "nullString") {
		t.Errorf("TestEnvT_GetOrDefault(): testOS.GetEnvOrDefault()")
	}
}

func TestOsT_GetEnvOrFatal(t *testing.T) {
	testOS.logFatallnFunc = log.Println
	testOS.GetEnvOrFatal("TestEnvironment_GetOrFatal()")
}

func TestOsT_Exists(t *testing.T) {
	if !testOS.Exists("/") {
		t.Errorf("TestEnvT_GetOrDefault(): testOS.Exists()")
	}
}
