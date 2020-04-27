package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonArrayParser_Parse(t *testing.T) {
	a := assert.New(t)

	meta := "json_array"
	testCases := []struct {
		name     string
		data     []byte
		nodes    Nodes
		errCheck func(error, ...interface{}) bool
	}{
		{
			name:     "empty data",
			data:     []byte(``),
			nodes:    nil,
			errCheck: a.Error,
		},
		{
			name: "json string",
			data: []byte(`{"a" : "b"}`),
			nodes: map[Key]Node{
				"a": NewNode([]byte(`"b""`), meta),
			},
			errCheck: a.NoError,
		},
		{
			name: "json obj",
			data: []byte(`{ "a" : {"foo" : "bar"}, "b": null}`),
			nodes: map[Key]Node{
				"a": NewNode([]byte(`{"foo":"bar"}`), meta),
				"b": NewNode([]byte(`null`), meta),
			},
			errCheck: a.NoError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			p := NewJSONArrayParser(meta)
			nodes, err := p.Parse(tc.data)
			tc.errCheck(err)
			if err != nil {
				a.Equal(tc.nodes, nodes)
			}
		})
	}
}
