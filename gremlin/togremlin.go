package gremlin

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/clbanning/mxj"
)

// const for edge node, designates an edge origin relationship
const gremlinEdgeFrom = "_from"

// const for edge node, designates an edge dest relationship
const gremlinEdgeTo = "_to"

// prefix for graph edge relationship
const graphParentVerb = "_has"

func TranslateJSON(input interface{}) map[string][]map[string]interface{} {
	rtn := make(map[string][]map[string]interface{})

	switch v := input.(type) {
	case []uint8:
		s := reflect.ValueOf(input)

		var tempInterface interface{}

		err := json.Unmarshal(s.Bytes(), &tempInterface)
		if err != nil {
			panic(err)
		}

		reflectedJSON := reflect.ValueOf(tempInterface)

		var emptyGrmData gremlinData

		translateJSONNodesRecursive(reflectedJSON, "", emptyGrmData, rtn)

	default:
		fmt.Printf("TranslateJSON doesn't handle type %T!\n", v)
	}

	return rtn
}

// https://gist.github.com/hvoecking/10772475
// Travel nodes and build the gremlin graph data structure
func translateJSONNodesRecursive(vl reflect.Value, parent string,
	grmData gremlinData, rtnMap map[string][]map[string]interface{}) {

	switch vl.Kind() {

	case reflect.Slice:
		for i := 0; i < vl.Len(); i += 1 {
			translateJSONNodesRecursive(vl.Index(i), parent, grmData, rtnMap)
		}
	case reflect.Map:
		for _, key := range vl.MapKeys() {
			translateJSONNodesRecursive(vl.MapIndex(key), key.String(), grmData, rtnMap)
		}

	case reflect.Interface:

		if len(vl.Interface().(map[string]interface{})) == 1 {
			for k, v := range vl.Interface().(map[string]interface{}) {
				translateJSONNodesRecursive(reflect.ValueOf(v), k, grmData, rtnMap)
			}
		} else {
			var grmData gremlinData

			// if parent is not null and we don't have it in our map, make some room
			makeMemoryForNode(parent, grmData, rtnMap)
			t := make(map[string]interface{})

			for k, v := range vl.Interface().(map[string]interface{}) {
				collectValuesAndKeys(t, k, v, parent, grmData)
			}
			if _, ok := rtnMap[parent]; ok {

				// Add node to collection
				rtnMap[parent] = append(rtnMap[parent], t)
			}
		}

	default:
		fmt.Println("uncaught")
	}

}

// TranslateXML data structure as is
func TranslateXML(input interface{}) map[string][]map[string]interface{} {
	rtn := make(map[string][]map[string]interface{})

	switch v := input.(type) {
	case []uint8:
		s := reflect.ValueOf(input)

		// Unmarshal to [map]interface{}
		mvj, err := mxj.NewMapXml(s.Bytes())

		if err != nil {
			log.Fatal(err)
		}

		var emptyGrmData gremlinData

		translateXMLNodesRecursive(mvj, "", emptyGrmData, rtn)

		// Remove nodes with empty values
		for k, v := range rtn {
			for _, mv := range v {
				if len(mv) == 0 {
					delete(rtn, k)
				}
			}
		}

	default:
		fmt.Printf("TranslateXML doesn't handle type %T!\n", v)
	}

	return rtn
}

func TranslateJSONWithKey(input interface{},
	keys interface{}) map[string][]map[string]interface{} {
	var rtn map[string][]map[string]interface{}

	switch v := input.(type) {
	case []uint8:

		s := reflect.ValueOf(input)

		var tempInterface interface{}

		err := json.Unmarshal(s.Bytes(), &tempInterface)
		if err != nil {
			panic(err)
		}

		reflectedJSON := reflect.ValueOf(tempInterface)

		rtn = translateJSONNodes(reflectedJSON, keys, "")

	default:
		fmt.Printf("TranslateJSONWithKey doesn't handle type %T!\n", v)
	}

	return rtn
}

