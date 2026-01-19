package test_test

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"github.com/metalim/jsonmap"
	"github.com/zeebo/assert"
)

type test struct {
	name string
	json string
}

var tests = []test{
	{"PlainJSON", `{"a":1,"c":3,"d":4,"b":2,"e":5}`},
	{"NestedJSON", `{"1":1,"c":null,"d":"dd","b":true,"e":[2,null,{"x":1,"z":3,"y":2},"zzz"],"f":{"a":1,"c":3,"d":4,"b":2,"e":5}}`},
	{"EscapedJSON", `{"\"":"\"","\\":"\\"}`},
}

type IMapShort interface {
	Len() int
	Set(key Key, value Value)
	Get(key Key) (value Value, ok bool)
	Delete(key Key)
	Keys() []Key
}

func TestSerialization(t *testing.T) {
	t.Run("JSONMap", func(t *testing.T) {
		testSerialization(t, jsonmap.New)
	})
}

func testSerialization[T IMapShort](t *testing.T, newMap func() T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// due to random nature of Go map iteration, we need to test it many times
			for i := 0; i < TEST_ITERATIONS; i++ {
				m := newMap()
				testSerializationOnce(t, test.json, m)
				if test.name == "PlainJSON" {
					verifyPlainJSON(t, m)
				}
			}
		})
	}
}

func testSerializationOnce(t *testing.T, testData string, m IMapShort) {
	err := json.Unmarshal([]byte(testData), &m)
	assert.Equal(t, err, nil)

	data, err := json.Marshal(m)
	assert.Equal(t, err, nil)
	assert.Equal(t, string(data), testData)
}

func verifyPlainJSON(t *testing.T, m IMapShort) {
	assert.Equal(t, m.Len(), 5)

	// keys
	assert.Equal(t, len(m.Keys()), 5)
	assert.Equal(t, strings.Join(m.Keys(), ","), "a,c,d,b,e")

	// values
	v, ok := m.Get("a")
	assert.True(t, ok)
	assert.Equal(t, v.(float64), 1.)
	v, ok = m.Get("c")
	assert.True(t, ok)
	assert.Equal(t, v.(float64), 3.)
	v, ok = m.Get("d")
	assert.True(t, ok)
	assert.Equal(t, v.(float64), 4.)
	v, ok = m.Get("b")
	assert.True(t, ok)
	assert.Equal(t, v.(float64), 2.)
	v, ok = m.Get("e")
	assert.True(t, ok)
	assert.Equal(t, v.(float64), 5.)
	v, ok = m.Get("f")
	assert.False(t, ok)
	assert.Equal(t, v, nil)
}

func encodeJSON(v any, escapeHTML bool) (string, error) {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(escapeHTML)
	err := enc.Encode(v)
	if err != nil {
		return "", err
	}
	return string(bytes.TrimRight(buf.Bytes(), "\n")), nil
}

func TestHTMLEscaping(t *testing.T) {
	const rawJSON = `{"range":">1.0.0 && <2.0.0"}`
	const escapedJSON = `{"range":"\u003e1.0.0 \u0026\u0026 \u003c2.0.0"}`
	const nestedRawJSON = `{"name":"test","deps":{"a":">1.0","b":"<2.0 & >=1.5"},"tags":["<root>","&"]}`
	const nestedArrayJSON = `{"name":"test","list":[{"op":">="},{"op":"<="}]}`

	t.Run("DefaultMarshal", func(t *testing.T) {
		m := jsonmap.New()
		err := json.Unmarshal([]byte(rawJSON), m)
		assert.NoError(t, err)

		data, err := json.Marshal(m)
		assert.NoError(t, err)
		assert.Equal(t, string(data), escapedJSON)
	})

	t.Run("DefaultEncode", func(t *testing.T) {
		m := jsonmap.New()
		err := json.Unmarshal([]byte(rawJSON), m)
		assert.NoError(t, err)

		data, err := encodeJSON(m, true)
		assert.NoError(t, err)
		assert.Equal(t, data, escapedJSON)
	})

	t.Run("EncodeNoEscape", func(t *testing.T) {
		m := jsonmap.New()
		m.SetEscapeHTML(false)
		err := json.Unmarshal([]byte(rawJSON), m)
		assert.NoError(t, err)

		data, err := encodeJSON(m, false)
		assert.NoError(t, err)
		assert.Equal(t, data, rawJSON)
	})

	t.Run("EncodeNoEscapeNested", func(t *testing.T) {
		m := jsonmap.New()
		m.SetEscapeHTML(false)
		err := json.Unmarshal([]byte(nestedRawJSON), m)
		assert.NoError(t, err)

		data, err := encodeJSON(m, false)
		assert.NoError(t, err)
		assert.Equal(t, data, nestedRawJSON)
	})

	t.Run("EncodeNoEscapeNestedArray", func(t *testing.T) {
		m := jsonmap.New()
		m.SetEscapeHTML(false)
		err := json.Unmarshal([]byte(nestedArrayJSON), m)
		assert.NoError(t, err)

		data, err := encodeJSON(m, false)
		assert.NoError(t, err)
		assert.Equal(t, data, nestedArrayJSON)
	})
}
