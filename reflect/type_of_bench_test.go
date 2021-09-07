package reflect

import (
	"log"
	"os"
	"reflect"
	"testing"
)

func BenchmarkReflectTypeOf(b *testing.B) {
	var iface interface{}
	iface = log.New(os.Stderr, "", log.LstdFlags)
	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		reflect.TypeOf(iface)
	}
}
