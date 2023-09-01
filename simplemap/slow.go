package jsonmap

import "sort"

// Delete deletes the key from the map.
// O(n) time if the key exists, because it uses KeyIndex().
// O(1) time if the key does not exist.
//
//	m.Delete(key)
func (m *Map) Delete(key Key) {
	if _, ok := m.values[key]; !ok {
		return
	}
	delete(m.values, key)
	i := m.KeyIndex(key)
	m.keys = append(m.keys[:i], m.keys[i+1:]...)
}

// Same as Set, but pushes the key to the end
// O(n) if key exists, as it uses Delete()
// O(1) if key is new.
//
//	m.Push(key, value)
func (m *Map) Push(key Key, value Value) {
	m.Delete(key)
	m.keys = append(m.keys, key)
	m.values[key] = value
}

// Same as Set, but pushes the key to the front if it is new.
// O(n) if key is new.
// O(1) if key exists.
//
//	m.SetFront(key, value)
func (m *Map) SetFront(key Key, value Value) {
	if _, ok := m.values[key]; !ok {
		m.keys = append(m.keys, "")
		copy(m.keys[1:], m.keys)
		m.keys[0] = key
	}
	m.values[key] = value
}

// PushFront is same as SetFront, but moves the element to the front of the map, as if it was just added.
// O(n) time.
//
//	m.PushFront(key, value)
func (m *Map) PushFront(key Key, value Value) {
	m.Delete(key)
	m.keys = append(m.keys, "")
	copy(m.keys[1:], m.keys)
	m.keys[0] = key
	m.values[key] = value
}

// PopFront removes the first element from the map and returns its key and value.
// Returns ok=false if the map is empty.
// O(n) time.
//
//	key, value, ok := m.PopFront()
func (m *Map) PopFront() (key Key, value Value, ok bool) {
	if len(m.keys) == 0 {
		return // ok=false
	}
	key = m.keys[0]
	value = m.values[key]
	m.Delete(key)
	return key, value, true
}

// KeyIndex returns index of key. O(n) time.
// If key is not in the map, it returns -1.
// O(n) time.
//
//	i := m.KeyIndex(key)
func (m *Map) KeyIndex(key Key) int {
	for i, k := range m.keys {
		if k == key {
			return i
		}
	}
	return -1
}

// Values returns all values in the map. O(n) time and O(n) space.
// O(1) time and space.
//
//	values := m.Values()
func (m *Map) Values() []Value {
	values := make([]Value, len(m.keys))
	for i, k := range m.keys {
		values[i] = m.values[k]
	}
	return values
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
