package jsonmap

import "encoding/json"

func (m *Map) MarshalJSON() ([]byte, error) {
	// TODO: implement
	return json.Marshal(m.m)
}
