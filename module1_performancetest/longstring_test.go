package longstring

import "testing"

func BenchmarkThroughConcatenation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ThroughConcatenation(100)
	}
}

func BenchmarkThroughArrangment(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ThroughArrangment(100)
	}
}

func BenchmarkThroughBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ThroughBuffer(100)
	}
}
