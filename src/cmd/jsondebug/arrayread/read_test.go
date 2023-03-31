package main

import "testing"

func Benchmark_readData(b *testing.B) {
	for i := 0; i < b.N; i++ {
		readData()
	}
}
