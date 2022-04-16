package utils

import (
    "runtime"
    "sync"
    "testing"
)

type BigDataStructure struct {
    ID     string
    Name   string
    Length int
}

var bigDataPool = &sync.Pool{
    New: func() interface{} {
        return new(BigDataStructure)
    },
}

// GOMAXPROCS=1 go test -bench=BenchmarkPool -benchmem -benchtime=1s
func BenchmarkPool_OutPool(b *testing.B) {
    var s *BigDataStructure
    for i := 0; i < b.N; i++ {
        for j := 0; j < 10000; j++ {
            s = &BigDataStructure{}
            s.ID = "1"
            s.Name = "Item-1"
            s.Length = j
        }
    }
}

func BenchmarkPool_Pool(b *testing.B) {
    var s *BigDataStructure
    for i := 0; i < b.N; i++ {
        for j := 0; j < 10000; j++ {
            s = bigDataPool.Get().(*BigDataStructure)
            s.ID = "1"
            s.Name = "Item-1"
            s.Length = j
            bigDataPool.Put(s)
        }
    }
}

func BenchmarkPool_PoolGC(b *testing.B) {
    var s *BigDataStructure
    for i := 0; i < b.N; i++ {
        for j := 0; j < 10000; j++ {
            s = bigDataPool.Get().(*BigDataStructure)
            s.ID = "1"
            s.Name = "Item-1"
            s.Length = j
            bigDataPool.Put(s)
        }
        runtime.GC()
    }
}
