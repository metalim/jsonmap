package jsonmap

// O(n) delete
func (m *Map) Delete(key string) {
	if _, ok := m.values[key]; !ok {
		return
	}
	delete(m.values, key)
	i := m.KeyIndex(key)
	m.keys = append(m.keys[:i], m.keys[i+1:]...)
}

// Same as Set, but pushes the key to the end
// O(n) if key exists, as it uses Delete()
func (m *Map) Push(key Key, value Value) {
	m.Delete(key)
	m.keys = append(m.keys, key)
	m.values[key] = value
}

// O(n) if key is new
func (m *Map) SetFront(key Key, value Value) {
	if _, ok := m.values[key]; !ok {
		m.keys = append(m.keys, "")
		copy(m.keys[1:], m.keys)
		m.keys[0] = key
	}
	m.values[key] = value
}

// O(n)
func (m *Map) PushFront(key Key, value Value) {
	m.Delete(key)
	m.keys = append(m.keys, "")
	copy(m.keys[1:], m.keys)
	m.keys[0] = key
	m.values[key] = value
}

// O(n)
func (m *Map) KeyIndex(key string) int {
	for i, k := range m.keys {
		if k == key {
			return i
		}
	}
	return -1
}

// O(n)
func (m *Map) Values() []Value {
	values := make([]Value, len(m.keys))
	for i, k := range m.keys {
		values[i] = m.values[k]
	}
	return values
}

// O(n) for existing keys, because it uses KeyIndex()
func (m *Map) GetElement(key Key) *Element {
	if _, ok := m.values[key]; !ok {
		return nil
	}
	return &Element{
		m:     m,
		index: m.KeyIndex(key),
	}
}
