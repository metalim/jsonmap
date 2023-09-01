// simplemap is a simpler implementation of jsonmap, but with O(N) Delete operation.
// Here's comparison of time complexity of operations, with differences marked with __bold__:
//
//	| Operation | jsonmap     | simplemap   |
//	|-----------|-------------|-------------|
//	| Clear     | O(1)        | O(1)        |
//	| Get       | O(1)        | O(1)        |
//	| Set       | O(1)        | O(1)        |
//	| Delete    | O(1)        | __O(N)__    |
//	| Push      | O(1)        | O(1)        |
//	| Pop       | O(1)        | O(1)        |
//	|           |             |             |
//	| First     | O(1)        | O(1)        |
//	| Last      | O(1)        | O(1)        |
//	| GetElement| O(1)        | __O(N)__    |
//	| el.Next   | O(1)        | O(1)        |
//	| el.Prev   | O(1)        | O(1)        |
//	|           |             |             |
//	| SetFront  | O(1)        | __O(N)__    |
//	| PushFront | O(1)        | __O(N)__    |
//	| PopFront  | O(1)        | __O(N)__    |
//	|           |             |             |
//	| KeyIndex  | O(N)        | O(N)        |
//	| Keys      | O(N)        | __O(1)__    |
//	| Values    | O(N)        | O(N)        |
//	| SortKeys  | O(N*log(N)) | O(N*log(N)) |
package jsonmap
