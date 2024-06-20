// jsonmap is an ordered map. Same as native Go map, but keeps order of insertion when iterating
// or serializing to JSON, and with additional methods to iterate from any point.
// Similar to native map, user has to take care of concurrency and nil map value.
//
// Create new map:
//
//	m := jsonmap.New()
//
// Set values (remembering order of insertion):
//
//	m.Set("someString", "value1")
//	m.Set("someNumber", 42.1)
//	m.Set("someArray", []string{"a", "b", "c"})
//
// Get value:
//
//	value, ok := m.Get("someString")
//	str, ok := value.(string) // value can be any type
//	fmt.Println(str, ok)
//
// Get value as specific type:
//
//	str, ok := jsonmap.GetAs[string](m, "someString")
//	fmt.Println(str, ok) // str is string
//
// Iterate forwards:
//
//	for elem := m.First(); elem != nil; elem = elem.Next() {
//	    fmt.Println(elem.Key(), elem.Value())
//	}
//
// Iterate backwards:
//
//	for elem := m.Last(); elem != nil; elem = elem.Prev() {
//	    fmt.Println(elem.Key(), elem.Value())
//	}
//
// Iterate from any point:
//
//	for elem := m.GetElement("someNumber"); elem != nil; elem = elem.Next() {
//	    fmt.Println(elem.Key(), elem.Value())
//	}
//
// Delete element:
//
//	m.Delete("someNumber")
//
// Push element to the end:
//
//	m.Push("someNumber", 42.1)
//
// Pop element from the end:
//
//	key, value, ok := m.Pop()
//
// Clear map:
//
//	m.Clear()
//
// Serialize to JSON:
//
//	data, err := json.Marshal(m)
//
// Deserialize from JSON:
//
//	err = json.Unmarshal(data, &m)
//
// or use jsonmap.Map as a field in a struct:
//
//	type MyStruct struct {
//	    SomeMap *jsonmap.Map `json:"someMap"`
//	}
//
// And serialize/deserialize the struct:
//
//	data, err := json.Marshal(myStruct)
//	err = json.Unmarshal(data, &myStruct)
//
// Time complexity of operations:
//
//	| Operation | Time        |
//	|-----------|-------------|
//	| Clear     | O(1)        |
//	| Get       | O(1)        |
//	| Set       | O(1)        |
//	| Delete    | O(1)        |
//	| Push      | O(1)        |
//	| Pop       | O(1)        |
//	|           |             |
//	| First     | O(1)        |
//	| Last      | O(1)        |
//	| GetElement| O(1)        |
//	| el.Next   | O(1)        |
//	| el.Prev   | O(1)        |
//	|           |             |
//	| SetFront  | O(1)        |
//	| PushFront | O(1)        |
//	| PopFront  | O(1)        |
//	|           |             |
//	| KeyIndex  | O(N)        |
//	| Keys      | O(N)        |
//	| Values    | O(N)        |
//	| SortKeys  | O(N*log(N)) |
package jsonmap
