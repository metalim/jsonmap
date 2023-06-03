package jsonmap

import "sort"

type Key = string
type Value = any
type Rank = int

type Map struct {
	values map[Key]Value
	keys   []Key

	// for Delete to be O(log N) we use ranks
	nextRank   Rank
	keyRanks   map[Key]Rank
	indexRanks []Rank
}

func New() *Map {
	return &Map{
		values:   make(map[Key]Value),
		keyRanks: make(map[Key]Rank),
	}
}

func (m *Map) Set(key Key, value Value) {
	if _, ok := m.values[key]; !ok {
		m.keys = append(m.keys, key)
		m.keyRanks[key] = m.nextRank
		m.indexRanks = append(m.indexRanks, m.nextRank)
		m.nextRank++
	}
	m.values[key] = value
}

// Same as Set, but pushes the key to the end
func (m *Map) Push(key Key, value Value) {
	m.Delete(key)
	m.keys = append(m.keys, key)
	m.values[key] = value
	m.keyRanks[key] = m.nextRank
	m.indexRanks = append(m.indexRanks, m.nextRank)
	m.nextRank++
}

func (m *Map) Get(key string) (value any, ok bool) {
	value, ok = m.values[key]
	return
}

// even when IndexKey is O(log N), Delete is still O(N) because of the slice
func (m *Map) Delete(key string) {
	if _, ok := m.values[key]; !ok {
		return
	}
	i := m.IndexKey(key)
	m.keys = append(m.keys[:i], m.keys[i+1:]...)
	m.indexRanks = append(m.indexRanks[:i], m.indexRanks[i+1:]...)
	delete(m.keyRanks, key)
	delete(m.values, key)
}

// O(log n) via binary search,
// or even O(1) via more advanced data structures. But lets keep it simple for now
func (m *Map) IndexKey(key string) int {
	rank, ok := m.keyRanks[key]
	if !ok {
		return -1
	}

	// find the index of the rank
	return sort.SearchInts(m.indexRanks, rank)
}

func (m *Map) Len() int {
	return len(m.values)
}

func (m *Map) Keys() []string {
	return m.keys
}
