package tree

import (
	"encoding/json"

	"github.com/pkg/errors"
)

func NewJSONObjectParser() *jsonObjectParser {
	return &jsonObjectParser{
		meta: "json_object",
	}
}

type jsonObjectParser struct {
	meta string
}

func (j *jsonObjectParser) Parse(data []byte) (Nodes, error) {
	type obj map[string]any

	o := make(obj)
	if err := json.Unmarshal(data, &o); err != nil {
		return nil, errors.Wrap(ErrCannotParse, err.Error())
	}

	nodes := make(Nodes)
	for k, v := range o {
		nodes[Key(k)] = NewNode(v.data, j.meta)
	}
	return nodes, nil
}
