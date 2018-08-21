package memreader

import (
	"testing"
)

func TestBuild(t *testing.T) {
	reader := Build(TypeColumn)
	_, isMemReaderCol := reader.(*memReaderCol)
	if !isMemReaderCol {
		t.Errorf("The returned MemReader should be memReaderCol.")
	}

	reader = Build(TypeRow)
	_, isMemReaderRow := reader.(*memReaderRow)
	if !isMemReaderRow {
		t.Errorf("The returned MemReader should be memReaderRow.")
	}
}
