package utils

import (
	// "strings"
	"unsafe"
)

//Check whether the characters exist in the array
func InListIsExist(source string, check []string) bool {
	for _, s := range check {
		if source == s {
			return true
		}
	}
	return false
}

//String to byte, only read-only
func StringToBytes(str string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&str))
	b := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&b))
}

// Byte to string, only read-only
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
