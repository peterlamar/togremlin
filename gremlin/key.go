package gremlin

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

// const for gremlin flag, designates a node's unique id
const gremlinKeyHeader = "_key"

// Useful internal format for Gremlin graph data
type gremlinData struct {
	Keys  map[string]string
	Edges map[string][]string
}

// Filles the getKeyData struct with values from the key file
// passed in
func getKeyData(keyFile interface{}) gremlinData {
	var gmData gremlinData

	gmData.Keys = make(map[string]string)
	gmData.Edges = make(map[string][]string, 0)

	switch v := keyFile.(type) {
	case []uint8:
		s := reflect.ValueOf(keyFile)

		jsonMap := make(map[string]interface{})
		err := json.Unmarshal(s.Bytes(), &jsonMap)
		if err != nil {
			panic(err)
		}

		getKeyDataRecursive(jsonMap, "", gmData)

	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}

	return gmData
}

// Recursive loop to traverse the key file
func getKeyDataRecursive(kv map[string]interface{}, parent string, grmData gremlinData) {

	for k, v := range kv {
		buildParentChild(k, v, parent, grmData)

		if mv, ok := v.(map[string]interface{}); ok {
			getKeyDataRecursive(mv, k, grmData)
		} else {
			buildGremlinHeaders(k, v, parent, grmData)
		}
	}
}

// Build the Gremlin headers
func buildGremlinHeaders(k string, v interface{}, parent string, grmData gremlinData) {

	if v == gremlinKeyHeader {
		grmData.Keys[parent] = k
		return
	}

}

// Builds the parent child edges of the key file
func buildParentChild(k string, v interface{}, parent string, grmData gremlinData) {

	if parent == k || parent == "" {
		return
	}

	if _, ok := v.(map[string]interface{}); ok {
		grmData.Edges[parent] = append(grmData.Edges[parent], k)
	}

}

// Check if object has a key in our gremlin file
//  and we want to make it a node
func hasKey(grmData gremlinData, node string) bool {
	_, rtn := grmData.Keys[node]
	return rtn
}

// Get valid edges with a key that point to the node
func getValidEdgesToNode(grmData gremlinData, node string) []string {
	var rtnEdges []string

	// Get edges that point to me
	myEdges := getEdgesToNode(grmData, node)

	// Find valid edges with key (or check their parents )
	for _, edge := range myEdges {
		// If it has a key, great
		if hasKey(grmData, edge) {

			rtnEdges = append(rtnEdges, edge)
		} else {
			// Otherwise, see if edge parent has a key
			parentEdges := getEdgesToNode(grmData, edge)

			if len(parentEdges) == 1 && hasKey(grmData, parentEdges[0]) {
				rtnEdges = append(rtnEdges, parentEdges[0])
			} else if len(parentEdges) > 1 {
				log.Fatal("Node has too many parents", edge)
			}
		}
	}

	return rtnEdges
}

// Get edges that point to a node
func getEdgesToNode(grmData gremlinData, node string) []string {
	var myEdges []string
	for edgeName, edges := range grmData.Edges {
		for _, edgeDest := range edges {
			if node == edgeDest {
				myEdges = append(myEdges, edgeName)
			}
		}
	}
	return myEdges
}
