package jsonmerge

import (
	// "fmt"
	"reflect"
	"testing"
)

func TestMergeOneLevel(t *testing.T) {
	m1 := Merger{"X": 1}
	t1 := Merger{"Y": 2}
	m2 := Merger{"X": 1, "Y": 2}
	m1.Merge(t1)
	eq := reflect.DeepEqual(m1, m2)
	if !eq {
		t.Errorf("first level failed")
	}
}
