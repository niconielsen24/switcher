package logic

import "switcher/game"

type Queue = *Queue_s

type Queue_s struct {
	length int
	head   *Node
}

type Node struct {
	Value game.Tile
	Next  *Node
}

func NewQueue() Queue {
	return nil
}

func newNode(value game.Tile) *Node {
	return &Node{Value: value}
}

func (q Queue) Enqueue(value game.Tile) Queue {
	node := newNode(value)
	if q == nil {
		q = &Queue_s{}
		q.length = 1
		q.head = node
		return q
	}

	if q.head == nil {
		q.head = node
		q.length = 1
	} else {
		current := q.head
		for current.Next != nil {
			current = current.Next
		}
		q.length++
		current.Next = node
	}

	return q
}

func (q Queue) Dequeue() (game.Tile, Queue) {
	if q == nil {
		var zeroValue game.Tile
		return zeroValue, nil
	}
	value := q.head.Value
	q.head = q.head.Next
	q.length--
	return value, q
}
