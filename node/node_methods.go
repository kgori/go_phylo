package node

import "fmt"

// import "github.com/kgori/go_phylo/lexer"

//****************************
// node methods
//****************************

func (parent *node) AddChild() *node {
	nodelet := New()
	nodelet.pointsDown = false
	child := New()
	parent.isLeaf = false

	// sync nodelets to nodeData
	parent.next, nodelet.next = nodelet, parent.next
	nodelet.nodeData = parent.nodeData

	// sync parent-child edgeData
	nodelet.out, child.out = child, nodelet
	edgeData := newEdgeData()
	nodelet.edgeData = edgeData
	child.edgeData = edgeData

	return child
}

func (n *node) AddBranchLength(branchLength float64) {
	n.edgeData.branchLength = branchLength
}

func (n *node) GetBranchLength() float64 {
	return n.edgeData.branchLength
}

func (n *node) AddSupportValue(supportValue float64) {
	n.edgeData.supportValue = supportValue
}

func (n *node) GetSupportValue() float64 {
	return n.edgeData.supportValue
}

func (n *node) AddNodeLabel(s string) {
	n.nodeData.label = s
}

func (n *node) String() string {
	return n.nodeData.label
}

func (n *node) AddMetadata(key, value string) {
	n.nodeData.metadata[key] = value
}

func (n *node) LookupMetadata(key string) string {
	return n.nodeData.metadata[key]
}

func (n *node) PostorderTraverse() *node {
	curr := n.next
Loop:
	for {
		if curr.pointsDown {
			break Loop
		}
		curr.out.PostorderTraverse()
		curr = curr.next
	}
	fmt.Printf("%s\n", curr)
	return curr
}

func (n *node) PreorderTraverse() *node {
	curr := n.next
	fmt.Printf("%s\n", curr)
Loop:
	for {
		if curr.pointsDown {
			break Loop
		}
		curr.out.PreorderTraverse()
		curr = curr.next
	}
	return curr
}

// func (n *node) Parse(input string) {
// 	l := lexer.New(input)
// 	s := NewStack()
// Loop:
// 	for {
// 		switch i := l.NextItem; i.WhatIs() {
// 		case ItemLBracket:
// 			break // new inner node
// 		case ItemRBracket:
// 			break // close inner node
// 		case ItemLabel:
// 			break // new leaf node
// 		case ItemComment:
// 			break
// 		case ItemNHXKey:
// 			break
// 		case ItemSupportValue:
// 			break
// 		case ItemBranchLength:
// 			break
// 		case ItemError:
// 			break
// 		case ItemEOF:
// 			break
// 		}

// 	}
// }
