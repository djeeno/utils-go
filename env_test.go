package utils_test

import (
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/newtstat/utils-go"
)

func TestGetOrDefaultString(t *testing.T) {
	testEnvKey := "UTILS_GO_ENV_TEST"
	backupTestEnvValue := os.Getenv(testEnvKey)
	defer os.Setenv(testEnvKey, backupTestEnvValue)

	testEnvDefaultValue := "defaultValue"
	testEnvValue := "value"

	t.Run("GetOrDefaultString/default", func(t *testing.T) {
		os.Setenv(testEnvKey, "")
		actual := utils.Env.GetOrDefaultString(testEnvKey, testEnvDefaultValue)
		if actual != testEnvDefaultValue {
			t.Fail()
		}
	})

	t.Run("GetOrDefaultString/get", func(t *testing.T) {
		os.Setenv(testEnvKey, testEnvValue)
		actual := utils.Env.GetOrDefaultString(testEnvKey, testEnvDefaultValue)
		if actual != testEnvValue {
			t.Fail()
		}
	})
}

func TestGetOrDefaultBool(t *testing.T) {
	testEnvKey := "UTILS_GO_ENV_TEST"
	backupTestEnvValue := os.Getenv(testEnvKey)
	defer os.Setenv(testEnvKey, backupTestEnvValue)

	testEnvDefaultValue := false
	testEnvValue := true

	t.Run("GetOrDefaultBool/default", func(t *testing.T) {
		os.Setenv(testEnvKey, "")
		actual := utils.Env.GetOrDefaultBool(testEnvKey, testEnvDefaultValue)
		if actual != testEnvDefaultValue {
			t.Fail()
		}
	})

	t.Run("GetOrDefaultBool/get", func(t *testing.T) {
		os.Setenv(testEnvKey, strconv.FormatBool(testEnvValue))
		actual := utils.Env.GetOrDefaultBool(testEnvKey, testEnvDefaultValue)
		if actual != testEnvValue {
			t.Fail()
		}
	})
}

func TestGetOrDefaultInt(t *testing.T) {
	testEnvKey := "UTILS_GO_ENV_TEST"
	backupTestEnvValue := os.Getenv(testEnvKey)
	defer os.Setenv(testEnvKey, backupTestEnvValue)

	testEnvDefaultValue := 1
	testEnvValue := 2

	t.Run("GetOrDefaultInt/default", func(t *testing.T) {
		os.Setenv(testEnvKey, "")
		actual := utils.Env.GetOrDefaultInt(testEnvKey, testEnvDefaultValue)
		if actual != testEnvDefaultValue {
			t.Fail()
		}
	})

	t.Run("GetOrDefaultInt/get", func(t *testing.T) {
		os.Setenv(testEnvKey, strconv.Itoa(testEnvValue))
		actual := utils.Env.GetOrDefaultInt(testEnvKey, testEnvDefaultValue)
		if actual != testEnvValue {
			t.Fail()
		}
	})
}

func TestGetOrDefaultSecond(t *testing.T) {
	testEnvKey := "UTILS_GO_ENV_TEST"
	backupTestEnvValue := os.Getenv(testEnvKey)
	defer os.Setenv(testEnvKey, backupTestEnvValue)

	testEnvDefaultValue := 1 * time.Second
	testEnvValue := 2 * time.Second

	t.Run("GetOrDefaultSecond/default", func(t *testing.T) {
		os.Setenv(testEnvKey, "")
		actual := utils.Env.GetOrDefaultSecond(testEnvKey, testEnvDefaultValue)
		if actual != testEnvDefaultValue {
			t.Fail()
		}
	})

	t.Run("GetOrDefaultSecond/get", func(t *testing.T) {
		os.Setenv(testEnvKey, "2")
		actual := utils.Env.GetOrDefaultSecond(testEnvKey, testEnvDefaultValue)
		if actual != testEnvValue {
			t.Fail()
		}
	})
}

