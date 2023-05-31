package jsonmap_test

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/metalim/jsonmap"
)

func TestJSON(t *testing.T) {
	m := jsonmap.New()

	// unmarshal
	json.Unmarshal([]byte(testData), &m)
	log.Println(m)

	assert(t, m.Len() == 5)

	// keys
	assert(t, len(m.Keys()) == 5)
	assert(t, m.Keys()[0] == "a")
	assert(t, m.Keys()[1] == "c")
	assert(t, m.Keys()[2] == "d")
	assert(t, m.Keys()[3] == "b")
	assert(t, m.Keys()[4] == "e")

	// values
	v, ok := m.Get("a")
	assert(t, ok)
	assert(t, v.(float64) == 1)
	v, ok = m.Get("c")
	assert(t, ok)
	assert(t, v.(float64) == 3)
	v, ok = m.Get("d")
	assert(t, ok)
	assert(t, v.(float64) == 4)
	v, ok = m.Get("b")
	assert(t, ok)
	assert(t, v.(float64) == 2)
	v, ok = m.Get("e")
	assert(t, ok)
	assert(t, v.(float64) == 5)
	v, ok = m.Get("f")
	assert(t, !ok)
	assert(t, v == nil)

	// marshal
	// data, err := json.Marshal(m)
	// assert(t, err == nil)
	// assert(t, string(data) == testData)
}

const testData = `{"a":1,"c":3,"d":4,"b":2,"e":5}`
