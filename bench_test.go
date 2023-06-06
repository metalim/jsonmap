package jsonmap_test

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"

	"github.com/metalim/jsonmap"
	simplemap "github.com/metalim/jsonmap/simplemap"
)

const (
	SET_KEYS    = 1e5
	DELETE_KEYS = 1e4
	GET_KEYS    = 1e5

	PREPARE_KEYS  = 1e7
	KEY_LEN       = 10
	LOG_KEY_STATS = true
)

func benchmarkSuite[T Map](b *testing.B, new func() T) {
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

func setupMap[T Map](b *testing.B, new func() T) (T, string) {
	const symbols = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_-"

	// single string for fast keys access
	var sb strings.Builder
	for j := 0; j < PREPARE_KEYS+KEY_LEN; j++ {
		sb.WriteByte(symbols[rand.Intn(len(symbols))])
	}
	keys := sb.String()

	m := new()
	rand.Seed(1)
	for j := 0; j < PREPARE_KEYS; j++ {
		index := rand.Intn(PREPARE_KEYS)
		m.Set(keys[index:index+KEY_LEN], j)
	}

	return m, keys
}

func benchmarkMapGet(b *testing.B, m Map, keys string) {
	var new, existing int
	for i := 0; i < b.N; i++ {
		_, ok := m.Get(keys[i%PREPARE_KEYS : i%PREPARE_KEYS+KEY_LEN])
		if LOG_KEY_STATS {
			if ok {
				existing++
			} else {
				new++
			}
		}
	}
	if LOG_KEY_STATS {
		b.Logf("N: %d, new: %d, existing: %d", b.N, new, existing)
	}
}

func benchmarkMapSetExisting(b *testing.B, m Map, keys string) {
	var new, existing int
	for i := 0; i < b.N; i++ {
		key := keys[i%PREPARE_KEYS : i%PREPARE_KEYS+KEY_LEN]
		if LOG_KEY_STATS {
			_, ok := m.Get(key)
			if ok {
				existing++
			} else {
				new++
			}
		}
		m.Set(key, i)
	}
	if LOG_KEY_STATS {
		b.Logf("N: %d, new: %d, existing: %d", b.N, new, existing)
	}
}

func benchmarkMapSetNew(b *testing.B, m Map, keys string) {
	var new, existing int
	for i := 0; i < b.N; i++ {
		key := keys[i%PREPARE_KEYS+1 : i%PREPARE_KEYS+KEY_LEN]
		if LOG_KEY_STATS {
			_, ok := m.Get(key)
			if ok {
				existing++
			} else {
				new++
			}
		}
		m.Set(key, i)
	}
	if LOG_KEY_STATS {
		b.Logf("N: %d, new: %d, existing: %d", b.N, new, existing)
	}
}

func benchmarkMapDelete(b *testing.B, m Map, keys string) {
	var new, existing int
	for i := 0; i < b.N; i++ {
		key := keys[i%PREPARE_KEYS : i%PREPARE_KEYS+KEY_LEN]
		if LOG_KEY_STATS {
			_, ok := m.Get(key)
			if ok {
				existing++
			} else {
				new++
			}
		}
		m.Delete(key)
	}
	if LOG_KEY_STATS {
		b.Logf("N: %d, new: %d, existing: %d", b.N, new, existing)
	}
}

var mapDefs = []struct {
	name string
	new  func() Map
}{
	{"jsonmap", func() Map { return jsonmap.New() }},
	{"simplemap", func() Map { return simplemap.New() }},
}

var ops = []struct {
	name      string
	benchmark func(*testing.B, Map, string)
}{
	{"Get", benchmarkMapGet},
	{"SetExisting", benchmarkMapSetExisting},
	{"SetNew", benchmarkMapSetNew},
	{"Delete", benchmarkMapDelete},
}

func Benchmark(b *testing.B) {
	b.Run("Suite", func(b *testing.B) {
		for _, mapDef := range mapDefs {
			b.Run(mapDef.name, func(b *testing.B) {
				benchmarkSuite(b, mapDef.new)
			})
		}
	})

	b.Run("Ops", func(b *testing.B) {
		for _, op := range ops {
			b.Run(op.name, func(b *testing.B) {
				for _, def := range mapDefs {
					b.Run(def.name, func(b *testing.B) {
						m, keys := setupMap(b, def.new)
						b.ResetTimer()
						op.benchmark(b, m, keys)
					})
				}
			})
		}
	})
}
