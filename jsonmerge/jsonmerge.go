package jsonmerge

import (
	//"encoding/json"
	"errors"
	"fmt"
	//"io"
)

type Merger interface {
	Merge(interface{}) error
}

type Accum struct {
	val interface{}
}

/*
	A	B	RES
	*	null	A
	null	*	B
	slice	slice	A+B
	slice	map	A
	map	slice	A
	map	map	A+B

*/

func (accum *Accum) mergeSlice(tomerge []interface{}) error {
	if accum.val == nil {
		accum.val = tomerge
		return nil
	}
	if a, ok := accum.val.([]interface{}); ok {
		accum.val = append(a, tomerge)
		return nil
	}
	return errors.New(fmt.Sprintf("Can't merge %T with %T", accum.val, tomerge))
}

func (accum *Accum) mergeMap(tomerge map[interface{}]interface{}) error {
	if a, ok := accum.(map[interface{}]interface{}); !ok {
		return errors.NewError
	for k, v := range accum.val {
		switch v.(type) {
		case []interface{}:
			accum.val[k].mergeSlice(v)
		case map[interface{}]interface{}:
			accum.val[k].mergeMap(v)
		default:
			accum.val[k] = v
		}
	}
	return nil
}

func (accum *Accum) Merge(tomerge interface{}) error {
	
}
