package tree

import (
	"strconv"

	"github.com/pkg/errors"
)

func NewEscapedJSONParser() *escapedJsonParser {
	return &escapedJsonParser{
		meta: "unescape",
	}
}

type escapedJsonParser struct {
	meta string
}

func (p *escapedJsonParser) Parse(data []byte) (Nodes, error) {
	if len(data) < 2 {
		return nil, ErrCannotParse
	}
	s, err := strconv.Unquote(string(data))
	if err != nil {
		return nil, errors.Wrap(err, "cannot unquote")
	}
	return map[Key]Node{
		"val": NewNode([]byte(s), p.meta),
	}, nil
}
