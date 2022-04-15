package utils

import (
	// "strings"
	"math/rand"
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

func RandString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := rand.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return BytesToString(bytes)
}

func RandByte(len int) []byte {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := rand.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return bytes
}

func SizeFormat(size float64) string {
	units := []string{"Byte", "KB", "MB", "GB", "TB", "PB", "EB"}
	n := 0
	for size > 1024 {
		size /= 1024
		n += 1
	}
	return strconv.FormatFloat(size, 'f', 2, 32) + " " + units[n]
}
