package memreader

type populater = func() [][]int64

type MemReader interface {
	ResetPopulater(other populater)
	Populate()
	Next(numMaxRows int) [][]int64
}

var (
	_ MemReader = (*memReaderCol)(nil)
	_ MemReader = (*memReaderRow)(nil)
)
