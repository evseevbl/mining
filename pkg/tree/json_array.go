package tree

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/pkg/errors"
)

func NewJSONListParser() *jsonArrayParser {
	return &jsonArrayParser{
		meta: "json_array",
	}
}

type jsonArrayParser struct {
	meta string
}

func (j *jsonArrayParser) Parse(data []byte) (Nodes, error) {
	type obj struct {
		Array []any `json:"array"`
	}

	data = []byte(fmt.Sprintf(`{"array": %s}`, string(data)))

	o := new(obj)
	if err := json.Unmarshal(data, &o); err != nil {
		return nil, errors.Wrap(ErrCannotParse, err.Error())
	}

	nodes := make(Nodes)
	for i, v := range o.Array {
		nodes[Key(strconv.Itoa(i))] = NewNode(v.data, j.meta)
	}
	return nodes, nil
}
