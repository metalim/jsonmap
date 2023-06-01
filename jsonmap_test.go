package jsonmap_test

import (
	"testing"

	"github.com/metalim/jsonmap"
)

func TestJSONMap(t *testing.T) {
	for i := 0; i < 100; i++ {
		testJSONMap(t, i)
	}
}

func testJSONMap(t *testing.T, i int) {
	m := jsonmap.New()
	m.Set("a", 1)
	m.Set("d", 2)
	m.Set("c", 3)
	m.Set("e", 5)
	m.Set("b", 6)
	m.Set("b", 7) // rewrite
	m.Delete("c")
	m.Delete("c") // no error
	m.Delete("e") // test for index bug: delete "e" and not "b"
	assertEqual(t, m.Len(), 3)

	// keys
	assertEqual(t, len(m.Keys()), 3)
	assertEqual(t, m.Keys()[0], "a")
	assertEqual(t, m.Keys()[1], "d")
	assertEqual(t, m.Keys()[2], "b")

	// values
	v, ok := m.Get("a")
	assert(t, ok)
	assertEqual(t, v.(int), 1)
	v, ok = m.Get("d")
	assert(t, ok)
	assertEqual(t, v.(int), 2)
	v, ok = m.Get("b")
	assert(t, ok)
	assertEqual(t, v.(int), 7)
	v, ok = m.Get("c")
	assert(t, !ok)
	assertEqual(t, v, nil)
	v, ok = m.Get("e")
	assert(t, !ok)
	assertEqual(t, v, nil)
}
