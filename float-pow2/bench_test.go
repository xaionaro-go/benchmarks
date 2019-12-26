package float_pow2

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/xaionaro-go/benchmarks/internal/helpers"
)

var (
	someInteger = 1
	someFloat = float64(1)
	result float64
)

func integerShift() {
	result = float64(someInteger << 8)
}

func pow2() {
	result = someFloat * math.Pow(2, 8)
}

func floatExpAdd() {
	someFloatBits := math.Float64bits(someFloat)
	result = math.Float64frombits(someFloatBits + 8 * 1<<52)
}

func BenchmarkIntegerShift(b *testing.B) {
	helpers.BenchmarkNoReturn(b, integerShift)
}
func BenchmarkPow(b *testing.B) {
	helpers.BenchmarkNoReturn(b, pow2)
}
func BenchmarkFloatExpAdd(b *testing.B) {
	helpers.BenchmarkNoReturn(b, floatExpAdd)
}

func TestIntegerShift(t *testing.T) {
	pow2()
	a := result
	integerShift()
	b := result
	floatExpAdd()
	c := result

	assert.Equal(t, a, b)
	assert.Equal(t, b, c)
}