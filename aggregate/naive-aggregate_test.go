package aggregate

import (
	"fmt"
	"testing"

	"github.com/zz-jason/home-work/aggregate/memreader"
)

func TestNaiveAggregate(t *testing.T) {
	opReader := memreader.Build(memreader.TypeRow)
	opReader.ResetPopulater(func() [][]int64 {
		return [][]int64{
			[]int64{10, 10},
			[]int64{1, 1},
			[]int64{2, 1},
			[]int64{2, 2},
			[]int64{2, 3},
			[]int64{2, 4},
			[]int64{2, 5},
			[]int64{2, 6},
			[]int64{2, 7},
			[]int64{3, 1},
			[]int64{3, 2},
			[]int64{3, 3},
			[]int64{3, 4},
			[]int64{3, 5},
			[]int64{3, 6},
			[]int64{3, 7},
			[]int64{3, 8},
			[]int64{3, 9},
			[]int64{3, 10},
			[]int64{2, 8},
			[]int64{2, 9},
			[]int64{2, 10},
			[]int64{1, 2},
			[]int64{1, 3},
			[]int64{1, 4},
			[]int64{1, 5},
			[]int64{1, 6},
			[]int64{1, 7},
			[]int64{1, 8},
			[]int64{1, 9},
			[]int64{4, 1},
			[]int64{4, 2},
			[]int64{4, 3},
			[]int64{4, 4},
			[]int64{4, 5},
			[]int64{4, 6},
			[]int64{4, 7},
			[]int64{4, 8},
			[]int64{4, 9},
			[]int64{4, 10},
			[]int64{1, 10},
			[]int64{0, 1},
			[]int64{0, 2},
			[]int64{0, 3},
			[]int64{0, 4},
			[]int64{0, 5},
			[]int64{0, 6},
			[]int64{0, 7},
			[]int64{0, 8},
			[]int64{0, 9},
			[]int64{0, 10},
		}
	})
	opReader.Populate()

	opAgg := new(naiveAggregate)
	opAgg.SetReader(opReader)

	result := fmt.Sprintf("%v", opAgg.GetResult())
	expect := "[[0 5.5] [1 5.5] [2 5.5] [3 5.5] [4 5.5] [10 10]]"

	if result != expect {
		t.Errorf("\nresult: %v\nexpect: %v\n", result, expect)
	}
}
