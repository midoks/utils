package utils

import (
   "testing"
)

// GOMAXPROCS=1 go test -bench=BenchmarkStackHeap -benchmem -benchtime=3s
// GOMAXPROCS=1 go test -bench=BenchmarkStackHeap_Heap -benchmem -benchtime=3s
func BenchmarkStackHeap_Heap(b *testing.B) {
   m := make([]*string, 1000)
   for i := 0; i < b.N; i++ {
      for i := 0; i < 1000; i++ {
         s := "test"
         m[i] = &s
      }
   }
}

// GOMAXPROCS=1 go test -bench=BenchmarkStackHeap_Stack -benchmem -benchtime=3s
func BenchmarkStackHeap_Stack(b *testing.B) {
   m := make([]string, 1000)
   for i := 0; i < b.N; i++ {
      for i := 0; i < 1000; i++ {
         s := "test"
         m[i] = s
      }
   }
}
