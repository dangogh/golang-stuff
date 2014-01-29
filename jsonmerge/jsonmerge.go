package jsonmerge

import (
	//"encoding/json"
	"fmt"
	//"io"
)

type Merger map[string]interface{}

func (r Merger) Merge(a Merger) error {
	for k, v := range a {
		switch v.(type) {
		case []interface{}:
			fmt.Printf("%v - slice %T\n", k, v)
		case map[interface{}]interface{}:
			fmt.Printf("%v - map %T\n", k, v)
		default:
			r[k] = v
		}
	}
	return nil
}
