package memreader

import (
	"math/rand"
	"time"
)

type memReaderCol struct {
	columns   [][]int64
	populater populater

	numReturnedRows int
}

func (r *memReaderCol) ResetPopulater(other populater) {
	r.populater = other
}

func (r *memReaderCol) Populate() {
	if r.populater != nil {
		r.columns = r.populater()
		return
	}

	numRows, numCols := 100000000, 2
	r.columns = make([][]int64, numCols)
	r.columns[0] = make([]int64, numRows)
	r.columns[1] = make([]int64, numRows)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < numRows; i++ {
		r.columns[0][i] = rand.Int63n(1000)
	}
	for i := 0; i < numRows; i++ {
		r.columns[1][i] = rand.Int63n(2000) - 1000
	}
}

func (r *memReaderCol) Next(numMaxRows int) [][]int64 {
	numRows := numMaxRows
	numRemainedRows := len(r.columns[0]) - r.numReturnedRows
	if numRemainedRows < numMaxRows {
		numRows = numRemainedRows
	}

	if numRows <= 0 {
		return nil
	}

	result := make([][]int64, 2)
	result[0] = r.columns[0][r.numReturnedRows : r.numReturnedRows+numRows]
	result[1] = r.columns[1][r.numReturnedRows : r.numReturnedRows+numRows]
	r.numReturnedRows += numRows
	return result
}
