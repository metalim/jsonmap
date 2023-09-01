package jsonmap_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/metalim/jsonmap"
	simplemap "github.com/metalim/jsonmap/simplemap"
)

const (
	TEST_ITERATIONS = 100
	plainJSON       = `{"a":1,"c":3,"d":4,"b":2,"e":5}`
	nestedJSON      = `{"1":1,"c":null,"d":"dd","b":true,"e":[2,null,{"x":1,"z":3,"y":2},"zzz"],"f":{"a":1,"c":3,"d":4,"b":2,"e":5}}`
)

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

func testMap[T IMapShort](t *testing.T, new func() T) {
	t.Run("PlainJSON", func(t *testing.T) {
		// due to random nature of Go map iteration, we need to test it many times
		for i := 0; i < TEST_ITERATIONS; i++ {
			m := new()
			testSerialization(t, plainJSON, m)
			verifyPlainJSON(t, m)
		}
	})

	t.Run("NestedJSON", func(t *testing.T) {
		for i := 0; i < TEST_ITERATIONS; i++ {
			m := new()
			testSerialization(t, nestedJSON, m)
		}
	})
}

func testSerialization(t *testing.T, testData string, m IMapShort) {
	err := json.Unmarshal([]byte(testData), &m)
	assertEqual(t, err, nil)

	data, err := json.Marshal(m)
	assertEqual(t, err, nil)
	assertEqual(t, string(data), testData)
}

func verifyPlainJSON(t *testing.T, m IMapShort) {
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
}
