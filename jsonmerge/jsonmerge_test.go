package jsonmerge

import (
	// "fmt"
	"reflect"
	"testing"
)

func TestMergeOneLevel(t *testing.T) {
	x1 := map[string]int{"X": 1}
	y1 := map[string]int{"Y": 2}
	var m1 Merger
	m1.Merge(x1)
	m1.Merge(y1)
	var res1 map[string]int = m1
	eq := reflect.DeepEqual(x1, res1)
	if !eq {
		t.Errorf("first level failed")
	}
}
