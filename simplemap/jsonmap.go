package jsonmap

type Key = string
type Value = any

type Map struct {
	values map[Key]Value
	keys   []Key
}

func New() *Map {
	return &Map{
		values: make(map[Key]Value),
	}
}

func (m *Map) Clear() {
	m.values = make(map[Key]Value)
	m.keys = m.keys[:0]
}

func (m *Map) Len() int {
	return len(m.values)
}

func (m *Map) Get(key string) (value any, ok bool) {
	value, ok = m.values[key]
	return
}

func (m *Map) Set(key Key, value Value) {
	if _, ok := m.values[key]; !ok {
		m.keys = append(m.keys, key)
	}
	m.values[key] = value
}

func (m *Map) Keys() []string {
	return m.keys
}

func (m *Map) Merge(other *Map) {
	for _, k := range other.keys {
		m.Set(k, other.values[k])
	}
}
