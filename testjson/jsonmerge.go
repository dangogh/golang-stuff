
package jsonmerge

import (
	//"encoding/json"
	"fmt"
	//"io"
)

type Merger interface{}

func (r Merger) Merge(a interface{}) (interface{}, error) {
	switch t := r.(type) {
	case []interface{}:
		fmt.Printf("slice %T\n", r)
	case map[interface{}]interface{}:
		fmt.Printf("map %T\n", r)
	default:
		fmt.Printf("Merging to a %T\n", r)
	}
}

