# Ordered map

![go test workflow](https://github.com/metalim/jsonmap/actions/workflows/gotest.yml/badge.svg)

Simple ordered map for Go, with JSON restrictions. The main purpose is to keep same order of keys after parsing JSON and generating it again, so Unmarshal followed by Marshal generates exactly the same JSON structure

Keys are strings, Values are any JSON values (number, string, boolean, null, array, map/object)

Storage is O(N), operations are O(1), except Delete, which is O(N) time. Delete can be reimplemented in O(log(N)) time with additional O(N) storage, or in O(1) time with more complicated data structures, but let's keep it simple for now

When Unmarshalling, **any nested map from JSON is created as ordered**, including maps in nested arrays

Inspired by https://github.com/iancoleman/orderedmap

## Usage

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/metalim/jsonmap"
)

const input = `{"an":"article","empty":null,"sub":{"x":1,"y":2},"bool":false,"array":[1,2,3]}`

func main() {
	m := jsonmap.New()

	// unmarshal, keeping order
	err := json.Unmarshal([]byte(input), &m)
	if err != nil {
		panic(err)
	}

	fmt.Println(m.Keys())         // [an empty sub bool array]
	fmt.Println(m.Get("an"))      // article true
	fmt.Println(m.Get("nothing")) // <nil> false

	// marshal, keeping order
	output, err := json.Marshal(&m)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(output) == input) // true

	// set new values, keeping order of existing keys
	m.Set("an", "bar")
	m.Set("truth", true)
	fmt.Println(m.Keys()) // [an empty sub bool array truth]

	// delete key "sub"
	m.Delete("sub")
	fmt.Println(m.Keys()) // [an empty bool array truth]

	// update value for key "an", and move it to the end
	m.Push("an", false)
	fmt.Println(m.Keys()) // [empty bool array truth an]

	data, err := json.Marshal(&m)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data)) // {"empty":null,"bool":false,"array":[1,2,3],"truth":true,"an":false}
}

```

## Alternatives

* https://github.com/iancoleman/orderedmap — also has O(N) time for Delete, but my implementation is cleaner
* https://github.com/wk8/go-ordered-map — has O(1) time for Delete due to linked list storage, but Unmarshal creates nested maps as vanilla unordered maps, which makes it useless for my purposes

Let me know of other alternatives, I'll add them here
