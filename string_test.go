package utils

import (
	"strings"
	"testing"
)

// go test -bench=BenchmarkSymbolEqual -benchtime=30s -cpuprofile mapCpu.out
// GOMAXPROCS=1 go test -bench=BenchmarkSymbolEqual -benchmem -benchtime=10s
func BenchmarkSymbolEqual(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t := "t"
		if t == "t" {
		} else {
			b.Log("TestStringEqual ok")
		}
	}
	b.StopTimer()
}

// go test -bench=BenchmarkStringsEqual -benchtime=30s -cpuprofile mapCpu.out
// GOMAXPROCS=1 go test -bench=BenchmarkStringsEqual -benchmem -benchtime=10s
func BenchmarkStringsEqual(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t := "t"
		if strings.EqualFold(t, "t") {
		} else {
			b.Log("TestStringEqual ok")
		}
	}
	b.StopTimer()
}
