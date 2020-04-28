package tree

import (
	"encoding/base64"

	"github.com/pkg/errors"
)

func NewBase64Decoder() *b64decoder {
	return &b64decoder{
		meta: "b64_decode",
	}
}

type b64decoder struct {
	meta string
}

func (p *b64decoder) Parse(data []byte) (Nodes, error) {
	if len(data) == 0 {
		return nil, ErrCannotParse
	}
	decoded , err := base64.StdEncoding.DecodeString(string(data))
	if err != nil {
		return nil, errors.Wrap(err, "cannot decode")
	}

	return map[Key]Node{
		"decoded": NewNode([]byte(string(decoded)), p.meta),
	}, nil
}
