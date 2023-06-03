package jsonmap

import "sort"

func (m *Map) SetFront(key Key, value Value) {
	if elem, ok := m.elements[key]; ok {
		elem.Value = value
		return
	}
	elem := &Element{
		Key:   key,
		Value: value,
	}
	m.elements[key] = elem
	if m.first == nil {
		m.first = elem
		m.last = elem
		return
	}
	m.first.prev = elem
	elem.next = m.first
	m.first = elem
}

func (m *Map) PushFront(key Key, value Value) {
	m.Delete(key)
	m.SetFront(key, value)
}

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
		return less(elements[i].Key, elements[j].Key)
	})
	m.first = elements[0]
	m.last = elements[len(elements)-1]
	for i := 0; i < len(elements)-1; i++ {
		elements[i].next = elements[i+1]
		elements[i+1].prev = elements[i]
	}
}

// not for hot path, as it's O(N)
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

// not for hot path, as it's O(N)
func (m *Map) Keys() []Key {
	keys := make([]Key, 0, len(m.elements))
	for elem := m.first; elem != nil; elem = elem.next {
		keys = append(keys, elem.Key)
	}
	return keys
}

// not for hot path, as it's O(N)
func (m *Map) Values() []Value {
	values := make([]Value, 0, len(m.elements))
	for elem := m.first; elem != nil; elem = elem.next {
		values = append(values, elem.Value)
	}
	return values
}

// not for hot path, as it's O(N)
func (m *Map) Elements() []*Element {
	elements := make([]*Element, 0, len(m.elements))
	for elem := m.first; elem != nil; elem = elem.next {
		elements = append(elements, elem)
	}
	return elements
}

// optional, as it can't be stopped
func (m *Map) ForEach(f func(key Key, value Value)) {
	for elem := m.first; elem != nil; elem = elem.next {
		f(elem.Key, elem.Value)
	}
}
