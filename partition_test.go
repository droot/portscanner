package scanner

import (
  "testing"

  "github.com/stretchr/testify/assert"
)

type testStruct struct {
  inputSlice []uint16
  numParts   int
  expParts   [][]uint16
}

func TestPartition(t *testing.T) {
  tests := []testStruct{
    {
      []uint16{1, 2, 3, 4},
      2,
      [][]uint16{
        []uint16{1, 2},
        []uint16{3, 4},
      },
    },
    {
      []uint16{1, 2, 3, 4},
      3,
      [][]uint16{
        []uint16{1},
        []uint16{2},
        []uint16{3, 4},
      },
    },
    {
      []uint16{1, 2, 3, 4},
      1,
      [][]uint16{
        []uint16{1, 2, 3, 4},
      },
    },
  }

  for _, tt := range tests {
    parts := partition(tt.inputSlice, tt.numParts)
    assert.Equal(t, len(parts), tt.numParts)
    assert.Equal(t, parts, tt.expParts)
  }
}
