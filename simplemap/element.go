package jsonmap

import "github.com/metalim/jsonmap"

// element is an element of a map, to be used in iteration.
//
//	for elem := m.First(); elem != nil; elem = elem.Next() {
//	    fmt.Println(elem.Key(), elem.Value())
//	}
type element struct {
	m     *Map
	index int
}

// Key returns the key of the element.
//
//	key := elem.Key()
func (e *element) Key() Key {
	return e.m.keys[e.index]
}

// Value returns the value of the element.
//
//	value := elem.Value()
func (e *element) Value() Value {
	return e.m.values[e.Key()]
}

// Next returns the next element in the map, for iteration.
// Returns nil if this is the last element.
// O(1) time.
//
//	for elem := m.First(); elem != nil; elem = elem.Next() {
//	    fmt.Println(elem.Key(), elem.Value())
//	}
func (e *element) Next() jsonmap.Element {
	if e.index == len(e.m.keys)-1 {
		return nil
	}
	return &element{
		m:     e.m,
		index: e.index + 1,
	}
}

// Prev returns the previous element in the map, for backwards iteration.
// Returns nil if this is the first element.
// O(1) time.
//
//	for elem := m.Last(); elem != nil; elem = elem.Prev() {
//	    fmt.Println(elem.Key(), elem.Value())
//	}
func (e *element) Prev() jsonmap.Element {
	if e.index == 0 {
		return nil
	}
	return &element{
		m:     e.m,
		index: e.index - 1,
	}
}

// First returns the first element in the map, for iteration.
// Returns nil if the map is empty.
// O(1) time.
//
//	for elem := m.First(); elem != nil; elem = elem.Next() {
//	    fmt.Println(elem.Key(), elem.Value())
//	}
func (m *Map) First() jsonmap.Element {
	return &element{
		m:     m,
		index: 0,
	}
}

// Last returns the last element in the map, for iteration for backwards iteration.
// Returns nil if the map is empty.
// O(1) time.
//
//	for elem := m.Last(); elem != nil; elem = elem.Prev() {
//	    fmt.Println(elem.Key(), elem.Value())
//	}
func (m *Map) Last() jsonmap.Element {
	return &element{
		m:     m,
		index: len(m.keys) - 1,
	}
}

// GetElement returns the element for the key, for iteration from a needle.
// Returns nil if the key is not in the map.
// O(n) for existing keys, because it uses KeyIndex().
func (m *Map) GetElement(key Key) jsonmap.Element {
	if _, ok := m.values[key]; !ok {
		return nil
	}
	return &element{
		m:     m,
		index: m.KeyIndex(key),
	}
}
