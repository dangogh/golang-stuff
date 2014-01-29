package jsonmerge

import (
	// "fmt"
	"testing"
)

func TestMergeOneLevel(t *testing.T) {
	t1 := map[string]int{"X": 1, "Y": 2, "Z": 3}
	t2 := map[string]string{"P": "A", "D": "B", "Q": "C"}
	t.Errorf("got there.. %T %T\n", t1, t2)
}
