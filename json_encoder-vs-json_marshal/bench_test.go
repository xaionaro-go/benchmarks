package jsonEncoderVsJSONMarshal

import (
	"bytes"
	"encoding/json"
	"sync"
	"testing"
)

var (
	data4 []int
	data256 []int
	data65536 []int
)

func init() {
	data4 = make([]int, 4)
	data256 = make([]int, 256)
	data65536 = make([]int, 65536)
}

func BenchmarkJsonEncoder4(b *testing.B) {
	for i:=0; i<b.N; i++ {
		var buf bytes.Buffer
		_ = json.NewEncoder(&buf).Encode(data4)
	}
}

func BenchmarkJsonEncoder256(b *testing.B) {
	for i:=0; i<b.N; i++ {
		var buf bytes.Buffer
		_ = json.NewEncoder(&buf).Encode(data256)
	}
}

func BenchmarkJsonEncoder65536(b *testing.B) {
	for i:=0; i<b.N; i++ {
		var buf bytes.Buffer
		_ = json.NewEncoder(&buf).Encode(data65536)
	}
}

var (
	bufPool = sync.Pool{
		New: func() interface{} {
			return &bytes.Buffer{}
		},
	}
)

func BenchmarkJsonEncoderWithBufferReuse4(b *testing.B) {
	for i:=0; i<b.N; i++ {
		buf := bufPool.Get().(*bytes.Buffer)
		_ = json.NewEncoder(buf).Encode(data4)
		buf.Reset()
		bufPool.Put(buf)
	}
}

func BenchmarkJsonEncoderWithBufferReuse256(b *testing.B) {
	for i:=0; i<b.N; i++ {
		buf := bufPool.Get().(*bytes.Buffer)
		_ = json.NewEncoder(buf).Encode(data256)
		buf.Reset()
		bufPool.Put(buf)
	}
}

func BenchmarkJsonEncoderWithBufferReuse65536(b *testing.B) {
	for i:=0; i<b.N; i++ {
		buf := bufPool.Get().(*bytes.Buffer)
		_ = json.NewEncoder(buf).Encode(data65536)
		buf.Reset()
		bufPool.Put(buf)
	}
}

func BenchmarkJsonMarshal4(b *testing.B) {
	for i:=0; i<b.N; i++ {
		_, _ = json.Marshal(data4)
	}
}

func BenchmarkJsonMarshal256(b *testing.B) {
	for i:=0; i<b.N; i++ {
		_, _ = json.Marshal(data256)
	}
}

func BenchmarkJsonMarshal65536(b *testing.B) {
	for i:=0; i<b.N; i++ {
		_, _ = json.Marshal(data65536)
	}
}
