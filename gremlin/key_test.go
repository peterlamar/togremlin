package gremlin

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetKeyData(t *testing.T) {

	keyFile := []byte(`{
	"note": {
		"timestamp": "_key"
	}
}`)

	var grmData gremlinData

	grmData.Keys = make(map[string]string)
	grmData.Edges = make(map[string][]string, 0)

	grmData.Keys["note"] = "timestamp"

	rtnData := getKeyData(keyFile)

	eq := cmp.Equal(grmData, rtnData)
	if !eq {
		fmt.Println("getKeyData return value generated in a different order " +
			"(this is ok and can randomly occur)")
		fmt.Println(cmp.Diff(grmData, rtnData))
	}

}
