package tree

import (
	"encoding/json"

	"github.com/pkg/errors"
)

func NewJSONArrayParser(meta string) *jsonArrayParser {
	return &jsonArrayParser{
		meta: meta,
	}
}

type jsonArrayParser struct {
	meta string
}

type any struct {
	data []byte
}

func (a *any) UnmarshalJSON(data []byte) error {
	a.data = data
	return nil
}

func (j *jsonArrayParser) Parse(data []byte) (Nodes, error) {
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
