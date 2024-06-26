package jsonmap

import (
	"fmt"
	"strings"
)

// String returns a string representation of the map. O(n) time.
func (m *Map) String() string {
	var b strings.Builder
	b.WriteString(`map[`)
	for el := m.First(); el != nil; el = el.Next() {
		if el != m.First() { // safe to compare pointers
			b.WriteByte(' ')
		}
		b.WriteString(el.Key())
		b.WriteByte(':')
		b.WriteString(fmt.Sprint(el.Value()))
	}
	b.WriteByte(']')
	return b.String()
}
