package tree

import (
	"fmt"
	"strings"
)

type tree struct {
	root Node
}

func (t *tree) PrettyPrint() {
	fmt.Println(string(t.root.Data()))
	t.printChildren(t.root, 0)
}

func (t *tree) printChildren(n Node, level int) {
	for key, child := range n.GetChildren() {
		fmt.Printf("%s-> %s:\t%s\t[%s]\n",
			strings.Repeat(" ", level),
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
