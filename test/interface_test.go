package test_test

import "github.com/metalim/jsonmap"

type Key = string
type Value = any
type IMap interface {
	Clear()
	Len() int
	Get(key Key) (value Value, ok bool)
	Set(key Key, value Value)
	Delete(key Key)
	Push(key Key, value Value)
	Pop() (key Key, value Value, ok bool)

	SetFront(key Key, value Value)
	PushFront(key Key, value Value)
	PopFront() (key Key, value Value, ok bool)

	KeyIndex(key Key) int
	MarshalJSON() ([]byte, error)
	UnmarshalJSON(data []byte) error
	SortKeys(less func(a, b Key) bool)
	String() string
	Keys() []Key
	Values() []Value
}
type IElementPointers interface {
	First() *jsonmap.Element
	Last() *jsonmap.Element
	GetElement(key Key) *jsonmap.Element
}

var _ IMap = (*jsonmap.Map)(nil)
var _ IElementPointers = (*jsonmap.Map)(nil)
