package decoders

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/pkg/errors"
)

func AsJSONArray() *jsonArrayDecoder {
	return &jsonArrayDecoder{
		name: "json_array",
	}
}

type jsonArrayDecoder struct {
	name
}

func (j *jsonArrayDecoder) Decode(data []byte) (map[string][]byte, error) {
	type obj struct {
		Array []any `json:"array"`
	}

	data = []byte(fmt.Sprintf(`{"array": %s}`, string(data)))

	o := new(obj)
	if err := json.Unmarshal(data, &o); err != nil {
		return nil, errors.Wrap(ErrCannotParse, err.Error())
	}

	values := make(map[string][]byte)
	for i, v := range o.Array {
		values[strconv.Itoa(i)] = v.data
	}
	return values, nil
}
