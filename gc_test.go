package utils

import (
    "fmt"
    "runtime"
    "testing"
    "time"
)

func gcTime() time.Duration {
    start := time.Now()
    runtime.GC()
    return time.Since(start)
}

// go test -v ./ -run TestGCStackHeap -benchmem -benchtime=3s

// go test -v ./ -run TestGCStackHeap_Stack -benchmem -benchtime=3s
func TestGCStackHeap_Stack(t *testing.T) {
    s := make(map[int]int, 5e6)
    for i := 0; i < 5e6; i++ {
        s[i] = i
    }
    fmt.Printf("With %T, GC took %s\n", s, gcTime())
    _ = s[0]
}

// go test -v ./ -run TestGCStackHeap_Stack -benchmem -benchtime=3s
func TestGCStackHeap_Heap(t *testing.T) {
    s := make(map[int]*int, 5e6)
    for i := 0; i < 5e6; i++ {
        s[i] = &i
    }
    fmt.Printf("With %T, GC took %s\n", s, gcTime())
    _ = s[0]
}
