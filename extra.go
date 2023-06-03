package jsonmap

func (m *Map) SetFront(key Key, value Value) {
	if elem, ok := m.elements[key]; ok {
		elem.value = value
		return
	}
	elem := &Element{
		key:   key,
		value: value,
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
