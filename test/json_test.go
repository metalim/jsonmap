package jsonmap_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/metalim/jsonmap"
	simplemap "github.com/metalim/jsonmap/simplemap"
	"github.com/zeebo/assert"
)

type test struct {
	name string
	json string
}

var tests = []test{
	{"PlainJSON", `{"a":1,"c":3,"d":4,"b":2,"e":5}`},
	{"NestedJSON", `{"1":1,"c":null,"d":"dd","b":true,"e":[2,null,{"x":1,"z":3,"y":2},"zzz"],"f":{"a":1,"c":3,"d":4,"b":2,"e":5}}`},
	{"EscapedJSON", `{"\"":"\"","\\":"\\"}`},
}

type IMapShort interface {
	Len() int
	Set(key Key, value Value)
	Get(key Key) (value Value, ok bool)
	Delete(key Key)
	Keys() []Key
}

func TestFastMap(t *testing.T) {
	testMap(t, jsonmap.New)
}

func TestSimpleMap(t *testing.T) {
	testMap(t, simplemap.New)
}

func testMap[T IMapShort](t *testing.T, newMap func() T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// due to random nature of Go map iteration, we need to test it many times
			for i := 0; i < TEST_ITERATIONS; i++ {
				m := newMap()
				testSerialization(t, test.json, m)
				if test.name == "PlainJSON" {
					verifyPlainJSON(t, m)
				}
			}
		})
	}
}

func testSerialization(t *testing.T, testData string, m IMapShort) {
	err := json.Unmarshal([]byte(testData), &m)
	assert.Equal(t, err, nil)

	data, err := json.Marshal(m)
	assert.Equal(t, err, nil)
	assert.Equal(t, string(data), testData)
}

func verifyPlainJSON(t *testing.T, m IMapShort) {
	assert.Equal(t, m.Len(), 5)

	// keys
	assert.Equal(t, len(m.Keys()), 5)
	assert.Equal(t, strings.Join(m.Keys(), ","), "a,c,d,b,e")

	// values
	v, ok := m.Get("a")
	assert.True(t, ok)
	assert.Equal(t, v.(float64), 1.)
	v, ok = m.Get("c")
	assert.True(t, ok)
	assert.Equal(t, v.(float64), 3.)
	v, ok = m.Get("d")
	assert.True(t, ok)
	assert.Equal(t, v.(float64), 4.)
	v, ok = m.Get("b")
	assert.True(t, ok)
	assert.Equal(t, v.(float64), 2.)
	v, ok = m.Get("e")
	assert.True(t, ok)
	assert.Equal(t, v.(float64), 5.)
	v, ok = m.Get("f")
	assert.False(t, ok)
	assert.Equal(t, v, nil)
}
