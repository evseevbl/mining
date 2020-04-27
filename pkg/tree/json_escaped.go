package tree

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
)

func NewEscapedJSONParser() *escapedJsonParser {
	return &escapedJsonParser{
		meta: "json_as_string",
		arr:  NewJSONArrayParser("json_array"),
	}
}

type escapedJsonParser struct {
	meta string
	arr  Parser
}

func (p *escapedJsonParser) Parse(data []byte) (Nodes, error) {
	if len(data) < 2 {
		return nil, ErrCannotParse
	}
	// data = data[1 : len(data)-1]
	s, err := strconv.Unquote(string(data))
	if err != nil {
		return nil, errors.Wrap(err, "cannot unquote")
	}
	fmt.Printf("<%s>:<%s>\n", string(data), s)
	return map[Key]Node{
		"0": NewNode([]byte(s), "escaped_json") ,
	}, nil
}
