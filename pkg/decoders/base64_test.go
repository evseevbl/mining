package decoders

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestB64decoder_Parse(t *testing.T) {
	a := assert.New(t)

	testCases := []struct {
		name     string
		data     []byte
		nodes    map[string][]byte
		errCheck func(error, ...interface{}) bool
	}{
		{
			name: "simple data",
			data: []byte(`eyJ1c2VyX2lkIjogMTAwNTAwfQ==`),
			nodes: map[string][]byte{
				value: []byte(`{"user_id": 100500}`),
			},
			errCheck: a.NoError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			p := AsBase64()
			ret, err := p.Decode(tc.data)
			tc.errCheck(err)
			a.Equal(tc.nodes, ret)
		})
	}
}
