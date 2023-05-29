package main

import (
	"fmt"

	"github.com/metalim/jsonmap"
)

func main() {
	demo1()
}

func demo1() {
	m := jsonmap.New()
	m.Set("a", "bar")
	m.Set("c", "qux")
	m.Set("aa", "quux")
	m.Set("ab", "corge")
	m.Set("ac", "grault")
	m.Set("ba", "garply")
	m.Set("bb", "waldo")
	m.Set("b", "fred")

	m.Delete("aa")
	m.Delete("bb")

	m2 := jsonmap.New()
	m2.Set("a", "bar")
	m2.Set("c", "qux")
	m.Set("ccc", m2)

	m.Set("aaa", "plugh")
	m.Set("aab", "xyzzy")
	m.Set("aac", "thud")
	m.Set("baa", "foo")

	fmt.Println(*m)
	var iterate func(m *jsonmap.Map, prefix string)
	iterate = func(m *jsonmap.Map, prefix string) {
		for _, key := range m.Keys() {
			val, _ := m.Get(key)
			if v, ok := val.(*jsonmap.Map); ok {
				iterate(v, prefix+key+"/")
				continue
			}
			fmt.Printf("%s%s: %v\n", prefix, key, val)
		}
	}
	iterate(m, "")
}
