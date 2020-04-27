package tree

import (
	"github.com/pkg/errors"
)

type Key string

type Nodes map[Key]Node

type Node interface {
	Data() []byte
	SetChildren(children Nodes) Node
	GetChildren() Nodes
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

func (n *node) SetChildren(children Nodes) Node {
	n.children = children
	return n
}

func (n *node) GetChildren() Nodes {
	return n.children
}

type Parser interface {
	Parse([]byte) (Nodes, error)
}

var (
	ErrCannotParse = errors.New("cannot parse")
)
