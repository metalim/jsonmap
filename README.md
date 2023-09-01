# Ordered map

[![go test workflow](https://github.com/metalim/jsonmap/actions/workflows/gotest.yml/badge.svg)](https://github.com/metalim/jsonmap/actions/workflows/gotest.yml)

Simple ordered map for Go, with JSON restrictions. The main purpose is to keep same order of keys after parsing JSON and generating it again, so Unmarshal followed by Marshal generates exactly the same JSON structure

Keys are strings, Values are any JSON values (number, string, boolean, null, array, map/object)

Storage is O(n), operations are O(1), except for optional operations in [slow.go](slow.go) file

When Unmarshalling, **any nested map from JSON is created as ordered jsonmap**, including maps in nested arrays

Alternative implementation is in [simplemap](simplemap), it has simple structure, O(1) Keys, but O(n) Delete operation. Check the difference by running `go test -bench . -benchmem`

Inspired by [wk8/go-ordered-map](https://github.com/wk8/go-ordered-map) and [iancoleman/orderedmap](https://github.com/iancoleman/orderedmap)

## Performance

Similar to Go native map, `jsonmap` has O(1) time for Get, Set, Delete. Additionally it has Push, First, Last, Next, Prev operations, which are also O(1)

```
➜ go test -bench . -benchmem
...

Benchmark/Ops/Get/jsonmap-10           10133562               118.1 ns/op             0 B/op          0 allocs/op
Benchmark/Ops/Get/simplemap-10         11165646               110.1 ns/op             0 B/op          0 allocs/op
Benchmark/Ops/Get/gomap-10             10916256               110.7 ns/op             0 B/op          0 allocs/op

Benchmark/Ops/SetExisting/jsonmap-10    3813949               278.3 ns/op            25 B/op          1 allocs/op
Benchmark/Ops/SetExisting/simplemap-10  4422174               239.2 ns/op            37 B/op          0 allocs/op
Benchmark/Ops/SetExisting/gomap-10      7771100               144.1 ns/op             7 B/op          0 allocs/op

Benchmark/Ops/SetNew/jsonmap-10         3636994               292.2 ns/op            55 B/op          1 allocs/op
Benchmark/Ops/SetNew/simplemap-10       4094977               267.8 ns/op           131 B/op          1 allocs/op
Benchmark/Ops/SetNew/gomap-10           6028083               168.5 ns/op            15 B/op          1 allocs/op

Benchmark/Ops/Delete/jsonmap-10         5974435               192.0 ns/op             0 B/op          0 allocs/op
Benchmark/Ops/Delete/simplemap-10           140           8937567 ns/op               0 B/op          0 allocs/op
Benchmark/Ops/Delete/gomap-10           9750571               114.8 ns/op             0 B/op          0 allocs/op
```

Suite benchmark does 10k of Set and Get, and 1k of Delete operations. `jsonmap` performance is on par with native Go map. `simplemap` is slower because of O(N) Delete

```
Benchmark/Suite/jsonmap-10                    33          33939048 ns/op        24056151 B/op     621895 allocs/op
Benchmark/Suite/simplemap-10                   2         922005438 ns/op        32889740 B/op     521865 allocs/op
Benchmark/Suite/gomap-10                      38          30634451 ns/op        23947561 B/op     521837 allocs/op
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
	// or simpler alternative, but with O(n) Delete()
	// jsonmap "github.com/metalim/jsonmap/simplemap"
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

* [iancoleman/orderedmap](https://github.com/iancoleman/orderedmap) — has O(n) time for Delete, similar to `simplemap`
* [wk8/go-ordered-map](https://github.com/wk8/go-ordered-map) — Unmarshal creates nested maps as native unordered maps, which makes it useless for my purposes

Let me know of other alternatives, I'll add them here
