package tree

import (
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
		{
			name:     "example from article",
			data:     []byte(` c=0&load%5B%5D=jquery-core,jquery-migrate&load%5B%5D=utils&ver=3.8.2&json={%22firstName%22:%22Иван%22,%22lastName%22:%22Иванов%22,%22address%22:{%22postalCode%22:101101},%22phoneNumbers%22:[%22812123-1234%22,%22916123-4567%22]}`),
			nodes:    nil,
			errCheck: a.NoError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tb := NewTreeBuilder(
				NewJSONArrayParser(meta),
				NewEscapedJSONParser(),
				NewUrlDecoder(),
			)
			tree, _ := tb.BuildTree(tc.data)
			_ = tree
		})
	}
}
