package uniuri

import (
	"testing"
)

var (
	sixtyFourChars = append(testCharSet, []byte{'+', '/'}...)
	sixtyFiveChars = append(sixtyFourChars, []byte{'.'}...)
	threeChars     = []byte{'a', 'b', 'c'}
)

func BenchmarkLen16Chars65(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = generate(StdLen, sixtyFiveChars)
	}
}

func BenchmarkLen16Chars64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = generate(StdLen, sixtyFourChars)
	}
}

func BenchmarkLen16Chars62(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = generate(StdLen, testCharSet)
	}
}

func BenchmarkLen16Chars3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = generate(StdLen, threeChars)
	}
}

func BenchmarkLen1024Chars65(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = generate(1024, sixtyFiveChars)
	}
}

func BenchmarkLen1024Chars64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = generate(1024, sixtyFourChars)
	}
}

func BenchmarkLen1024Chars62(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = generate(1024, testCharSet)
	}
}

func BenchmarkLen1024Chars3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = generate(1024, threeChars)
	}
}
