package main

import "testing"

func BenchmarkDoubler(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dubblenum(100)
	}
}

func BenchmarkDoubler2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		doubler(1, 100)
	}
}
