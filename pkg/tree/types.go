package tree

type Key string

type Nodes map[Key]Node

type Node interface {
	Data() []byte
	Meta() string
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

// Meta contains information about the decoder
func (n *node) Meta() string {
	return n.meta
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

