package jsonmap_test

import (
	"fmt"
	"testing"

	"github.com/metalim/jsonmap"
	jsonmap1 "github.com/metalim/jsonmap/jsonmap1"
	jsonmap2 "github.com/metalim/jsonmap/jsonmap2"
	jsonmap3 "github.com/metalim/jsonmap/jsonmap3"
	jsonmap4 "github.com/metalim/jsonmap/jsonmap4"
	jsonmap5 "github.com/metalim/jsonmap/jsonmap5"
)

const (
	// 10000 keys to set
	SET_KEYS = 100000
	// 1000 keys to delete
	DELETE_KEYS = 10000
	// 10000 keys to get
	GET_KEYS = 100000
)

type Map interface {
	Set(key string, value any)
	Get(key string) (value any, ok bool)
	Delete(key string)
}

func benchmarkMap[T Map](b *testing.B, new func() T) {
	for i := 0; i < b.N; i++ {
		m := new()
		for j := 0; j < SET_KEYS; j++ {
			m.Set(fmt.Sprintf("key%d", j), j)
		}
		for j := 0; j < DELETE_KEYS; j++ {
			m.Delete(fmt.Sprintf("key5%d", j))
		}
		for j := 0; j < 10000; j++ {
			m.Get(fmt.Sprintf("key%d", j))
		}
	}
}

func BenchmarkMap(b *testing.B) {
	benchmarkMap(b, jsonmap.New)
}
func BenchmarkMap1(b *testing.B) {
	benchmarkMap(b, jsonmap1.New)
}
func BenchmarkMap2(b *testing.B) {
	benchmarkMap(b, jsonmap2.New)
}
func BenchmarkMap3(b *testing.B) {
	benchmarkMap(b, jsonmap3.New)
}
func BenchmarkMap4(b *testing.B) {
	benchmarkMap(b, jsonmap4.New)
}
func BenchmarkMap5(b *testing.B) {
	benchmarkMap(b, jsonmap5.New)
}

func BenchmarkMap5Keys(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := jsonmap5.New()
		for j := 0; j < SET_KEYS; j++ {
			m.Set(fmt.Sprintf("key%d", j), j)
		}
		for j := 0; j < DELETE_KEYS; j++ {
			m.Delete(fmt.Sprintf("key5%d", j))
		}
		for _, k := range m.Keys() {
			m.Get(k)
		}
	}
}

func BenchmarkMap5Iter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := jsonmap5.New()
		for j := 0; j < SET_KEYS; j++ {
			m.Set(fmt.Sprintf("key%d", j), j)
		}
		for j := 0; j < DELETE_KEYS; j++ {
			m.Delete(fmt.Sprintf("key5%d", j))
		}
		for el := m.First(); el != nil; el = el.Next() {
			_ = el.Key
			_ = el.Value
		}
	}
}
