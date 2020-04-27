package tree

import (
	"github.com/pkg/errors"
)

type treeBuilder struct {
	parsers []Parser
}

func NewTreeBuilder(parsers ...Parser) *treeBuilder {
	return &treeBuilder{
		parsers: parsers,
	}
}

type tree struct {
	root Node
}

func (tb *treeBuilder) BuildTree(data []byte) (*tree, error) {
	t := &tree{
		root: NewNode(data, "root"),
	}

	if err := tb.ProcessNode(t.root); err != nil {
		return nil, errors.Wrap(err, "cannot process root node")
	}

	return t, nil
}

func (tb *treeBuilder) ProcessNode(n Node) error {
	for _, p := range tb.parsers {
		nodes, err := p.Parse(n.Data())
		if err == nil {
			n.SetChildren(nodes)
			for _, child := range n.GetChildren() {
				_ = tb.ProcessNode(child)
			}
			return nil
		}
	}
	return errors.New("all parsers failed")
}

