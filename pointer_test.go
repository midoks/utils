package utils

import (
	"testing"
)

type PointerTest struct {
	I int
}

func PointerTestFuncPoint() *PointerTest {
	t := &PointerTest{}
	t.I = 1
	return t
}

func PointerTestFuncNoPoint() PointerTest {
	t := PointerTest{}
	t.I = 1
	return t
}

// go test -bench=BenchmarkPointer -benchmem -benchtime=10s
func BenchmarkPointerStruckPoint(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = PointerTestFuncPoint()
	}
	b.StopTimer()
}

func BenchmarkPointerStruckNoPoint(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = PointerTestFuncPoint()
	}
	b.StopTimer()
}
