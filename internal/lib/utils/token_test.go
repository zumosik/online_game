package utils

import "testing"

// Benchmark generateToken
func Benchmark_generateToken(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generateToken()
	}
}
