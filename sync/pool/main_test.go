package main

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

var slicePool = sync.Pool{
	New: func() interface{} {
		s := make([]byte, 1024)
		return &s
	},
}

func BenchmarkWithPool(b *testing.B) {
	var mstats1, mstats2 runtime.MemStats
	runtime.ReadMemStats(&mstats1)
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			//runtime.GC()
			obj := slicePool.Get().(*[]byte)
			//obj[0] = "s"
			slicePool.Put(obj)
		}
	}

	runtime.ReadMemStats(&mstats2)
	unit := fmt.Sprintf("HeapInuse%d,%d", mstats2.HeapInuse, mstats1.HeapInuse)
	b.ReportMetric(float64(mstats2.HeapInuse-mstats1.HeapInuse), unit)
}

func BenchmarkNoPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			//runtime.GC()
			obj := make([]byte, 1024)
			obj[0] = 1
		}
	}
}

type Person struct {
	Age int
}

var personPool = sync.Pool{
	New: func() interface{} { return new(Person) },
}

func Benchmark2WithoutPool(b *testing.B) {
	var p *Person
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			//runtime.GC()
			p = new(Person)
			p.Age = 23
		}
	}
}
func Benchmark2WithPool(b *testing.B) {
	var p *Person
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			//runtime.GC()
			p = personPool.Get().(*Person)
			p.Age = 23
			personPool.Put(p)
		}
	}
}
