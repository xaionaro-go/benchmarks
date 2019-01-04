package utf8

import (
	"testing"
	"unicode/utf8"
)

func BenchmarkValidString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		utf8.ValidString("someString")
	}
}
