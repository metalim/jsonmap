package jsonmap

import (
	"bytes"
	"encoding/json"
)

func (m *Map) MarshalJSON() ([]byte, error) {
	// TODO: implement
	return json.Marshal(m.m)
}

func (m *Map) UnmarshalJSON(data []byte) error {
	r := json.NewDecoder(bytes.NewReader(data))
	m.m = make(map[string]any)

	err := r.Decode(&m.m)

	if err != nil {
		return err
	}

	// fake order
	m.keys = m.keys[:0]
	for k := range m.m {
		m.keys = append(m.keys, k)
	}

	return nil
}
