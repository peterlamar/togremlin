package gremlin

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTranslateXMLWtihKey(t *testing.T) {

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

	rtnData := TranslateXMLWithKey(rawInput, gremlinKeys)

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
		t.Errorf("TranslateXMLWithKey return value was incorrect")
		fmt.Println(cmp.Diff(expectedReturn, rtnData))
	}
}

func TestTranslateXML(t *testing.T) {
	rawInput := []byte(`<?xml version="1.0" encoding="UTF-8"?>
	 - <note>
					<timestamp>2018-08-25T18:42:58+00:00</timestamp>
				 <to>Humans</to>
				 <from>Dolphins</from>
				 <heading>So Long</heading>
				 <body>Thanks for All the Fish!</body>
		 </note>`)

	rtnData := TranslateXML(rawInput)

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
		t.Errorf("translateXML return value was incorrect")
		fmt.Println(cmp.Diff(expectedReturn, rtnData))
	}
}

func TestTranslateXMLTwoMessages(t *testing.T) {
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

	rtnData := TranslateXML(rawInput)

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
		t.Errorf("translateXML return value was incorrect")
		fmt.Println(cmp.Diff(expectedReturn, rtnData))
	}
}

func TestTranslateXMLWithParent(t *testing.T) {
	rawInput := []byte(`<?xml version="1.0" encoding="UTF-8"?>
<guide>
   <name>HitchHiker</name>
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
</guide>
			 `)

	rtnData := TranslateXML(rawInput)

	expectedReturn := map[string][]map[string]interface{}{
		"guide": []map[string]interface{}{
			0: map[string]interface{}{
				"name": "HitchHiker",
			},
		},
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
		t.Errorf("translateXML return value was incorrect")
		fmt.Println(cmp.Diff(expectedReturn, rtnData))
	}
}

func TestTranslateXMLWithParentAndKey(t *testing.T) {
	rawInput := []byte(`<?xml version="1.0" encoding="UTF-8"?>
<guide>
   <name>HitchHiker</name>
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
</guide>
			 `)

	gremlinKeys := []byte(`{
	"guide": {
		"name": "_key",
		"notes": {
			"note": {
				"timestamp": "_key"
			}
		}
	}
}`)

	rtnData := TranslateXMLWithKey(rawInput, gremlinKeys)

	expectedReturn := map[string][]map[string]interface{}{
		"guide": []map[string]interface{}{
			0: map[string]interface{}{
				"name": "HitchHiker",
				"_key": "HitchHiker",
			},
		},
		"_hasnote": []map[string]interface{}{
			0: map[string]interface{}{
				"_from": "guide/HitchHiker",
				"_to":   "note/2018-08-25T18:42:58+00:00",
			},
			1: map[string]interface{}{
				"_from": "guide/HitchHiker",
				"_to":   "note/2018-09-25T02:30:28+00:00",
			},
		},
		"note": []map[string]interface{}{
			0: map[string]interface{}{
				"timestamp": "2018-08-25T18:42:58+00:00",
				"to":        "Humans",
				"from":      "Dolphins",
				"heading":   "So Long",
				"body":      "Thanks for All the Fish!",
				"_key":      "2018-08-25T18:42:58+00:00",
			},
			{
				"timestamp": "2018-09-25T02:30:28+00:00",
				"to":        "Humans",
				"from":      "Douglas Adams",
				"heading":   "Space",
				"body":      "Space is big. You just won't believe how vastly, hugely, mind- bogglingly big it is. I mean, you may think it is a long way down the road to the chemists, but thats just peanuts to space.",
				"_key":      "2018-09-25T02:30:28+00:00",
			},
		},
	}

	eq := cmp.Equal(rtnData, expectedReturn)

	if !eq {
		t.Errorf("translateXML return value was incorrect")
		fmt.Println(cmp.Diff(expectedReturn, rtnData))
	}
}

func TestTranslateJSON(t *testing.T) {
	rawInput := []byte(`{
  "notes": {
    "note": [
      {
        "timestamp": "2018-08-25T18:42:58+00:00",
        "to": "Humans",
        "from": "Dolphins",
        "heading": "So Long",
        "body": "Thanks for All the Fish!"
      },
      {
        "timestamp": "2018-09-25T02:30:28+00:00",
        "to": "Humans",
        "from": "Douglas Adams",
        "heading": "Space",
        "body": "Space is big. You just won't believe how vastly, hugely, mind- bogglingly big it is. I mean, you may think it is a long way down the road to the chemists, but thats just peanuts to space."
      }
    ]
  }
}`)

	rtnData := TranslateJSON(rawInput)

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
		t.Errorf("TestTranslateJSON return value was incorrect")
		fmt.Println(cmp.Diff(expectedReturn, rtnData))
	}
}
