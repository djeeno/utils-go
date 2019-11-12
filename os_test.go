package utils

import (
	"log"
	"os"
	"testing"
)

func TestOsT_GetEnvOrDefault(t *testing.T) {
	t.Helper()
	testOS := OS

	t.Run("normal/OS.GetEnvOrDefault/Get", func(t *testing.T) {
		if err := os.Setenv("TEST_ENV", "assert"); err != nil {
			t.Fatalf("TestOsT_GetEnvOrDefault(): os.Setenv(): %v", err)
		}
		if "assert" != testOS.GetEnvOrDefault("TEST_ENV", "nullString") {
			t.Errorf("TestOsT_GetEnvOrDefault(): testOS.GetEnvOrDefault()")
		}
	})

	t.Run("normal/OS.GetEnvOrDefault/Default", func(t *testing.T) {
		if err := os.Setenv("TEST_ENV", ""); err != nil {
			t.Fatalf("TestOsT_GetEnvOrDefault(): os.Setenv(): %v", err)
		}
		if "nullString" != testOS.GetEnvOrDefault("TEST_ENV", "nullString") {
			t.Errorf("TestOsT_GetEnvOrDefault(): testOS.GetEnvOrDefault()")
		}
	})
}

func TestOsT_GetEnvOrFatal(t *testing.T) {
	t.Helper()
	testOS := OS

	testOS.logFatalfFn = log.Printf
	testOS.GetEnvOrFatal("TestOsT_GetEnvOrFatal")
}

func TestOsT_Exists(t *testing.T) {
	t.Helper()
	testOS := OS

	if !testOS.Exists("/") {
		t.Errorf("TestOsT_Exists(): testOS.Exists()")
	}
}
