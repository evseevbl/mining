package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestB64decoder_Parse(t *testing.T) {
	a := assert.New(t)

	meta := "b64_decode"
	testCases := []struct {
		name     string
		data     []byte
		nodes    Nodes
		errCheck func(error, ...interface{}) bool
	}{
		{
			name: "simple data",
			data: []byte(`eyJ1c2VyX2lkIjogMTAwNTAwfQ==`),
			nodes: map[Key]Node{
				"decoded": NewNode([]byte(`{"user_id": 100500}`), meta),
			},
			errCheck: a.NoError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			p := NewBase64Decoder()
			ret, err := p.Parse(tc.data)
			tc.errCheck(err)
			a.Equal(tc.nodes, ret)
		})
	}
}
