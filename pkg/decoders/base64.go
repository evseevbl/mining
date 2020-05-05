package decoders

import (
	"encoding/base64"

	"github.com/pkg/errors"
)

func AsBase64() *b64decoder {
	return &b64decoder{
		name: "b64_decode",
	}
}

type b64decoder struct {
	name
}

func (p *b64decoder) Decode(data []byte) (map[string][]byte, error) {
	if len(data) == 0 {
		return nil, ErrCannotParse
	}
	decoded, err := base64.StdEncoding.DecodeString(string(data))
	if err != nil {
		return nil, errors.Wrap(err, "cannot decode")
	}

	return map[string][]byte{
		value: decoded,
	}, nil
}
