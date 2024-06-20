package jsonmap

type Key = string
type Value = any

// Element of a map, to be used in iteration.
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
// O(1) time.
//
//	for elem := m.First(); elem != nil; elem = elem.Next() {
//	    fmt.Println(elem.Key(), elem.Value())
//	}
func (e *Element) Next() *Element {
	return e.next
}

// Prev returns the previous element in the map, for backwards iteration.
// Returns nil if this is the first element.
// O(1) time.
//
//	for elem := m.Last(); elem != nil; elem = elem.Prev() {
//	    fmt.Println(elem.Key(), elem.Value())
//	}
func (e *Element) Prev() *Element {
	return e.prev
}
