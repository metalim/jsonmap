package jsonmap

import "encoding/json"

func (m *Map) MarshalJSON() ([]byte, error) {
	// TODO: implement
	return json.Marshal(m.m)
}

func (m *Map) UnmarshalJSON(data []byte) error {
	// TODO: implement
	return json.Unmarshal(data, &m.m)
}
