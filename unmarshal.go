package jsonmap

import (
	"bytes"
	"encoding/json"
	"errors"
)

// UnmarshalJSON implements json.Unmarshaler interface.
// It supports nested maps and arrays.
//
// Note: it does not clear the map before unmarshaling.
// If you want to clear it, call Clear() before UnmarshalJSON().
//
//	err := m.UnmarshalJSON([]byte(`{"a":1,"b":2}`))
func (m *Map) UnmarshalJSON(data []byte) error {
	d := json.NewDecoder(bytes.NewReader(data))
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if tok != json.Delim('{') {
		return errors.New("expected '{'")
	}

	return decodeMap(d, m)
}

func decodeMap(d *json.Decoder, m *Map) error {
	for {
		// key or end
		tok, err := d.Token()
		if err != nil {
			return err
		}

		if tok == json.Delim('}') {
			return nil
		}

		key, ok := tok.(string)
		if !ok {
			return errors.New("expected string key")
		}

		// value
		tok, err = d.Token()
		if err != nil {
			return err
		}

		switch tok {

		case json.Delim('{'):
			m2 := New()
			m2.SetEscapeHTML(m.escapeHTML)
			err = decodeMap(d, m2)
			if err != nil {
				return err
			}
			m.Push(key, m2)

		case json.Delim('['):
			a, err := decodeArray(d, m.escapeHTML)
			if err != nil {
				return err
			}
			m.Push(key, a)

		default:
			m.Push(key, tok)
		}
	}
}

func decodeArray(d *json.Decoder, escapeHTML bool) ([]any, error) {
	a := make([]any, 0)
	for {
		tok, err := d.Token()
		if err != nil {
			return a, err
		}

		switch tok {

		case json.Delim(']'):
			return a, nil

		case json.Delim('{'):
			m := New()
			m.SetEscapeHTML(escapeHTML)
			err = decodeMap(d, m)
			if err != nil {
				return a, err
			}
			a = append(a, m)

		case json.Delim('['):
			a2, err := decodeArray(d, escapeHTML)
			if err != nil {
				return a, err
			}
			a = append(a, a2)

		default:
			a = append(a, tok)
		}
	}
}
