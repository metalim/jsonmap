package main

import (
	"encoding/json"
	"fmt"

	// "github.com/metalim/jsonmap"
	// or simpler alternative, but with O(n) Delete()
	jsonmap "github.com/metalim/jsonmap/simplemap"
)

const input = `{"an":"article","empty":null,"sub":{"s":1,"e":2,"x":3,"y":4},"bool":false,"array":[1,2,3]}`

func main() {
	m := jsonmap.New()

	// unmarshal, keeping order
	err := json.Unmarshal([]byte(input), &m)
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

	if string(output) == input {
		fmt.Println("output == input")
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
