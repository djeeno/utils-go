package utils

func Slice() *sliceT {
	return &_slice
}

var _slice = sliceT{
	Int: intT{},
}

type sliceT struct {
	Int intT
}

type intT struct{}

func (intT) Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
