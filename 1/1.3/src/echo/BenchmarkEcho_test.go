package echo

import (
	"testing"
)

func BenchmarkEcho1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := []string{"./echo1", "a", "b", "c"}
		echo1(input)
	}
}

func BenchmarkEcho2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := []string{"./echo2", "a", "b", "c"}
		echo2(input)
	}
}
