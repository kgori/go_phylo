package main

import (
	"fmt"
	"github.com/kgori/go_phylo/node"
)

func main() {
	// imagine constructing this tree:
	// ((A:0.1, B:0.2)95:0.3, C:0.25, D:0.15);
	tree := node.New()
	ABclade := tree.AddChild()
	A := ABclade.AddChild()
	B := ABclade.AddChild()
	C := tree.AddChild()
	D := tree.AddChild()
	ABclade.AddBranchLength(0.3)
	ABclade.AddSupportValue(95)
	tree.AddNodeLabel("tree")
	A.AddNodeLabel("A")
	B.AddNodeLabel("B")
	C.AddNodeLabel("C")
	D.AddNodeLabel("D")
	ABclade.AddNodeLabel("ABclade")
	A.AddBranchLength(0.1)
	B.AddBranchLength(0.2)
	C.AddBranchLength(0.25)
	D.AddBranchLength(0.15)

	// tree._preorderTraverse()
	s := node.NewStack()
	s.Push(A)
	n := s.Pop()
	fmt.Println(n)
}
