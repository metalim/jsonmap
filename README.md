# Ordered map

[![go test workflow](https://github.com/metalim/jsonmap/actions/workflows/gotest.yml/badge.svg)](https://github.com/metalim/jsonmap/actions/workflows/gotest.yml)
[![go report](https://goreportcard.com/badge/github.com/metalim/jsonmap)](https://goreportcard.com/report/github.com/metalim/jsonmap)
[![codecov](https://codecov.io/gh/metalim/jsonmap/graph/badge.svg?token=HLGJ7U07JH)](https://codecov.io/gh/metalim/jsonmap)
[![go doc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/metalim/jsonmap)
[![Ask DeepWiki](https://deepwiki.com/badge.svg)](https://deepwiki.com/metalim/jsonmap)
[![MIT license](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/mit/)

Simple ordered map for Go, with JSON restrictions. The main purpose is to keep same order of keys after parsing JSON and generating it again, so Unmarshal followed by Marshal generates exactly the same JSON structure

Keys are strings, Values are any JSON values (number, string, boolean, null, array, map/object)

Storage is O(n), operations are O(1), except for optional operations in [slow.go](slow.go) file

When Unmarshalling, **any nested map from JSON is created as ordered jsonmap**, including maps in nested arrays

Inspired by [wk8/go-ordered-map](https://github.com/wk8/go-ordered-map) and [iancoleman/orderedmap](https://github.com/iancoleman/orderedmap)

## Performance

Similar to Go native map, `jsonmap` has O(1) time for Get, Set, Delete. Additionally it has Push, First, Last, Next, Prev operations, which are also O(1). Suite benchmark does 10k of Set and Get, and 1k of Delete operations. `jsonmap` performance is on par with native Go map

```
➜ go test -bench . -benchmem ./test
goos: darwin
goarch: arm64
pkg: github.com/metalim/jsonmap/test
Benchmark/Suite/jsonmap-10            36          31614692 ns/op        17046302 B/op     623006 allocs/op
Benchmark/Suite/gomap-10              48          25116164 ns/op        14791564 B/op     522925 allocs/op
Benchmark/Ops/Get/jsonmap-10            11191284               112.3 ns/op             0 B/op          0 allocs/op
Benchmark/Ops/Get/gomap-10              11543529                97.88 ns/op            0 B/op          0 allocs/op
Benchmark/Ops/SetExisting/jsonmap-10     5847512               291.3 ns/op           106 B/op          1 allocs/op
Benchmark/Ops/SetExisting/gomap-10       7608140               174.6 ns/op            89 B/op          1 allocs/op
Benchmark/Ops/SetNew/jsonmap-10          2548009               429.3 ns/op           242 B/op          2 allocs/op
Benchmark/Ops/SetNew/gomap-10            5143821               212.3 ns/op           133 B/op          1 allocs/op
Benchmark/Ops/Delete/jsonmap-10          7642382               156.0 ns/op             0 B/op          0 allocs/op
Benchmark/Ops/Delete/gomap-10           10592187               102.4 ns/op             0 B/op          0 allocs/op
PASS
ok      github.com/metalim/jsonmap/test 143.577s
```

## Installation
```bash
$ go get github.com/metalim/jsonmap
```

## Usage

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/metalim/jsonmap"
)

const sampleJSON = `{"an":"article","empty":null,"sub":{"s":1,"e":2,"x":3,"y":4},"bool":false,"array":[1,2,3]}`

func main() {
	m := jsonmap.New()

	// unmarshal, keeping order
	err := json.Unmarshal([]byte(sampleJSON), &m)
	if err != nil {
		panic(err)
	}

	// get values
	val, ok := m.Get("an")
	fmt.Println("an: ", val, ok) // article true
	val, ok = m.Get("non-existant")
	fmt.Println("non-existant", val, ok) // <nil> false

	// marshal, keeping order
	output, err := json.Marshal(&m)
	if err != nil {
		panic(err)
	}

	if string(output) == sampleJSON {
		fmt.Println("output == sampleJSON")
	}

	// iterate
	fmt.Println("forward order:")
	for el := m.First(); el != nil; el = el.Next() {
		fmt.Printf("\t%s: %v\n", el.Key(), el.Value())
	}
	fmt.Println()

	fmt.Println("backwards order:")
	for el := m.Last(); el != nil; el = el.Prev() {
		fmt.Printf("\t%s: %v\n", el.Key(), el.Value())
	}
	fmt.Println()

	fmt.Println(`forward from key "sub":`)
	for el := m.GetElement("sub"); el != nil; el = el.Next() {
		fmt.Printf("\t%s: %v\n", el.Key(), el.Value())
	}
	fmt.Println()

	// print map
	fmt.Println(m) // map[an:article empty:<nil> sub:map[s:1 e:2 x:3 y:4] bool:false array:[1 2 3]]

	// set new values, keeping order of existing keys
	m.Set("an", "bar")
	m.Set("truth", true)
	fmt.Println(m) // map[an:bar empty:<nil> sub:map[s:1 e:2 x:3 y:4] bool:false array:[1 2 3] truth:true]

	// delete key "sub"
	m.Delete("sub")
	fmt.Println(m) // map[an:bar empty:<nil> bool:false array:[1 2 3] truth:true]

	// update value for key "an", and move it to the end
	m.Push("an", "end")
	fmt.Println(m) // map[empty:<nil> bool:false array:[1 2 3] truth:true an:end]

	data, err := json.Marshal(&m)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data)) // {"empty":null,"bool":false,"array":[1,2,3],"truth":true,"an":"end"}
}

```

## Alternatives

* [iancoleman/orderedmap](https://github.com/iancoleman/orderedmap) — has O(n) time for Delete
* [wk8/go-ordered-map](https://github.com/wk8/go-ordered-map) — Unmarshal creates nested maps as native unordered maps, which makes it useless for my purposes

Let me know of other alternatives, I'll add them here
