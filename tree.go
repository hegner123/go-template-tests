package main

import (
	"fmt"
	"strings"

	tdb "github.com/hegner123/go-template-tests/sqlc"
)

type Scanner func(parent, child tdb.Page) int

type Tree struct {
	fn    Scanner
	Label string
	Root  *Node
}

type Node struct {
	Val     tdb.Page
	Juniors []*Node
}

func Scan(parent, child tdb.Page) int {
	if parent.PageID == child.ParentID.Int32 {
		return 1
	} else {
		return 0
	}
}

func NewTree(fn Scanner, label string) *Tree {
	return &Tree{fn: fn, Label: label}
}

func (tr *Tree) Add(page tdb.Page) {
	tr.Root = tr.Root.Add(tr.fn, page)
}

func (n *Node) Add(fn Scanner, page tdb.Page) *Node {
	if n == nil {
		return &Node{Val: page}
	}
	switch r := fn(n.Val, page); {
	case r == 1:
		n.Juniors = append(n.Juniors, &Node{Val: page})
	case r == 0:
		for i := 0; i < len(n.Juniors); i++ {
			n.Juniors[i].Add(fn, page)
		}
	}

	return n
}

func FormatJson(s string) string {
	return strings.TrimSpace(s)
}

func (n *Node) PrintTree(level int) {
	if n == nil {
		return
	}
	indent := ""
	for i := 0; i < level; i++ {
		indent += "  "
	}
	fmt.Printf("%s%d (%s)\n", indent, n.Val.PageID, n.Val.Label)
	for _, junior := range n.Juniors {
		junior.PrintTree(level + 1)
	}
}
