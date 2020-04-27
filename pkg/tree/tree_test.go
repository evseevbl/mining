package tree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTreeBuilder_BuildTree(t *testing.T) {
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
			name:     "simple json",
			data:     []byte(`{"foo":"bar"}`),
			nodes:    nil,
			errCheck: a.NoError,
		},
		{
			name:     "nested json",
			data:     []byte(`{"foo": "bar", "a": {"price": 100500}}`),
			nodes:    nil,
			errCheck: a.NoError,
		},
		{
			name:     "json as string",
			data:     []byte(`{"foo": "bar", "a": "{\"price\": 100500}" }`),
			nodes:    nil,
			errCheck: a.NoError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tb := NewTreeBuilder(
				NewJSONArrayParser(meta),
				NewEscapedJSONParser(),
			)
			tree, _ := tb.BuildTree(tc.data)
			fmt.Println(tree)
		})
	}
}
