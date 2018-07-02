package random

import (
	"strings"
	"testing"

	"madscientists.co/attest"
)

func TestRandomAsciiPrintable(t *testing.T) {
	test := attest.Test{t}
	tChar := AsciiPrintable()
	test.AttestGreaterThan(rune(32), tChar)
	test.AttestLessThan(rune(127), tChar)
}

func BenchmarkRandomAsciiPrintable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AsciiPrintable()
	}
}

func TestRandomAlphanumeric(t *testing.T) {
	test := attest.Test{t}
	tChar := Alphanumeric()
	test.AttestGreaterThan(-1, strings.Index(alphanumeric, string(tChar)))
}
func BenchmarkRandomAlphanumeric(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Alphanumeric()
	}
}

func TestRandomString(t *testing.T) {
	test := attest.Test{t}
	tStr := String(20)
	test.AttestEquals(20, len(tStr))
	for _, char := range tStr {
		test.AttestGreaterThan(rune(32), char) // ASCII chars 0-32 and 127 are
		test.AttestLessThan(rune(127), char)   // nonprintable characters.
	}
}

func TestRandomAlphanumericString(t *testing.T) {
	test := attest.Test{t}
	tStr := AlphanumericString(20)
	test.AttestEquals(20, len(tStr))
	for _, char := range tStr {
		test.AttestGreaterThan(-1, strings.Index(alphanumeric, string(char)))
	}
}
func BenchmarkRandomString(b *testing.B) {
	b.Run("Length 1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			String(1)
		}
	})
	b.Run("Length 5", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			String(4)
		}
	})
	b.Run("Length 20", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			String(16)
		}
	})
	b.Run("Length 50", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			String(64)
		}
	})
	b.Run("Length 100", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			String(256)
		}
	})
	b.Run("Length 500", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			String(1024)
		}
	})
}
func BenchmarkRandomAlphanumericString(b *testing.B) {
	b.Run("Length 1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			AlphanumericString(1)
		}
	})
	b.Run("Length 5", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			AlphanumericString(4)
		}
	})
	b.Run("Length 20", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			AlphanumericString(16)
		}
	})
	b.Run("Length 50", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			AlphanumericString(64)
		}
	})
	b.Run("Length 100", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			AlphanumericString(256)
		}
	})
	b.Run("Length 500", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			AlphanumericString(1024)
		}
	})
}
