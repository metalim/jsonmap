package jsonmap

type Key = string
type Value = any

// Element is an element of a map, to be used in iteration.
//
//	for elem := m.First(); elem != nil; elem = elem.Next() {
//	    fmt.Println(elem.Key(), elem.Value())
//	}
type Element struct {
	key   Key
	value Value

	next, prev *Element
}

// Key returns the key of the element.
//
//	key := elem.Key()
func (e *Element) Key() Key {
	return e.key
}

// Value returns the value of the element.
//
//	value := elem.Value()
func (e *Element) Value() Value {
	return e.value
}

// Next returns the next element in the map, for iteration.
// Returns nil if this is the last element.
//
//	for elem := m.First(); elem != nil; elem = elem.Next() {
//	    fmt.Println(elem.Key(), elem.Value())
//	}
func (e *Element) Next() *Element {
	return e.next
}

// Prev returns the previous element in the map, for backwards iteration.
// Returns nil if this is the first element.
//
//	for elem := m.Last(); elem != nil; elem = elem.Prev() {
//	    fmt.Println(elem.Key(), elem.Value())
//	}
func (e *Element) Prev() *Element {
	return e.prev
}

// Map is a map with ordered elements.
// Same as native map, but keeps order of insertion when serializing to JSON or iterating,
// and with additional methods iterate from any point.
// Similar to native map, user has to take care of concurrency and nil map value.
type Map struct {
	elements    map[Key]*Element
	first, last *Element
}

// New returns a new map.
//
//	m := jsonmap.New()
func New() *Map {
	return &Map{
		elements: make(map[Key]*Element),
	}
}

// Clear removes all elements from the map.
//
//	m.Clear()
func (m *Map) Clear() {
	m.elements = make(map[Key]*Element)
	m.first = nil
	m.last = nil
}

// Len returns the number of elements in the map, similar to len(m) for native map.
//
//	l := m.Len()
func (m *Map) Len() int {
	return len(m.elements)
}

// Get returns the value for the key, similar to m[key] for native map.
// Returns ok=false if the key is not in the map.
//
//	value, ok := m.Get(key)
func (m *Map) Get(key string) (value any, ok bool) {
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
func GetAs[T any](m *Map, key string) (value T, ok bool) {
	elem, ok := m.elements[key]
	if !ok {
		return
	}
	value, ok = elem.value.(T)
	return
}

// Set sets the value for the key.
// If key is already in the map, it replaces the value, but keeps the original order of the element.
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
//
//	m.Delete(key)
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

// Push is same as Set, but moves the element to the end of the map, as if it was just added.
//
//	m.Push(key, value)
func (m *Map) Push(key Key, value Value) {
	m.Delete(key)
	m.Set(key, value)
}

// Pop removes the last element from the map and returns it.
// Returns ok=false if the map is empty.
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
func (m *Map) First() *Element {
	return m.first
}

// Last returns the last element in the map, for iteration for backwards iteration.
// Returns nil if the map is empty.
func (m *Map) Last() *Element {
	return m.last
}

// GetElement returns the element for the key, for iteration from a needle.
// Returns nil if the key is not in the map.
func (m *Map) GetElement(key string) *Element {
	return m.elements[key]
}
