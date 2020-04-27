package tree

import (
	"github.com/pkg/errors"
)

type Key string

type Nodes map[Key]Node

type Node interface {
	Data() []byte
	// WithResults(children Nodes) Node
}

type node struct {
	data     []byte
	meta     string
	children Nodes
}

func (n *node) Data() []byte {
	return n.data
}

func NewNode(data []byte, meta string) *node {
	return &node{
		data:     data,
		meta:     meta,
		children: nil,
	}
}

func (n *node) SetResults(children Nodes) {
	return
}

type Parser interface {
	Parse([]byte) (Nodes, error)
}

var (
	ErrCannotParse = errors.New("cannot parse")
)
