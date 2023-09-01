package jsonmap

type Key = string
type Value = any

// Map is a map[string]any with saved order of elements.
// It is useful for marshaling/unmarshaling JSON objects.
// This is a simplified version of jsonmap.Map, but with O(n) Delete.
type Map struct {
	values map[Key]Value
	keys   []Key
}

// New returns a new Map. O(1) time.
//
//	m := jsonmap.New()
func New() *Map {
	return &Map{
		values: make(map[Key]Value),
	}
}

// Clear removes all elements from the map. O(1) time.
//
//	m.Clear()
func (m *Map) Clear() {
	m.values = make(map[Key]Value)
	m.keys = m.keys[:0]
}

// Len returns the number of elements in the map. O(1) time.
//
//	n := m.Len()
func (m *Map) Len() int {
	return len(m.values)
}

// Get returns the value for the key.
// Returns ok=false if the key is not in the map.
// O(1) time.
//
//	value, ok := m.Get(key)
func (m *Map) Get(key Key) (value Value, ok bool) {
	value, ok = m.values[key]
	return
}

// Set sets the value for the key.
// If key is already in the map, it replaces the value, but keeps the original order of the element.
// O(1) time.
//
//	m.Set(key, value)
func (m *Map) Set(key Key, value Value) {
	if _, ok := m.values[key]; !ok {
		m.keys = append(m.keys, key)
	}
	m.values[key] = value
}

// Keys returns all keys in the map. O(n) time and O(n) space.
// O(1) time, as keys slice is stored in the map.
//
//	keys := m.Keys()
func (m *Map) Keys() []Key {
	return m.keys
}

// Helper function to get the value as a specific type.
// Returns ok=false if the key is not in the map or the value is not of the requested type.
//
//	str, ok := jsonmap.GetAs[string](m, key)
func GetAs[T any](m *Map, key Key) (value T, ok bool) {
	elem, ok := m.values[key]
	if !ok {
		return
	}
	value, ok = elem.(T)
	return
}

// Pop removes the last element from the map and returns it.
// Returns ok=false if the map is empty.
// O(1) time.
//
//	key, value, ok := m.Pop()
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
