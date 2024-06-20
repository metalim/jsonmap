package jsonmap

// SetFront sets the value for the key.
// If key is already in the map, it replaces the value, but keeps the original order of the element.
// If key is not in the map, it adds the element to the front of the map.
// O(1) time.
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

// PushFront is same as SetFront, but moves the element to the front of the map, as if it was just added.
// O(1) time.
//
//	m.PushFront(key, value)
func (m *Map) PushFront(key Key, value Value) {
	m.Delete(key)
	m.SetFront(key, value)
}

// PopFront removes the first element from the map and returns its key and value.
// Returns ok=false if the map is empty.
// O(1) time.
//
//	key, value, ok := m.PopFront()
func (m *Map) PopFront() (key Key, value Value, ok bool) {
	if m.first == nil {
		return // ok=false
	}
	key = m.first.key
	value = m.first.value
	ok = true
	m.Delete(key)
	return
}
