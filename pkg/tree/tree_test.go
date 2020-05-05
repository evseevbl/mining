package tree

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/evseevbl/mining/pkg/decoders"
)

func TestTreeBuilder_BuildTree(t *testing.T) {
	a := assert.New(t)

	testCases := []struct {
		name        string
		data        []byte
		exp         tree
		allDecoders []decoder
		errCheck    func(error, ...interface{}) bool
	}{
		{
			name: "empty data",
			data: []byte(``),
			allDecoders: []decoder{
				decoders.AsBase64(),
				NewDummyDecoder("someName", nil, someErr),
			},
			errCheck: a.Error,
		},
		{
			name:     "simple json",
			data:     []byte(`{"foo":"bar"}`),
			errCheck: a.NoError,
			allDecoders: []decoder{
				decoders.AsJSONObject(),
			},
		},
		{
			name:     "nested json",
			data:     []byte(`{"foo": "bar", "a": {"price": 100500}}`),
			errCheck: a.NoError,
			allDecoders: []decoder{
				decoders.AsJSONObject(),
			},
		},
		{
			name: "json as string",
			data: []byte(`{"foo": "bar", "a": "{\"price\": 100500}" }`),
			allDecoders: []decoder{
				decoders.AsJSONObject(),
				decoders.Unescape(),
			},
			errCheck: a.NoError,
		},
		{
			name: "example from article",
			data: []byte(` c=0&load%5B%5D=jquery-core,jquery-migrate&load%5B%5D=utils&ver=3.8.2&json={%22firstName%22:%22Иван%22,%22lastName%22:%22Иванов%22,%22address%22:{%22postalCode%22:101101},%22phoneNumbers%22:[%22812123-1234%22,%22916123-4567%22]}`),
			allDecoders: []decoder{
				decoders.AsBase64(),
				decoders.AsJSONObject(),
				decoders.Unescape(),
				decoders.UrlDecode(),
				decoders.AsJSONArray(),
			},
			errCheck: a.NoError,
		},
		{
			name: "agar.io"		,
			data: []byte(`type=userAction&data=eyJjaGFubmVsIjoid2ViX3dpZGdldCIsInVzZXJBY3Rpb24iOnsiY2F0ZWdvcnkiOiJhcGkiLCJhY3Rpb24iOiJ6RS5oaWRlIiwibGFiZWwiOm51bGwsInZhbHVlIjp7ImFyZ3MiOm51bGx9fSwiYnVpZCI6Ijg2NzliZjEyMmYxMDE4YzkxZWVmNTBmYThlY2Q3ZWIwIiwic3VpZCI6IjdkZThkNzdlNTZiMjZjYTJhN2E4ZTlkMzg0ZjI0YTA2IiwidmVyc2lvbiI6IjQ3OGY0YTA0ZiIsInRpbWVzdGFtcCI6IjIwMjAtMDQtMjhUMDk6MTk6MjUuMDQ1WiIsInVybCI6Imh0dHBzOi8vYWdhci5pby8jZmZhIn0%3D`),
			allDecoders: []decoder{
				decoders.AsJSONObject(),
				decoders.AsBase64(),
				decoders.UrlDecode(),
				decoders.AsJSONArray(),
			},
			errCheck:a.NoError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tb := NewTreeBuilder(
				tc.allDecoders...
			)
			tree, err := tb.BuildTree(tc.data)
			if err == nil {
				tree.PrettyPrint()
			}
		})
	}
}

var (
	someResult = map[string][]byte{
		"foo": []byte(`"bar"`),
	}

	someErr = errors.New("mock error")
)

func NewDummyDecoder(name string, results map[string][]byte, err error) *dummyDecoder {
	return &dummyDecoder{
		name:    name,
		results: results,
		err:     err,
	}
}

type dummyDecoder struct {
	name    string
	results map[string][]byte
	err     error
}

func (d *dummyDecoder) Name() string {
	return d.name
}

func (d *dummyDecoder) Decode(_ []byte) (map[string][]byte, error) {
	return d.results, d.err
}
