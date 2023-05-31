package jsonmap_test

import (
	"testing"

	"github.com/metalim/jsonmap"
)

func TestJSONMap(t *testing.T) {
	m := jsonmap.New()
	m.Set("a", 1)
	m.Set("d", 2)
	m.Set("c", 3)
	m.Set("e", 5)
	m.Set("b", 6)
	m.Set("b", 7) // rewrite
	m.Delete("c")
	m.Delete("c") // no error
	m.Delete("e") // test for index bug: deletes "e" and not "b"
	assert(t, m.Len() == 3)

	// keys
	assert(t, len(m.Keys()) == 3)
	assert(t, m.Keys()[0] == "a")
	assert(t, m.Keys()[1] == "d")
	assert(t, m.Keys()[2] == "b")

	// values
	v, ok := m.Get("a")
	assert(t, ok)
	assert(t, v.(int) == 1)
	v, ok = m.Get("d")
	assert(t, ok)
	assert(t, v.(int) == 2)
	v, ok = m.Get("b")
	assert(t, ok)
	assert(t, v.(int) == 7)
	v, ok = m.Get("c")
	assert(t, !ok)
	assert(t, v == nil)
	v, ok = m.Get("e")
	assert(t, !ok)
	assert(t, v == nil)
}
