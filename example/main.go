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
