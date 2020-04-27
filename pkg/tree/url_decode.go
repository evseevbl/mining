package tree

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/pkg/errors"
)

func NewUrlDecoder() *urlDecoder {
	return &urlDecoder{
		meta: "url_decode",
	}
}

type urlDecoder struct {
	meta string
}

func (p *urlDecoder) Parse(data []byte) (Nodes, error) {
	str := string(data)
	if !strings.ContainsAny(str, "&=") {
		return nil, errors.New("nothing to decode")
	}


	ret, err := url.ParseQuery(str)
	if err != nil {
		return nil, errors.Wrap(err, "parseQuery")
	}

	nodes := make(Nodes)
	for key, values := range ret {
		if len(values) == 1 {
			nodes[Key(key)] = NewNode([]byte(values[0]), p.meta)
			continue
		}
		for i, val := range values {
			nodeKey := fmt.Sprintf("%s_%02d", key, i)
			nodes[Key(nodeKey)] = NewNode([]byte(val), p.meta)
		}
	}
	return nodes, nil
}
