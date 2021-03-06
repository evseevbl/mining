package decoders

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/evseevbl/mining/pkg/tree"
)

func TestUrlDecoder_Parse(t *testing.T) {
	a := assert.New(t)

	meta := "url_decode"
	testCases := []struct {
		name     string
		data     []byte
		nodes    tree.Nodes
		errCheck func(error, ...interface{}) bool
	}{
		{
			name: "with json",
			data: []byte(`c=0&load%5B%5D=jquery-core,jquery-migrate&load%5B%5D=utils&ver=3.8.2&json={%22name%22:%22Иван%20Иванов%22}`),
			nodes: map[tree.Key]tree.Node{
				"c": tree.NewNode([]byte(`0`), meta),
			},
			errCheck: a.NoError,
		},
	}

	p := UrlDecode()
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			nodes, err := p.Decode(tc.data)
			tc.errCheck(err)
			if err != nil {
				a.Equal(tc.nodes, nodes)
			}
		})
	}
}
