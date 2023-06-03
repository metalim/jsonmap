package jsonmap

import "sort"

// O(n)
func (m *Map) KeyIndex(key string) int {
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

// O(n)
func (m *Map) Keys() []Key {
	keys := make([]Key, 0, len(m.elements))
	for elem := m.first; elem != nil; elem = elem.next {
		keys = append(keys, elem.key)
	}
	return keys
}

// O(n)
func (m *Map) Values() []Value {
	values := make([]Value, 0, len(m.elements))
	for elem := m.first; elem != nil; elem = elem.next {
		values = append(values, elem.value)
	}
	return values
}

// O(n)
func (m *Map) Elements() []*Element {
	elements := make([]*Element, 0, len(m.elements))
	for elem := m.first; elem != nil; elem = elem.next {
		elements = append(elements, elem)
	}
	return elements
}

// O(n log n)
func (m *Map) SortKeys(less func(a, b Key) bool) {
	if m.first == nil {
		return
	}
	if m.first == m.last {
		return
	}
	elements := make([]*Element, 0, len(m.elements))
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

// optional, as it can't be stopped
func (m *Map) ForEach(f func(key Key, value Value)) {
	for elem := m.first; elem != nil; elem = elem.next {
		f(elem.key, elem.value)
	}
}
