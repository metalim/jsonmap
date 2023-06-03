package jsonmap

// for compatibility with main version of jsonmap

type Element struct {
	m     *Map
	index int
}

func (m *Map) First() *Element {
	return &Element{
		m:     m,
		index: 0,
	}
}

func (m *Map) Last() *Element {
	return &Element{
		m:     m,
		index: len(m.keys) - 1,
	}
}

func (e *Element) Key() Key {
	return e.m.keys[e.index]
}

func (e *Element) Value() Value {
	return e.m.values[e.Key()]
}

func (e *Element) Next() *Element {
	if e.index == len(e.m.keys)-1 {
		return nil
	}
	return &Element{
		m:     e.m,
		index: e.index + 1,
	}
}

func (e *Element) Prev() *Element {
	if e.index == 0 {
		return nil
	}
	return &Element{
		m:     e.m,
		index: e.index - 1,
	}
}
