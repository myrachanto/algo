package main

import "testing"

// go test -bench=.
// go test -bench=. -benchmem // memeory allocations
// go test -bench=. -benchtime 10s
func BenchmarkRecursiveFibinacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RecursiveFibinacci(30)
	}
}
func BenchmarkFibinacciIterative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibinacciIterative(30)
	}
}
func BenchmarkRecursiveFibinacciEnhanced(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RecursiveFibinacciEnhanced(30)
	}
}
