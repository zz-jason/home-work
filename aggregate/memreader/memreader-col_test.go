package memreader

import (
	"testing"
)

func TestMemReaderCol(t *testing.T) {
	reader := new(memReaderCol)
	reader.ResetPopulater(func() [][]int64 {
		return [][]int64{
			[]int64{1, 2, 3, 4, 5, 6},
			[]int64{10, 20, 30, 40, 50, 60},
		}
	})
	reader.Populate()
	for i := int64(1); i <= 6; i++ {
		cols := reader.Next(1)
		if cols[0][0] != i {
			t.Errorf("\nexpect: %v\nresult: %v\n", i, cols[0][0])
		}
		if cols[1][0] != i*10 {
			t.Errorf("\nexpect: %v\nresult: %v\n", i*10, cols[1][0])
		}
	}

	reader = new(memReaderCol)
	reader.Populate()
	for {
		cols := reader.Next(1)
		if cols == nil {
			break
		}
		if !(cols[0][0] >= 0 && cols[0][0] < 1000) {
			t.Errorf("key should in the range [0, 1000)")
		}
		if !(cols[1][0] >= -1000 && cols[1][0] < 1000) {
			t.Errorf("val should in the range [-1000, 1000)")
		}
	}
}
