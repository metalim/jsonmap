package jsonmap

type Key = string
type Value = any

type Element struct {
	key   Key
	value Value

	next, prev *Element
}

func (e *Element) Key() Key {
	return e.key
}

func (e *Element) Value() Value {
	return e.value
}

func (e *Element) Next() *Element {
	return e.next
}

func (e *Element) Prev() *Element {
	return e.prev
}

type Map struct {
	elements    map[Key]*Element
	first, last *Element
}

func New() *Map {
	return &Map{
		elements: make(map[Key]*Element),
	}
}

func (m *Map) Clear() {
	m.elements = make(map[Key]*Element)
	m.first = nil
	m.last = nil
}

func (m *Map) Len() int {
	return len(m.elements)
}

func (m *Map) First() *Element {
	return m.first
}

func (m *Map) Last() *Element {
	return m.last
}

func (m *Map) GetElement(key string) *Element {
	return m.elements[key]
}

func (m *Map) Get(key string) (value any, ok bool) {
	elem, ok := m.elements[key]
	if !ok {
		return
	}
	return elem.value, true
}

func GetAs[T any](m *Map, key string) (value T, ok bool) {
	elem, ok := m.elements[key]
	if !ok {
		return
	}
	value, ok = elem.value.(T)
	return
}

func (m *Map) Set(key Key, value Value) {
	if elem, ok := m.elements[key]; ok {
		elem.value = value
		return
	}
	elem := &Element{
		key:   key,
		value: value,
	}
	m.elements[key] = elem
	if m.last == nil {
		m.first = elem
		m.last = elem
		return
	}
	m.last.next = elem
	elem.prev = m.last
	m.last = elem
}

func (m *Map) Delete(key string) {
	elem, ok := m.elements[key]
	if !ok {
		return
	}
	if elem.prev == nil {
		m.first = elem.next
	} else {
		elem.prev.next = elem.next
	}
	if elem.next == nil {
		m.last = elem.prev
	} else {
		elem.next.prev = elem.prev
	}
	delete(m.elements, key)
}

func (m *Map) Push(key Key, value Value) {
	m.Delete(key)
	m.Set(key, value)
}

func (m *Map) Merge(other *Map) {
	for elem := other.First(); elem != nil; elem = elem.Next() {
		m.Set(elem.key, elem.value)
	}
}
