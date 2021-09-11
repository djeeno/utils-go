package utils

// SliceUtils is an empty structure that is prepared only for creating methods.
type SliceUtils struct{}

// Slice is an entity that allows the methods of SliceUtils to be executed from outside the package without initializing SliceUtils.
var Slice *SliceUtils

// ContainsInt returns whether or not the passed slice contains the passed value.
func (*SliceUtils) ContainsInt(slice []int, value int) bool {
	for _, elem := range slice {
		if value == elem {
			return true
		}
	}
	return false
}

// ContainsString returns whether or not the passed slice contains the passed value.
func (*SliceUtils) ContainsString(slice []string, value string) bool {
	for _, elem := range slice {
		if value == elem {
			return true
		}
	}
	return false
}
