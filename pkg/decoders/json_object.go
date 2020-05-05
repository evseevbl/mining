package decoders

import (
	"encoding/json"

	"github.com/pkg/errors"
)

func AsJSONObject() *jsonObjectDecoder {
	return &jsonObjectDecoder{
		name: "json_object",
	}
}

type jsonObjectDecoder struct {
	name
}

func (j *jsonObjectDecoder) Decode(data []byte) (map[string][]byte, error) {
	type obj map[string]any

	o := make(obj)
	if err := json.Unmarshal(data, &o); err != nil {
		return nil, errors.Wrap(ErrCannotParse, err.Error())
	}

	values := make(map[string][]byte)
	for k, v := range o {
		values[k] = v.data
	}
	return values, nil
}
