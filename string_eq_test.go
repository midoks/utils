package utils

import (
	"strings"
	"testing"
)

// go test -bench=BenchmarkEqualSymbol -benchtime=30s -cpuprofile mapCpu.out
// GOMAXPROCS=1 go test -bench=BenchmarkEqualSymbol -benchmem -benchtime=10s
func BenchmarkEqualSymbol(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t := "t"
		if t == "t" {
		} else {
			b.Error("TestStringEqual ok")
		}
	}
	b.StopTimer()
}

// go test -bench=BenchmarkEqualSymbolToLower -benchtime=30s -cpuprofile mapCpu.out
// GOMAXPROCS=1 go test -bench=BenchmarkEqualSymbolToLower -benchmem -benchtime=10s
func BenchmarkEqualSymbolToLower(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t := "t"
		if t == strings.ToLower("T") {
		} else {
			b.Error("TestStringEqual ok")
		}
	}
	b.StopTimer()
}

// go test -bench=BenchmarkEqualSymbolNEQ -benchtime=30s -cpuprofile mapCpu.out
// GOMAXPROCS=1 go test -bench=BenchmarkEqualSymbolNEQ -benchmem -benchtime=10s
func BenchmarkEqualSymbolNEQ(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t := "t"
		if t == "TT" {
		}
	}
	b.StopTimer()
}

// go test -bench=BenchmarkEqualSymbolNEQ2 -benchtime=30s -cpuprofile mapCpu.out
// GOMAXPROCS=1 go test -bench=BenchmarkEqualSymbolNEQ2 -benchmem -benchtime=10s
func BenchmarkEqualSymbolNEQ2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t := "t"
		if t == "T" {
		}
	}
	b.StopTimer()
}

// go test -bench=BenchmarkEqualStrings -benchtime=30s -cpuprofile mapCpu.out
// GOMAXPROCS=1 go test -bench=BenchmarkEqualStrings -benchmem -benchtime=10s
func BenchmarkEqualStrings(b *testing.B) {
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
