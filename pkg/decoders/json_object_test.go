package decoders

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonArrayParser_Parse(t *testing.T) {
	a := assert.New(t)

	testCases := []struct {
		name     string
		data     []byte
		values   map[string][]byte
		errCheck func(error, ...interface{}) bool
	}{
		{
			name:     "empty data",
			data:     []byte(``),
			values:   nil,
			errCheck: a.Error,
		},
		{
			name: "json string",
			data: []byte(`{"a" : "b"}`),
			values: map[string][]byte{
				"a": []byte(`"b""`),
			},
			errCheck: a.NoError,
		},
		{
			name: "json obj",
			data: []byte(`{ "a" : {"foo" : "bar"}, "b": null}`),
			values: map[string][]byte{
				"a": []byte(`{"foo":"bar"}`),
				"b": []byte(`null`),
			},
			errCheck: a.NoError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			p := AsJSONObject()
			nodes, err := p.Decode(tc.data)
			tc.errCheck(err)
			if err != nil {
				a.Equal(tc.values, nodes)
			}
		})
	}
}
