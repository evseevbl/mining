package decoders

import (
	"strconv"

	"github.com/pkg/errors"
)

func Unescape() *unescaper {
	return &unescaper{
		name: "unescape",
	}
}

type unescaper struct {
	name
}

func (p *unescaper) Decode(data []byte) (map[string][]byte, error) {
	if len(data) < 2 {
		return nil, ErrCannotParse
	}
	s, err := strconv.Unquote(string(data))
	if err != nil {
		return nil, errors.Wrap(err, "cannot unquote")
	}
	return map[string][]byte{
		value: []byte(s),
	}, nil
}
