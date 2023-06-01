package jsonmap

import (
	"bytes"
	"encoding/json"
	"errors"
)

func (m *Map) UnmarshalJSON(data []byte) error {
	if m.m == nil {
		m.m = make(map[string]any)
	}

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

		tok, err = d.Token()
		if err != nil {
			return err
		}

		switch tok {

		case json.Delim('{'):
			m2 := New()
			err = decodeMap(d, m2)
			if err != nil {
				return err
			}
			m.Push(key, m2)

		case json.Delim('['):
			a := make([]any, 0)
			err = decodeArray(d, a)
			if err != nil {
				return err
			}
			m.Push(key, a)

		default:
			m.Push(key, tok)
		}
	}
}

func decodeArray(d *json.Decoder, a []any) error {
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}

		if tok == json.Delim(']') {
			return nil
		}

		switch tok {

		case json.Delim('{'):
			m := New()
			err = decodeMap(d, m)
			if err != nil {
				return err
			}
			a = append(a, m)

		case json.Delim('['):
			a2 := make([]any, 0)
			err = decodeArray(d, a2)
			if err != nil {
				return err
			}
			a = append(a, a2)

		default:
			a = append(a, tok)
		}
	}
}
