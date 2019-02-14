package replacer

import (
	"bytes"
	"testing"
	"strings"
)

func BenchmarkReplace(b *testing.B) {
	s := `loooooooooooooooooooooooooooooooooooooooooooooong input string`
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			strings.Replace(strings.Replace(strings.Replace(s, `one`, `two`, -1), `inp`, `nimb`, -1), `ring`, `grok`, -1)
		}
	})
}

func BenchmarkReplacer(b *testing.B) {
	s := `loooooooooooooooooooooooooooooooooooooooooooooong input string`
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			replacer := strings.NewReplacer(`one`, `two`, `inp`, `nimb`, `ring`, `grok`)
			replacer.Replace(s)
		}
	})
}

func BenchmarkBytesReplace(b *testing.B) {
	bs := []byte(`loooooooooooooooooooooooooooooooooooooooooooooong input string`)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			bytes.Replace(bytes.Replace(bytes.Replace(bs, []byte(`one`), []byte(`two`), -1), []byte(`inp`), []byte(`nimb`), -1), []byte(`ring`), []byte(`grok`), -1)
		}
	})
}

