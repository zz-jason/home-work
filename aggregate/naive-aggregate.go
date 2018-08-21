package aggregate

import (
	"sort"

	"github.com/zz-jason/home-work/aggregate/memreader"
)

type naiveAggregate struct {
	reader memreader.MemReader
}

func (e *naiveAggregate) SetReader(reader memreader.MemReader) {
	e.reader = reader
}

func (e *naiveAggregate) GetResult() [][]float64 {
	allInputs := e.read()
	allInputs = e.sort(allInputs)
	return e.aggregate(allInputs)
}

func (e *naiveAggregate) read() [][]int64 {
	allInputs := make([][]int64, 0, 1<<20)
	for {
		input := e.reader.Next(1)
		if input == nil {
			return allInputs
		}
		allInputs = append(allInputs, input...)
	}
}

func (e *naiveAggregate) sort(allInputs [][]int64) [][]int64 {
	sort.Slice(allInputs, func(i, j int) bool { return allInputs[i][0] < allInputs[j][0] })
	return allInputs
}

func (e *naiveAggregate) aggregate(allInputs [][]int64) [][]float64 {
	if len(allInputs) <= 0 {
		return nil
	}

	result := make([][]float64, 0, 1024)
	key, sum, count := allInputs[0][0], float64(allInputs[0][1]), 1
	for i := 1; i < len(allInputs); i++ {
		if allInputs[i][0] != key {
			result = append(result, []float64{float64(key), sum / float64(count)})
			key, sum, count = allInputs[i][0], float64(allInputs[i][1]), 1
			continue
		}
		sum += float64(allInputs[i][1])
		count += 1
	}

	result = append(result, []float64{float64(key), sum / float64(count)})
	return result
}
