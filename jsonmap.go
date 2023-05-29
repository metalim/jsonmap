package jsonmap

type Key = string
type Value = any
type Order = int

type element struct {
	value Value
	order Order
}

type Map struct {
	m         map[Key]element
	keys      []Key
	nextOrder Order
}

func New() *Map {
	return &Map{
		m: make(map[Key]element),
	}
}

func (m *Map) Set(key Key, value Value) {
	if el, ok := m.m[key]; ok {
		m.m[key] = element{value: value, order: el.order}
		return
	}
	m.keys = append(m.keys, key)
	m.m[key] = element{value: value, order: m.nextOrder}
	m.nextOrder++
}

func (m *Map) Get(key string) (value any, ok bool) {
	el, ok := m.m[key]
	return el.value, ok
}

func (m *Map) Delete(key string) {
	if _, ok := m.m[key]; !ok {
		return
	}
	delete(m.m, key)
	i := m.index(key)
	m.keys = append(m.keys[:i], m.keys[i+1:]...)
}

func (m *Map) index(key string) int {
	for i, k := range m.keys {
		if k == key {
			return i
		}
	}
	return -1
}

func (m *Map) Len() int {
	return len(m.m)
}

func (m *Map) Keys() []string {
	return m.keys
}
