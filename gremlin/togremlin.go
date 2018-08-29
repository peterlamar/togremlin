package gremlin

import (
	"fmt"
	"log"
	"reflect"

	"github.com/clbanning/mxj"
)

// Translate data structure into hash map of json array objects
func Translate(input interface{},
	keys interface{}) map[string][]map[string]interface{} {

	var rtn map[string][]map[string]interface{}

	switch v := input.(type) {
	case []uint8:
		s := reflect.ValueOf(input)

		// Unmarshal to [map]interface{}
		mvj, err := mxj.NewMapXml(s.Bytes())

		if err != nil {
			log.Fatal(err)
		}

		translateNodes(mvj, keys, "")

	default:
		fmt.Printf("Translate doesn't handle type %T!\n", v)
	}

	return rtn
}

// Translate nodes into the graph output
func translateNodes(m map[string]interface{}, keyFile interface{},
	parent string) map[string][]map[string]interface{} {

	var rtn map[string][]map[string]interface{}

	return rtn
}
