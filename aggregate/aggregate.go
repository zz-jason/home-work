package aggregate

import (
	"github.com/zz-jason/home-work/aggregate/memreader"
)

type Aggregate interface {
	SetReader(memreader.MemReader)
	GetResult() [][]float64
}
