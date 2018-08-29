package gremlin

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTranslate(t *testing.T) {

	rawInput := []byte(`<?xml version="1.0" encoding="UTF-8"?>
   - <note>
          <timestamp>2018-08-25T18:42:58+00:00</timestamp>
         <to>Humans</to>
         <from>Dolphins</from>
         <heading>So Long</heading>
         <body>Thanks for all the fish!</body>
     </note>`)

	gremlinKeys := []byte(`{
	"note": {
		"timestamp": "_key"
	}
}`)

	rtnData := Translate(rawInput, gremlinKeys)

	expectedReturn := map[string][]map[string]interface{}{
		"note": []map[string]interface{}{
			0: map[string]interface{}{
				"timestamp": "2018-08-25T18:42:58+00:00",
				"to":        "Humans",
				"from":      "Dolphins",
				"heading":   "So Long",
				"_key":      "2018-08-25T18:42:58+00:00",
			},
		},
	}

	eq := cmp.Equal(rtnData, expectedReturn)

	if !eq {
		t.Errorf("translate return value was incorrect")
	}

}
