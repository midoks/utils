package utils

import (
	"strings"
	"testing"
)

func InListIsExistTmp(source string, check []string) bool {
	for _, s := range check {
		if strings.EqualFold(source, s) {
			return true
		}
	}
	return false
}

// GOMAXPROCS=1 go test -bench=BenchmarkInListIsExistTmpEqual -benchmem -benchtime=10s
func BenchmarkInListIsExistTmpEqual(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t := []string{"t", "t1"}

		if InListIsExistTmp("t", t) {
		} else {
			b.Log("TestStringEqual ok")
		}
	}
	b.StopTimer()
}

// GOMAXPROCS=1 go test -bench=BenchmarkInListIsExist -benchmem -benchtime=10s
func BenchmarkInListIsExist(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t := []string{"t", "t1"}

		if InListIsExist("t", t) {
		} else {
			b.Log("TestStringEqual ok")
		}
	}
	b.StopTimer()
}
