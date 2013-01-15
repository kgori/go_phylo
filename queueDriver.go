package main

import (
	"fmt"
	"github.com/kgori/go_phylo/node"
)

func main() {
	q := node.NewQueue(3)
	m := node.New()
	m.AddNodeLabel("m")
	n := node.New()
	n.AddNodeLabel("n")
	o := node.New()
	o.AddNodeLabel("o")
	q.Remove()
	q.Add(m)
	q.Add(o)
	q.Add(n)
	q.Add(o)
	One := q.Remove()
	Two := q.Remove()
	Thr := q.Remove()
	fmt.Println(One)
	fmt.Println(Two)
	fmt.Println(Thr)
}
