package mut

import (
	"testing"
)

func Benchmark10WritingMut(b *testing.B) {
	for N := 0; N < b.N; N++ {
		TenPercentsWriting()
	}
}

func Benchmark50WritingMut(b *testing.B) {
	for N := 0; N < b.N; N++ {
		FiftyPercentsWriting()
	}
}

func Benchmark90WritingMut(b *testing.B) {
	for N := 0; N < b.N; N++ {
		NintyPercentsWriting()
	}
}
