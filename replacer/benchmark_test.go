package replacer

import (
	"testing"
	"strings"
)

func BenchmarkReplace(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			strings.Replace(strings.Replace(strings.Replace(`input string`, `one`, `two`, -1), `inp`, `nimb`, -1), `ring`, `grok`, -1)
		}
	})
}

func BenchmarkReplacer(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			replacer := strings.NewReplacer(`one`, `two`, `inp`, `nimb`, `ring`, `grok`)
			replacer.Replace(`input string`)
		}
	})
}
