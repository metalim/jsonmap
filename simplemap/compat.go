package jsonmap

import "sort"

// for compatibility with main version of jsonmap

type Element struct {
	m     *Map
	index int
}

func (m *Map) First() *Element {
	return &Element{
		m:     m,
		index: 0,
	}
}

func (m *Map) Last() *Element {
	return &Element{
		m:     m,
		index: len(m.keys) - 1,
	}
}

func (e *Element) Key() Key {
	return e.m.keys[e.index]
}

func (e *Element) Value() Value {
	return e.m.values[e.Key()]
}

func (e *Element) Next() *Element {
	if e.index == len(e.m.keys)-1 {
		return nil
	}
	return &Element{
		m:     e.m,
		index: e.index + 1,
	}
}

func (e *Element) Prev() *Element {
	if e.index == 0 {
		return nil
	}
	return &Element{
		m:     e.m,
		index: e.index - 1,
	}
}

// O(1)
func (m *Map) Pop() (key Key, value Value, ok bool) {
	if len(m.keys) == 0 {
		return // ok=false
	}
	key = m.keys[len(m.keys)-1]
	value = m.values[key]
	delete(m.values, key)
	m.keys = m.keys[:len(m.keys)-1]
	return key, value, true
}

// SortKeys sorts keys in the map. O(n*log(n)) time.
//
//	m.SortKeys(func(a, b Key) bool {
//		return a < b
//	})
func (m *Map) SortKeys(less func(a, b Key) bool) {
	sort.Slice(m.keys, func(i, j int) bool {
		return less(m.keys[i], m.keys[j])
	})
}
