package bench

import (
	"testing"

	"github.com/bsm/vast"

	"github.com/xaionaro-go/benchmarks/internal/helpers"
)

func throughPointer() uint64 {
	for cIdx, _ := range creatives {
		creative := &creatives[cIdx]

		for mIdx, _ := range creative.Linear.MediaFiles {
			mediaFile := &creative.Linear.MediaFiles[mIdx]
			mediaFile.URI = vast.URI("/done")
		}

		creative.Sequence = 1
	}
	return 0
}

func usingCopy() uint64 {
	for cIdx, creative := range creatives {
		//isUpdated := false

		for mIdx, _ := range creative.Linear.MediaFiles {
			mediaFile := &creative.Linear.MediaFiles[mIdx]
			mediaFile.URI = vast.URI("/done")
		}

		creative.Sequence = 1
		//isUpdated = true

		//if !isUpdated {
		//	continue
		//}
		creatives[cIdx] = creative
	}
	return 0
}

func BenchmarkModifyThroughPointer(b *testing.B) {
	initCreatives()
	helpers.Benchmark(b, throughPointer)
}

func BenchmarkModifyUsingCopy(b *testing.B) {
	initCreatives()
	helpers.Benchmark(b, usingCopy)
}
