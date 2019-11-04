package utils

import (
	"log"
	"os"
	"testing"
)

func TestEnvT_GetOrDefault(t *testing.T) {
	// Default
	if err := os.Setenv("TEST_ENV", ""); err != nil {
		t.Errorf("test: TestEnvT_GetOrDefault(): os.Setenv(): %v", err)
	}
	if "nullString" != OS.GetEnvOrDefault("TEST_ENV", "nullString") {
		t.Errorf("test: TestEnvT_GetOrDefault(): Env.GetOrDefault()")
	}

	// Get
	if err := os.Setenv("TEST_ENV", "assert"); err != nil {
		t.Errorf("test: TestEnvT_GetOrDefault(): os.Setenv(): %v", err)
	}
	if "assert" != OS.GetEnvOrDefault("TEST_ENV", "nullString") {
		t.Errorf("test: TestEnvT_GetOrDefault(): Env.GetOrDefault()")
	}
}

func TestEnvironment_GetOrFatal(t *testing.T) {
	OS := osT{logFatallnFunc: log.Println,}
	OS.GetEnvOrFatal("test: TestEnvironment_GetOrExit()")
}

func TestOsT_Exists(t *testing.T) {
	if !OS.Exists("/") {
		t.Errorf("test: TestOsT_Exists(): OS.Exists()")
	}
}
