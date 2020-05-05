package decoders

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/pkg/errors"
)

func UrlDecode() *urlDecoder {
	return &urlDecoder{
		name: "url_decode",
	}
}

type urlDecoder struct {
	name
}

func (p *urlDecoder) Decode(data []byte) (map[string][]byte, error) {
	str := string(data)
	if !strings.ContainsAny(str, "&=") {
		return nil, errors.New("nothing to decode")
	}

	ret, err := url.ParseQuery(str)
	if err != nil {
		return nil, errors.Wrap(err, "parseQuery")
	}

	results := make(map[string][]byte)
	for key, values := range ret {
		if len(values) == 1 {
			results[key] = []byte(values[0])
			continue
		}
		for i, val := range values {
			results[fmt.Sprintf("%s_%02d", key, i)] = []byte(val)
		}
	}
	return results, nil
}
