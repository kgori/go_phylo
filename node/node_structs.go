package node

//****************************
// node and associated structs
// with constructors
//****************************

/*
Implements Felsenstein's Unrooted Tree Data Structure
from Inferring Phylogenies (2004) p. 589 
*/
type node struct {
	next            *node
	out             *node
	isLeaf          bool
	pointsDown      bool
	outPointsToRoot bool
	nodeData        *nodeData
	edgeData        *edgeData
}

type nodeData struct {
	label    string
	metadata map[string]string
}

type edgeData struct {
	branchLength float64
	supportValue float64
}

func New() *node {
	n := &node{
		isLeaf:     true,
		pointsDown: true,
	}
	n.next = n
	n.nodeData = newNodeData()
	return n
}

func newNodeData() *nodeData {
	nD := &nodeData{metadata: make(map[string]string)}
	return nD
}

func newEdgeData() *edgeData {
	eD := &edgeData{}
	return eD
}
