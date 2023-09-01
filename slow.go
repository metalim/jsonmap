package jsonmap

import "sort"

// KeyIndex returns index of key. O(n) time.
// If key is not in the map, it returns -1.
// O(n) time if the key exists.
// O(1) time if the key does not exist.
//
//	i := m.KeyIndex(key)
func (m *Map) KeyIndex(key Key) int {
	elem, ok := m.elements[key]
	if !ok {
		return -1
	}
	i := 0
	for e := m.first; e != elem; e = e.next {
		i++
	}
	return i
}

// Keys returns all keys in the map. O(n) time and space.
//
//	keys := m.Keys()
func (m *Map) Keys() []Key {
	keys := make([]Key, 0, len(m.elements))
	for elem := m.first; elem != nil; elem = elem.next {
		keys = append(keys, elem.key)
	}
	return keys
}

// Values returns all values in the map. O(n) time and space.
//
//	values := m.Values()
func (m *Map) Values() []Value {
	values := make([]Value, 0, len(m.elements))
	for elem := m.first; elem != nil; elem = elem.next {
		values = append(values, elem.value)
	}
	return values
}

// SortKeys sorts keys in the map. O(n*log(n)) time, O(n) space.
// O(n*log(n)) time, O(n) space.
//
//	m.SortKeys(func(a, b Key) bool {
//		return a < b
//	})
func (m *Map) SortKeys(less func(a, b Key) bool) {
	if m.Len() < 2 {
		return
	}
	elements := make([]*element, 0, len(m.elements))
	for elem := m.first; elem != nil; elem = elem.next {
		elements = append(elements, elem)
	}

	sort.Slice(elements, func(i, j int) bool {
		return less(elements[i].key, elements[j].key)
	})
	m.first = elements[0]
	m.last = elements[len(elements)-1]
	for i := 0; i < len(elements)-1; i++ {
		elements[i].next = elements[i+1]
		elements[i+1].prev = elements[i]
	}
}
