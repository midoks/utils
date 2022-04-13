package utils

import (
    "fmt"
    "math/rand"
    "testing"
)

// GOMAXPROCS=1 go test -bench=BenchmarkConnect -benchmem -benchtime=3s
func BenchmarkConnectFmtSprintf(b *testing.B) {
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _ = fmt.Sprintf("connect:%s", "test")
    }
    b.StopTimer()
}

func BenchmarkConnectPlus(b *testing.B) {
    b.ResetTimer()
    for i := 0; i < b.N; i++ {

        _ = "connect:" + "test"
    }
    b.StopTimer()
}

func RandString_Op(len int) string {
    bytes := make([]byte, len)
    for i := 0; i < len; i++ {
        b := rand.Intn(26) + 65
        // fmt.Println(b)
        bytes[i] = byte(b)
    }
    return string(bytes)
}

// GOMAXPROCS=1 go test -bench=BenchmarkTmpRand -benchmem -benchtime=3s
func BenchmarkTmpRand(b *testing.B) {
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _ = RandString(10)
    }
    b.StopTimer()
}

func BenchmarkTmpRandOp(b *testing.B) {
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _ = RandString_Op(10)
    }
    b.StopTimer()
}

func BenchmarkTmpRandByte(b *testing.B) {
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _ = RandByte(10)
    }
    b.StopTimer()
}
