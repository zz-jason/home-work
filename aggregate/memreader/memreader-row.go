package memreader

import (
	"math/rand"
	"time"
)

type memReaderRow struct {
	rows      [][]int64
	populater populater

	numReturnedRows int
}

func (r *memReaderRow) ResetPopulater(other populater) {
	r.populater = other
}

func (r *memReaderRow) Populate() {
	if r.populater != nil {
		r.rows = r.populater()
		return
	}

	numRows, numCols := 100000000, 2
	r.rows = make([][]int64, numRows)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < numRows; i++ {
		r.rows[i] = make([]int64, numCols)
		r.rows[i][0] = rand.Int63n(1000)
		r.rows[i][1] = rand.Int63n(2000) - 1000
	}
}

func (r *memReaderRow) Next(numMaxRows int) [][]int64 {
	numRows := numMaxRows
	numRemainedRows := len(r.rows) - r.numReturnedRows
	if numRemainedRows < numMaxRows {
		numRows = numRemainedRows
	}

	if numRows <= 0 {
		return nil
	}

	result := r.rows[r.numReturnedRows : r.numReturnedRows+numRows]
	r.numReturnedRows += numRows
	return result
}
