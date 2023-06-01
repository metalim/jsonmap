package jsonmap_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/metalim/jsonmap"
)

func TestSimpleJSON(t *testing.T) {
	// due to random nature of Go map iteration, we need to test it many times
	for i := 0; i < 100; i++ {
		testJSON(t, i, `{"a":1,"c":3,"d":4,"b":2,"e":5}`, func(m *jsonmap.Map) {
			assertEqual(t, m.Len(), 5)

			// keys
			assertEqual(t, len(m.Keys()), 5)
			assertEqual(t, strings.Join(m.Keys(), ","), "a,c,d,b,e")

			// values
			v, ok := m.Get("a")
			assert(t, ok)
			assertEqual(t, v.(float64), 1.)
			v, ok = m.Get("c")
			assert(t, ok)
			assertEqual(t, v.(float64), 3.)
			v, ok = m.Get("d")
			assert(t, ok)
			assertEqual(t, v.(float64), 4.)
			v, ok = m.Get("b")
			assert(t, ok)
			assertEqual(t, v.(float64), 2.)
			v, ok = m.Get("e")
			assert(t, ok)
			assertEqual(t, v.(float64), 5.)
			v, ok = m.Get("f")
			assert(t, !ok)
			assertEqual(t, v, nil)

		})
	}
}

func TestNestedJSON(t *testing.T) {
	for i := 0; i < 100; i++ {
		testJSON(t, i, `{"1":1,"c":null,"d":"dd","b":true,"e":[2,null,3.5,"zzz"],"f":{"a":1,"c":3,"d":4,"b":2,"e":5}}`, nil)
	}
}

func testJSON(t *testing.T, i int, testData string, verify func(m *jsonmap.Map)) {
	m := jsonmap.New()

	// unmarshal
	err := json.Unmarshal([]byte(testData), &m)
	assertEqual(t, err, nil)

	// verify
	if verify != nil {
		verify(m)
	}

	// marshal
	data, err := json.Marshal(m)
	assertEqual(t, err, nil)
	assertEqual(t, string(data), testData)
}
