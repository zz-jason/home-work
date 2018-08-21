package memreader

import (
	"testing"
)

func TestMemReaderRow(t *testing.T) {
	reader := new(memReaderRow)
	reader.ResetPopulater(func() [][]int64 {
		return [][]int64{
			[]int64{1, 10},
			[]int64{2, 20},
			[]int64{3, 30},
			[]int64{4, 40},
			[]int64{5, 50},
			[]int64{6, 60},
		}
	})
	reader.Populate()
	for i := int64(1); i <= 6; i++ {
		rows := reader.Next(1)
		if rows[0][0] != i {
			t.Errorf("\nexpect: %v\nresult: %v\n", i, rows[0][0])
		}
		if rows[0][1] != i*10 {
			t.Errorf("\nexpect: %v\nresult: %v\n", i*10, rows[0][1])
		}
	}

	reader = new(memReaderRow)
	reader.Populate()
	for {
		rows := reader.Next(1)
		if rows == nil {
			break
		}
		if !(rows[0][0] >= 0 && rows[0][0] < 1000) {
			t.Errorf("key should in the range [0, 1000)")
		}
		if !(rows[0][1] >= -1000 && rows[0][1] < 1000) {
			t.Errorf("val should in the range [-1000, 1000)")
		}
	}
}