func translateJSONNodes(vl reflect.Value, keyFile interface{},
	parent string) map[string][]map[string]interface{} {

	grmData := getKeyData(keyFile)

	rtn := make(map[string][]map[string]interface{})

	translateJSONNodesRecursive(vl, "", grmData, rtn)

	return rtn

}

// TranslateXML data structure into hash map of json array objects
func TranslateXMLWithKey(input interface{},
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

		rtn = translateXMLNodes(mvj, keys, "")

	default:
		fmt.Printf("TranslateXMLWithKey doesn't handle type %T!\n", v)
	}

	return rtn
}

// TranslateXML nodes into the graph output
func translateXMLNodes(m map[string]interface{}, keyFile interface{},
	parent string) map[string][]map[string]interface{} {

	grmData := getKeyData(keyFile)

	rtn := make(map[string][]map[string]interface{})

	translateXMLNodesRecursive(m, "", grmData, rtn)

	return rtn
}

// Travel nodes and build the gremlin graph data structure
func translateXMLNodesRecursive(m map[string]interface{}, parent string,
	grmData gremlinData, rtnMap map[string][]map[string]interface{}) {

	// if parent is not null and we don't have it in our map, make some room
	makeMemoryForNode(parent, grmData, rtnMap)

	t := make(map[string]interface{})

	for k, v := range m {
		collectValuesAndKeys(t, k, v, parent, grmData)
	}

	if _, ok := rtnMap[parent]; ok {
		// Create edge nodes and load them into memory.
		createParentEdges(t, parent, grmData, rtnMap)

		// Add node to collection
		rtnMap[parent] = append(rtnMap[parent], t)
	}

	// Go through loop and now handle struct and arrays
	for k, v := range m {
		// If memeber is another json struct
		if mv, ok := v.(map[string]interface{}); ok {

			translateXMLNodesRecursive(mv, k, grmData, rtnMap)
			// If member is json array
		} else if mvs, ok := v.([]interface{}); ok {

			for _, mv := range mvs {
				if mv, ok := mv.(map[string]interface{}); ok {
					translateXMLNodesRecursive(mv, k, grmData, rtnMap)
				}
			}
		}
	}
}

// Collect values and keys
func collectValuesAndKeys(t map[string]interface{}, k string, v interface{},
	parent string, grmData gremlinData) {

	// If member is a string, add it
	if _, ok := v.(string); ok {
		t[k] = v
	}

	// If this value is a _key, duplicate it
	if value, ok := grmData.Keys[parent]; ok && value == k {
		t[gremlinKeyHeader] = v
	}

}

// Allocate memory for new Node
func makeMemoryForNode(parent string,
	grmData gremlinData, rtnMap map[string][]map[string]interface{}) {

	// If parent is empty, then this is the initial pass
	if parent == "" {
		return
	}

	// If the node already exists, no need to allocate memory
	if _, ok := rtnMap[parent]; ok {
		return
	}

	// If we have no keys then allocate the memory
	if len(grmData.Keys) == 0 {
		rtnMap[parent] = make([]map[string]interface{}, 0)
		return
	}

	if hasKey(grmData, parent) {
		rtnMap[parent] = make([]map[string]interface{}, 0)
	}

}

// Check if we need to make a parent/child edge and create it if necessary
func createParentEdges(currentNode map[string]interface{}, currentNodeName string,
	grmData gremlinData, rtnMap map[string][]map[string]interface{}) {

	validEdges := getValidEdgesToNode(grmData, currentNodeName)

	// Create edges
	for _, edge := range validEdges {
		if len(rtnMap[edge]) == 0 {
			log.Fatal("Key doesn't exist for edge ", edge)
		}
		element := rtnMap[edge][len(rtnMap[edge])-1]

		t := make(map[string]interface{})
		t[gremlinEdgeFrom] = edge + "/" + element[gremlinKeyHeader].(string)
		t[gremlinEdgeTo] = currentNodeName + "/" + currentNode[gremlinKeyHeader].(string)

		// Check if we need to grab values edge values from the collection
		edgename := graphParentVerb + currentNodeName

		rtnMap[edgename] = append(rtnMap[edgename], t)

	}
}
