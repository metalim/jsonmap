package jsonmap_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/metalim/jsonmap"
	simplemap "github.com/metalim/jsonmap/simplemap"
)

const (
	PREPARE_KEYS = 1e7
	SET_KEYS     = 1e5
	DELETE_KEYS  = 1e4
	GET_KEYS     = 1e5
)

func benchmarkMap[T Map](b *testing.B, new func() T) {
	for i := 0; i < b.N; i++ {
		m := new()
		for j := 0; j < SET_KEYS; j++ {
			m.Set(fmt.Sprintf("key%d", j), j)
		}
		for j := 0; j < DELETE_KEYS; j++ {
			m.Delete(fmt.Sprintf("key5%d", j))
		}
		for j := 0; j < GET_KEYS; j++ {
			m.Get(fmt.Sprintf("key%d", j))
		}
	}
}

func setupMap[T Map](b *testing.B, new func() T) (T, []string) {
	m := new()
	keys := make([]string, PREPARE_KEYS)
	for j := 0; j < PREPARE_KEYS; j++ {
		keys[j] = fmt.Sprintf("key%d", j)
		m.Set(keys[j], j)
	}
	rand.Shuffle(len(keys), func(i, j int) {
		keys[i], keys[j] = keys[j], keys[i]
	})
	return m, keys
}

func benchmarkMapGet(b *testing.B, m Map, keys []string) {
	for i := 0; i < b.N; i++ {
		m.Get(keys[i%len(keys)])
	}
}
func benchmarkMapSetExisting(b *testing.B, m Map, keys []string) {
	for i := 0; i < b.N; i++ {
		m.Set(keys[i%len(keys)], i)
	}
}
func benchmarkMapSetNew(b *testing.B, m Map, keys []string) {
	for i, k := range keys {
		keys[i] = k + "x"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Set(keys[i%len(keys)], i)
	}
}
func benchmarkMapDelete(b *testing.B, m Map, keys []string) {
	for i := 0; i < b.N; i++ {
		m.Delete(keys[i%len(keys)])
	}
}

var mapDefs = []struct {
	name string
	new  func() Map
}{
	{"jsonmap", func() Map { return jsonmap.New() }},
	{"simplemap", func() Map { return simplemap.New() }},
}

var benchDefs = []struct {
	name  string
	bench func(*testing.B, Map, []string)
}{
	{"Get", benchmarkMapGet},
	{"SetExisting", benchmarkMapSetExisting},
	{"SetNew", benchmarkMapSetNew},
	{"Delete", benchmarkMapDelete},
}

func BenchmarkMap(b *testing.B) {
	benchmarkMap(b, jsonmap.New)
}
func BenchmarkSimpleMap(b *testing.B) {
	benchmarkMap(b, simplemap.New)
}

func BenchmarkMapOps(b *testing.B) {
	for _, benchDef := range benchDefs {
		b.Run(benchDef.name, func(b *testing.B) {
			for _, def := range mapDefs {
				b.Run(def.name, func(b *testing.B) {
					m, keys := setupMap(b, def.new)
					b.ResetTimer()
					benchDef.bench(b, m, keys)
				})
			}
		})
	}
}
