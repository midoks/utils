package utils

import (
    "strings"
    "testing"
)

var strs = []string{
    "here's",
    "a",
    "some",
    "long",
    "list",
    "of",
    "strings",
    "for",
    "you",
}

func buildStrNaive() string {
    var s string
    for _, v := range strs {
        s += v
    }
    return s
}

func buildStrBuilder() string {
    b := strings.Builder{}

    // Grow the buffer to a decent length, so we don't have to continually
    // re-allocate.
    b.Grow(60)

    for _, v := range strs {
        b.WriteString(v)
    }

    return b.String()
}

// go test -bench=BenchmarkStringBuild -benchmem -benchtime=3s
func BenchmarkStringBuildNaive(b *testing.B) {
    for i := 0; i < b.N; i++ {
        _ = buildStrNaive()
    }
}
func BenchmarkStringBuildBuilder(b *testing.B) {
    for i := 0; i < b.N; i++ {
        _ = buildStrBuilder()
    }
}