func TestGetString(t *testing.T) {
	testEnvKey := "UTILS_GO_ENV_TEST"
	backupTestEnvValue := os.Getenv(testEnvKey)
	defer os.Setenv(testEnvKey, backupTestEnvValue)

	testEnvSuccessValue := "value"
	testEnvErrorValue := ""

	t.Run("GetString/success", func(t *testing.T) {
		os.Setenv(testEnvKey, testEnvSuccessValue)
		actual, err := utils.Env.GetString(testEnvKey)
		if err != nil {
			t.Fail()
		}
		if actual != testEnvSuccessValue {
			t.Fail()
		}
	})

	t.Run("GetString/error", func(t *testing.T) {
		os.Setenv(testEnvKey, "")
		actual, err := utils.Env.GetString(testEnvKey)
		if err == nil {
			t.Fail()
		}
		if actual != testEnvErrorValue {
			t.Fail()
		}
	})
}

func TestGetBool(t *testing.T) {
	testEnvKey := "UTILS_GO_ENV_TEST"
	backupTestEnvValue := os.Getenv(testEnvKey)
	defer os.Setenv(testEnvKey, backupTestEnvValue)

	testEnvSuccessValue := true
	testEnvErrorValue := false

	t.Run("GetBool/success", func(t *testing.T) {
		os.Setenv(testEnvKey, strconv.FormatBool(testEnvSuccessValue))
		actual, err := utils.Env.GetBool(testEnvKey)
		if err != nil {
			t.Fail()
		}
		if actual != testEnvSuccessValue {
			t.Fail()
		}
	})

	t.Run("GetBool/empty", func(t *testing.T) {
		os.Setenv(testEnvKey, "")
		actual, err := utils.Env.GetBool(testEnvKey)
		if err == nil {
			t.Fail()
		}
		if actual != testEnvErrorValue {
			t.Fail()
		}
	})

	t.Run("GetBool/invalid", func(t *testing.T) {
		os.Setenv(testEnvKey, "invalid")
		actual, err := utils.Env.GetBool(testEnvKey)
		if err == nil {
			t.Fail()
		}
		if actual != testEnvErrorValue {
			t.Fail()
		}
	})
}

func TestGetInt(t *testing.T) {
	testEnvKey := "UTILS_GO_ENV_TEST"
	backupTestEnvValue := os.Getenv(testEnvKey)
	defer os.Setenv(testEnvKey, backupTestEnvValue)

	testEnvSuccessValue := 1
	testEnvErrorValue := 0

	t.Run("GetInt/success", func(t *testing.T) {
		os.Setenv(testEnvKey, strconv.Itoa(testEnvSuccessValue))
		actual, err := utils.Env.GetInt(testEnvKey)
		if err != nil {
			t.Fail()
		}
		if actual != testEnvSuccessValue {
			t.Fail()
		}
	})

	t.Run("GetInt/empty", func(t *testing.T) {
		os.Setenv(testEnvKey, "")
		actual, err := utils.Env.GetInt(testEnvKey)
		if err == nil {
			t.Fail()
		}
		if actual != testEnvErrorValue {
			t.Fail()
		}
	})

	t.Run("GetInt/invalid", func(t *testing.T) {
		os.Setenv(testEnvKey, "invalid")
		actual, err := utils.Env.GetInt(testEnvKey)
		if err == nil {
			t.Fail()
		}
		if actual != testEnvErrorValue {
			t.Fail()
		}
	})
}

func TestGetSecond(t *testing.T) {
	testEnvKey := "UTILS_GO_ENV_TEST"
	backupTestEnvValue := os.Getenv(testEnvKey)
	defer os.Setenv(testEnvKey, backupTestEnvValue)

	testEnvSuccessValue := 1 * time.Second
	testEnvErrorValue := 0 * time.Second

	t.Run("GetSecond/success", func(t *testing.T) {
		os.Setenv(testEnvKey, "1")
		actual, err := utils.Env.GetSecond(testEnvKey)
		if err != nil {
			t.Fail()
		}
		if actual != testEnvSuccessValue {
			t.Fail()
		}
	})

	t.Run("GetSecond/empty", func(t *testing.T) {
		os.Setenv(testEnvKey, "")
		actual, err := utils.Env.GetSecond(testEnvKey)
		if err == nil {
			t.Fail()
		}
		if actual != testEnvErrorValue {
			t.Fail()
		}
	})

	t.Run("GetSecond/invalid", func(t *testing.T) {
		os.Setenv(testEnvKey, "invalid")
		actual, err := utils.Env.GetSecond(testEnvKey)
		if err == nil {
			t.Fail()
		}
		if actual != testEnvErrorValue {
			t.Fail()
		}
	})
}
