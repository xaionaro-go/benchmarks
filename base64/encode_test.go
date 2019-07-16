package base64

import (
	"encoding/base64"
	"testing"
)

func BenchmarkEncode(b *testing.B) {
	data := []byte{0, 1, 2, 3}

	for i := 0; i < b.N; i++ {
		base64.StdEncoding.EncodeToString(data)
	}
}
