package aggregate

import (
	"testing"

	"github.com/zz-jason/home-work/aggregate/memreader"
)

func BenchmarkNaiveAggregate(b *testing.B) {
	opReader := memreader.Build(memreader.TypeRow)
	opReader.Populate()

	opAgg := new(naiveAggregate)
	opAgg.SetReader(opReader)

	b.ResetTimer()
	for counter := 0; counter < b.N; counter++ {
		opAgg.GetResult()
	}
}
