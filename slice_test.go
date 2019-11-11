package utils

import "testing"

var testSlice = Slice

var testIntSliceA = []int{1, 4, 1, 4, 2, 1, 3, 5, 6}
var testIntSliceB = []int{1, 4, 1, 4, 2, 1, 3, 5, 6}
var testIntSliceC = []int{1, 4, 1, 4, 2, 1, 3, 5, 6, 2}
var testIntSliceD = []int{1, 7, 3, 2, 0, 5, 0, 8, 0}

func TestIntT_Equal(t *testing.T) {
	t.Helper()

	if !testSlice.Int.Equal(testIntSliceA, testIntSliceB) {
		t.Errorf("testIntSliceA=%#v not equal testIntSliceB=%#v\n", testIntSliceA, testIntSliceB)
	}

	if testSlice.Int.Equal(testIntSliceA, testIntSliceC) {
		t.Errorf("testIntSliceA=%#v not equal testIntSliceC=%#v\n", testIntSliceA, testIntSliceC)
	}

	if testSlice.Int.Equal(testIntSliceA, testIntSliceD) {
		t.Errorf("testIntSliceA=%#v not equal testIntSliceD=%#v\n", testIntSliceA, testIntSliceD)
	}
}
