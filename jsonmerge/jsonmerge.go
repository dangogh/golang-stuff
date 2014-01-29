package jsonmerge

import (
	//"encoding/json"
	"fmt"
	//"io"
)

type Merger map[string]interface{}

func (r Merger) Merge(a interface{}) error {
	for k, v := range r {
		switch v.(type) {
		case []interface{}:
			fmt.Printf("%v - slice %T\n", k, v)
		case map[interface{}]interface{}:
			fmt.Printf("%v - map %T\n", k, v)
		default:
			fmt.Printf("%v - Merging to a %T\n", k, v)
		}
	}
	return nil
}
