package jsonmap

import (
	"fmt"
	"strings"
)

func (m *Map) String() string {
	var b strings.Builder
	b.WriteString(`map[`)
	for el := m.First(); el != nil; el = el.Next() {
		if el != m.First() { // safe to compare pointers for main version
			b.WriteByte(' ')
		}
		b.WriteString(el.Key())
		b.WriteByte(':')
		b.WriteString(fmt.Sprint(el.Value()))
	}
	b.WriteByte(']')
	return b.String()
}
