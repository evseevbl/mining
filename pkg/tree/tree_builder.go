package tree

import (
	"fmt"

	"github.com/pkg/errors"
)

type decoder interface {
	Name() string
	Decode([]byte) (map[string][]byte, error)
}

type treeBuilder struct {
	decoders []decoder
}

func NewTreeBuilder(parsers ...decoder) *treeBuilder {
	return &treeBuilder{
		decoders: parsers,
	}
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
	successes := make([]*decoderResult, 0, len(tb.decoders))
	for _, p := range tb.decoders {
		values, err := p.Decode(n.Data())
		if err == nil {
			successes = append(successes, &decoderResult{
				values: values,
				name:   p.Name(),
			})
		}
	}
	switch len(successes) {
	case 0:
		return errors.New("all decoders failed")
	case 1: // unique interpretation
		n.SetChildren(successes[0].asNodes())
		for _, child := range n.GetChildren() {
			_ = tb.ProcessNode(child)
		}
		return nil
	default:
		decodeVariants := make(Nodes)
		for _, result := range successes {
			// same data as parent
			variant := NewNode(n.Data(), result.name)
			variant.SetChildren(result.asNodes())

			key := Key(fmt.Sprintf("_as_%s", result.name))
			decodeVariants[key] = variant
		}
		n.SetChildren(decodeVariants)
		// first level nodes already have results
		for _, processedNode := range n.GetChildren() {
			for _, unprocessedNode := range processedNode.GetChildren() {
				_ = tb.ProcessNode(unprocessedNode)
			}
		}
		return nil
	}
	return errors.New("all decoders failed")
}

type decoderResult struct {
	values map[string][]byte
	name   string
}

func (dr *decoderResult) asNodes() Nodes {
	nodes := make(Nodes)
	for k, v := range dr.values {
		nodes[Key(k)] = NewNode(v, dr.name)
	}
	return nodes
}
