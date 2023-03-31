package main

import "testing"

func Benchmark_parseData(b *testing.B) {
	d, err := readFile()
	if err != nil {
		return
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		parseData(d)
	}
}
