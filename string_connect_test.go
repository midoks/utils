package utils

import (
    "fmt"
    "math"
    "math/rand"
    "strconv"
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

/////////////////
func logn(n, b float64) float64 {
    return math.Log(n) / math.Log(b)
}

func humanateBytes(s uint64, base float64, sizes []string) string {
    if s < 10 {
        return fmt.Sprintf("%d B", s)
    }
    e := math.Floor(logn(float64(s), base))
    suffix := sizes[int(e)]
    val := float64(s) / math.Pow(base, math.Floor(e))
    f := "%.0f"
    if val < 10 {
        f = "%.1f"
    }

    return fmt.Sprintf(f, val) + " " + suffix
}

var units = []string{"Byte", "KB", "MB", "GB", "TB", "PB", "EB"}

// FileSize calculates the file size and generate user-friendly string.
func FileSize(s int64) string {
    return humanateBytes(uint64(s), 1024, units)
}

func SizeFormat_Op(size float64) string {
    var units = []string{"Byte", "KB", "MB", "GB", "TB", "PB", "EB"}
    n := 0
    for size > 1024 {
        size /= 1024
        n += 1
    }
    return strconv.FormatFloat(size, 'f', 2, 64) + units[n]
}

func SizeFormat(size float64) string {
    var units = []string{"Byte", "KB", "MB", "GB", "TB", "PB", "EB"}
    n := 0
    for size > 1024 {
        size /= 1024
        n += 1
    }
    return fmt.Sprintf("%.2f", size) + units[n]
}

// GOMAXPROCS=1 go test -bench=BenchmarkTmpSize -benchmem -benchtime=3s
func BenchmarkTmpSizeFormatOp(b *testing.B) {
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _ = SizeFormat_Op(10000000)
    }
    b.StopTimer()
}

func BenchmarkTmpSizeFormatFloat(b *testing.B) {
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _ = float64(1000000)
    }
    b.StopTimer()
}

func BenchmarkTmpSizeFormatInt64(b *testing.B) {
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _ = int64(1000000)
    }
    b.StopTimer()
}

func BenchmarkTmpSizeFormat(b *testing.B) {
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _ = SizeFormat(10000000)
    }
    b.StopTimer()
}

func BenchmarkTmpSizeFile(b *testing.B) {
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _ = FileSize(10000000)
    }
    b.StopTimer()
}

// GOMAXPROCS=1 go test -bench=BenchmarkVarConst -benchmem -benchtime=3s
func BenchmarkVarConst_VAR(b *testing.B) {
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        var i = "2006-01-02 15:04:05"
        _ = i + "1"
    }
    b.StopTimer()
}

func BenchmarkVarConst_Const(b *testing.B) {
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        const i = "2006-01-02 15:04:05"
    }
    b.StopTimer()
}
