package jsonmap

// Map is a map with saved order of elements.
// It is useful for marshaling/unmarshaling JSON objects.
//
// Same as native map, but it keeps order of insertion when iterating or
// serializing to JSON, and has additional methods to iterate from any element.
// Similar to native map, user has to take care of concurrent access and nil map value.
type Map struct {
	elements    map[Key]*Element
	first, last *Element
	escapeHTML  bool
}

// New returns a new map. O(1) time.
//
//	m := jsonmap.New()
func New() *Map {
	return &Map{
		elements:   make(map[Key]*Element),
		escapeHTML: true,
	}
}

// Clear removes all elements from the map. O(1) time.
//
//	m.Clear()
func (m *Map) Clear() {
	m.elements = make(map[Key]*Element)
	m.first = nil
	m.last = nil
}

// Len returns the number of elements in the map, similar to len(m) for native map. O(1) time.
//
//	l := m.Len()
func (m *Map) Len() int {
	return len(m.elements)
}

// Get returns the value for the key, similar to m[key] for native map.
// Returns ok=false if the key is not in the map.
// O(1) time.
//
//	value, ok := m.Get(key)
func (m *Map) Get(key Key) (value Value, ok bool) {
	elem, ok := m.elements[key]
	if !ok {
		return // ok=false
	}
	return elem.value, true
}

// Helper function to get the value as a specific type.
// Returns ok=false if the key is not in the map or the value is not of the requested type.
//
//	str, ok := jsonmap.GetAs[string](m, key)
func GetAs[T any](m *Map, key Key) (value T, ok bool) {
	elem, ok := m.elements[key]
	if !ok {
		return
	}
	value, ok = elem.value.(T)
	return
}

// Set sets the value for the key.
// If key is already in the map, it replaces the value, but keeps the original order of the element.
// O(1) time.
//
//	m.Set(key, value)
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

// Delete removes the element from the map.
// O(1) time.
//
//	m.Delete(key)
func (m *Map) Delete(key Key) {
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

// Push is same as Set, but moves the element to the end of the map, as if it was just added.
// O(1) time.
//
//	m.Push(key, value)
func (m *Map) Push(key Key, value Value) {
	m.Delete(key)
	m.Set(key, value)
}

// Pop removes the last element from the map and returns it.
// Returns ok=false if the map is empty.
// O(1) time.
//
//	key, value, ok := m.Pop()
func (m *Map) Pop() (key Key, value Value, ok bool) {
	if m.last == nil {
		return // ok=false
	}
	key = m.last.key
	value = m.last.value
	m.Delete(key)
	return key, value, true
}

// First returns the first element in the map, for iteration.
// Returns nil if the map is empty.
// O(1) time.
func (m *Map) First() *Element {
	return m.first
}

// Last returns the last element in the map, for iteration for backwards iteration.
// Returns nil if the map is empty.
// O(1) time.
func (m *Map) Last() *Element {
	return m.last
}

// GetElement returns the element for the key, for iteration from a needle.
// Returns nil if the key is not in the map.
// O(1) time.
func (m *Map) GetElement(key Key) *Element {
	if el, ok := m.elements[key]; ok {
		return el
	}
	return nil
}

// SetEscapeHTML sets whether HTML characters should be escaped when marshaling to JSON.
// This setting is propagated to nested maps and arrays when unmarshaling from JSON.
// Set to true by default.
// O(1) time.
func (m *Map) SetEscapeHTML(escape bool) {
	m.escapeHTML = escape
}
