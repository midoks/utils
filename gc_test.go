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

// go test -v -run TestStackHeap
// go test -v -run TestStackHeap_GC_Stack
func TestStackHeap_GC_Stack(t *testing.T) {
    s := make(map[int]int, 5e5)
    for i := 0; i < 5e5; i++ {
        s[i] = i
    }
    fmt.Printf("With %T, GC took %s\n", s, gcTime())
    _ = s[0]
}

// go test -v -run TestStackHeap_GC_Heap
func TestStackHeap_GC_Heap(t *testing.T) {
    s := make(map[int]*int, 5e5)
    for i := 0; i < 5e5; i++ {
        s[i] = &i
    }
    fmt.Printf("With %T, GC took %s\n", s, gcTime())
    _ = s[0]
}
