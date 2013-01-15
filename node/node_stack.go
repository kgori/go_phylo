package node

type stack struct {
	array  [10]*node
	slice  []*node
	length int
}

type Stack interface {
	Push(n *node)
	Pop() *node
	Delete()
}

func NewStack() Stack {
	s := &stack{}
	return s
}

func (s *stack) Push(n *node) {
	s.slice = append(s.slice, n)
	s.length++
}
func (s *stack) Pop() *node {
	node := s.slice[s.length-1]
	s.Delete()
	return node
}
func (s *stack) Delete() {
	s.slice = s.slice[:s.length-1]
	s.length--
}
