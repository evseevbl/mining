package tree

import (
	"fmt"
	"strings"

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

func (t *tree) PrettyPrint() {
	fmt.Println(string(t.root.Data()))
	t.printChildren(t.root, 0)
}

func (t *tree) printChildren(n Node, level int) {
	for key, child := range n.GetChildren() {
		fmt.Printf("%s> %s= %s [%s]\n",
			strings.Repeat("--", level),
			key,
			string(t.shorten(child.Data())),
			child.Meta(),
		)
		t.printChildren(child, level+1)
	}
}

func (tree) shorten(b []byte) []byte {
	switch {
	case len(b) == 0:
		return []byte{}
	case len(b) < 30:
		return b
	default:
		s := []byte(string(b[0:10]))
		s = append(s, []byte(`...`)...)
		return append(s, b[len(b)-10:len(b)-1]...)
	}
}
