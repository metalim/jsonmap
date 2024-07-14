package test_test

import (
	"testing"

	"github.com/metalim/jsonmap"
	"github.com/zeebo/assert"
)

// due to random nature of Go map iteration, we need to test it many times
const TEST_ITERATIONS = 100

func TestJSONMap(t *testing.T) {
	for i := 0; i < TEST_ITERATIONS; i++ {
		testJSONMapOnce(t, i)
	}
}

func testJSONMapOnce(t *testing.T, i int) {
	m := jsonmap.New()

	// no elements
	assert.Equal(t, m.Len(), 0)
	assert.Equal(t, len(m.Keys()), 0)
	assert.Equal(t, m.String(), "map[]")
	assert.Nil(t, m.First())
	assert.Nil(t, m.Last())

	// 1 element
	m.Set("y", 1)
	assert.Equal(t, m.Len(), 1)
	assert.Equal(t, len(m.Keys()), 1)
	assert.Equal(t, m.Keys()[0], "y")
	assert.Equal(t, m.First().Key(), m.Keys()[0])
	assert.Equal(t, m.First().Key(), m.Last().Key())
	assert.Nil(t, m.First().Next())
	assert.Nil(t, m.Last().Prev())
	assert.Equal(t, m.String(), "map[y:1]")

	// delete single element
	m.Delete("y")
	assert.Equal(t, m.Len(), 0)
	assert.Equal(t, len(m.Keys()), 0)
	assert.Equal(t, m.String(), "map[]")
	assert.Nil(t, m.First())
	assert.Nil(t, m.Last())

	// 1 element
	m.Set("z", 1)
	assert.Equal(t, m.Len(), 1)
	assert.Equal(t, len(m.Keys()), 1)
	assert.Equal(t, m.Keys()[0], "z")
	assert.Equal(t, m.First().Key(), m.Keys()[0])
	assert.Equal(t, m.First().Key(), m.Last().Key())
	assert.Nil(t, m.First().Next())
	assert.Nil(t, m.Last().Prev())
	assert.Equal(t, m.String(), "map[z:1]")

	// clear map
	m.Clear()
	assert.Equal(t, m.Len(), 0)
	assert.Equal(t, len(m.Keys()), 0)
	assert.Equal(t, m.String(), "map[]")
	assert.Nil(t, m.First())
	assert.Nil(t, m.Last())

	// 2 elements
	m.Set("a", 1)
	m.Set("d", "2")
	assert.Equal(t, m.Len(), 2)
	assert.Equal(t, len(m.Keys()), 2)
	assert.Equal(t, m.Keys()[0], "a")
	assert.Equal(t, m.Keys()[1], "d")
	assert.Equal(t, m.First().Key(), m.Keys()[0])
	assert.Equal(t, m.Last().Key(), m.Keys()[1])
	assert.Nil(t, m.First().Prev())
	assert.Equal(t, m.First().Next(), m.Last())
	assert.Equal(t, m.Last().Prev(), m.First())
	assert.Nil(t, m.Last().Next())
	assert.Equal(t, m.String(), "map[a:1 d:2]")

	m.Set("c", 3)
	m.Set("e", 5)
	m.Set("b", 6)
	m.Set("b", 7) // rewrite
	m.Delete("c")
	m.Delete("c") // no error
	m.Delete("e") // test for index bug: delete "e" and not "b"
	assert.Equal(t, m.Len(), 3)

	// keys in insertion order
	assert.Equal(t, len(m.Keys()), 3)
	assert.Equal(t, m.Keys()[0], "a")
	assert.Equal(t, m.Keys()[1], "d")
	assert.Equal(t, m.Keys()[2], "b")

	// values
	v, ok := m.Get("a")
	assert.True(t, ok)
	assert.Equal(t, v.(int), 1)
	vInt, ok := jsonmap.GetAs[int](m, "a")
	assert.True(t, ok)
	assert.Equal(t, vInt, 1)
	vString, ok := jsonmap.GetAs[string](m, "a")
	assert.False(t, ok)
	assert.Equal(t, vString, "")

	v, ok = m.Get("d")
	assert.True(t, ok)
	assert.Equal(t, v.(string), "2")
	vString, ok = jsonmap.GetAs[string](m, "d")
	assert.True(t, ok)
	assert.Equal(t, vString, "2")
	vInt, ok = jsonmap.GetAs[int](m, "d")
	assert.False(t, ok)
	assert.Equal(t, vInt, 0)

	vInt, ok = jsonmap.GetAs[int](m, "nonexistent")
	assert.False(t, ok)
	assert.Equal(t, vInt, 0)

	v, ok = m.Get("b")
	assert.True(t, ok)
	assert.Equal(t, v.(int), 7)
	v, ok = m.Get("c")
	assert.False(t, ok)
	assert.Nil(t, v)
	v, ok = m.Get("e")
	assert.False(t, ok)
	assert.Nil(t, v)
}
