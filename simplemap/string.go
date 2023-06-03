package jsonmap

import (
	"fmt"
	"strings"
)

func (m *Map) String() string {
	var b strings.Builder
	b.WriteString(`map[`)
	for i, k := range m.keys {
		if i > 0 {
			b.WriteByte(' ')
		}
		b.WriteString(k)
		b.WriteByte(':')
		b.WriteString(fmt.Sprint(m.values[k]))
	}
	b.WriteByte(']')
	return b.String()
}
