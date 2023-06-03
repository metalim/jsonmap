package jsonmap_test

import (
	"fmt"
	"testing"

	"github.com/metalim/jsonmap"
	simplemap "github.com/metalim/jsonmap/simplemap"
)

const (
	// 10000 keys to set
	SET_KEYS = 100000
	// 1000 keys to delete
	DELETE_KEYS = 10000
	// 10000 keys to get
	GET_KEYS = 100000
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
		for j := 0; j < 10000; j++ {
			m.Get(fmt.Sprintf("key%d", j))
		}
	}
}

func BenchmarkMap(b *testing.B) {
	benchmarkMap(b, jsonmap.New)
}
func BenchmarkSimpleMap(b *testing.B) {
	benchmarkMap(b, simplemap.New)
}
