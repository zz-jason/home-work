package memreader

import (
	"fmt"
)

type TypeCode int

const (
	TypeColumn TypeCode = iota
	TypeRow
)

// Build creates a MemReader implementation according to the required TypeCode.
// Only two type are supported now, e.g, TypeColumn and TypeRow.
func Build(code TypeCode) MemReader {
	switch code {
	case TypeColumn:
		return new(memReaderCol)
	case TypeRow:
		return new(memReaderRow)
	}
	panic(fmt.Sprintf("Build: unsupported TypeCode(%v)", code))
}
