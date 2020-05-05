package decoders

import (
	"github.com/pkg/errors"
)

type DecodeResult struct {
	Values map[string][]byte
	Meta   string
}

var (
	ErrCannotParse = errors.New("cannot parse")
)
