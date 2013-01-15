package node

import "log"

type queue struct {
	channel chan *node
}

type Queue interface {
	Add(n *node)
	Remove() *node
}

func NewQueue(size int) Queue {
	c := make(chan *node, size)
	q := &queue{
		channel: c,
	}
	return q
}

func (q *queue) Add(n *node) {
	select {
	case q.channel <- n:
	default:
	}
}

func (q *queue) Remove() *node {
	select {
	case n := <-q.channel:
		return n
	default:
		log.Println("Removing from empty queue")
	}
	return nil
}
