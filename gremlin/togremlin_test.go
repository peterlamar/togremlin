package gremlin

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

/*
func TestTranslateWtihKey(t *testing.T) {

	rawInput := []byte(`<?xml version="1.0" encoding="UTF-8"?>
	        <note>
	             <timestamp>2018-08-25T18:42:58+00:00</timestamp>
	            <to>Humans</to>
	            <from>Dolphins</from>
	            <heading>So Long</heading>
	            <body>Thanks for All the Fish!</body>
	        </note>`)

	gremlinKeys := []byte(`{
	   	"note": {
	   		"timestamp": "_key"
	   	}
	   }`)

	rtnData := TranslateWithKey(rawInput, gremlinKeys)

	expectedReturn := map[string][]map[string]interface{}{
		"note": []map[string]interface{}{
			0: map[string]interface{}{
				"timestamp": "2018-08-25T18:42:58+00:00",
				"to":        "Humans",
				"from":      "Dolphins",
				"heading":   "So Long",
				"body":      "Thanks for All the Fish!",
				"_key":      "2018-08-25T18:42:58+00:00",
			},
		},
	}

	eq := cmp.Equal(rtnData, expectedReturn)

	if !eq {
		t.Errorf("TranslateWithKey return value was incorrect")
		fmt.Println(cmp.Diff(expectedReturn, rtnData))
	}
}

func TestTranslate(t *testing.T) {
	rawInput := []byte(`<?xml version="1.0" encoding="UTF-8"?>
	 - <note>
					<timestamp>2018-08-25T18:42:58+00:00</timestamp>
				 <to>Humans</to>
				 <from>Dolphins</from>
				 <heading>So Long</heading>
				 <body>Thanks for All the Fish!</body>
		 </note>`)

	rtnData := Translate(rawInput)

	expectedReturn := map[string][]map[string]interface{}{
		"note": []map[string]interface{}{
			0: map[string]interface{}{
				"timestamp": "2018-08-25T18:42:58+00:00",
				"to":        "Humans",
				"from":      "Dolphins",
				"heading":   "So Long",
				"body":      "Thanks for All the Fish!",
			},
		},
	}

	eq := cmp.Equal(rtnData, expectedReturn)

	if !eq {
		t.Errorf("translate return value was incorrect")
		fmt.Println(cmp.Diff(expectedReturn, rtnData))
	}
}
*/
func TestTranslateTwoMessages(t *testing.T) {
	rawInput := []byte(`<?xml version="1.0" encoding="UTF-8"?>
		<notes>
		   <note>
		      <timestamp>2018-08-25T18:42:58+00:00</timestamp>
		      <to>Humans</to>
		      <from>Dolphins</from>
		      <heading>So Long</heading>
		      <body>Thanks for All the Fish!</body>
		   </note>
		   <note>
		      <timestamp>2018-09-25T02:30:28+00:00</timestamp>
		      <to>Humans</to>
		      <from>Douglas Adams</from>
		      <heading>Space</heading>
		      <body>Space is big. You just won't believe how vastly, hugely, mind- bogglingly big it is. I mean, you may think it is a long way down the road to the chemists, but thats just peanuts to space.</body>
		   </note>
		</notes>
			 `)

	rtnData := Translate(rawInput)

	expectedReturn := map[string][]map[string]interface{}{
		"note": []map[string]interface{}{
			0: map[string]interface{}{
				"timestamp": "2018-08-25T18:42:58+00:00",
				"to":        "Humans",
				"from":      "Dolphins",
				"heading":   "So Long",
				"body":      "Thanks for All the Fish!",
			},
			{
				"timestamp": "2018-09-25T02:30:28+00:00",
				"to":        "Humans",
				"from":      "Douglas Adams",
				"heading":   "Space",
				"body":      "Space is big. You just won't believe how vastly, hugely, mind- bogglingly big it is. I mean, you may think it is a long way down the road to the chemists, but thats just peanuts to space.",
			},
		},
	}

	eq := cmp.Equal(rtnData, expectedReturn)

	if !eq {
		t.Errorf("translate return value was incorrect")
		fmt.Println(cmp.Diff(expectedReturn, rtnData))
	}
}
