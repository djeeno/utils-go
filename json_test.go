package utils

import (
	"io"
	"strings"
	"testing"
)

var testJSON = JSON

type testJSONT struct {
	test bool
}

func TestJsonT_Unmarshal(t *testing.T) {
	t.Helper()

	// normal
	{
		testStruct := testJSONT{}
		r := strings.NewReader(`{"test":true}`)
		if err := testJSON.Unmarshal(r, &testStruct); err != nil {
			t.Errorf("TestJsonT_Unmarshal(): testJSON.Unmarshal(): err != nil: %v", err)
		}
	}
	// non-normal
	{
		testJSON.ioutilReadAllFunc = func(r io.Reader) (bytes []byte, e error) {
			return nil, ErrorDummyErrorForTest
		}
		testStruct := testJSONT{}
		r := strings.NewReader(`{"test":true}`)
		if err := testJSON.Unmarshal(r, &testStruct); err == nil {
			t.Errorf("TestJsonT_Unmarshal(): testJSON.Unmarshal(): err == nil: %v", err)
		}
	}
	{
		testStruct := testJSONT{}
		r := strings.NewReader(`{"test":true}` + `UnmarshalError`)
		if err := testJSON.Unmarshal(r, &testStruct); err == nil {
			t.Errorf("TestJsonT_Unmarshal(): testJSON.Unmarshal(): err == nil: %v", err)
		}
	}
}
