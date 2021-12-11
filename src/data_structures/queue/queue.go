package queue

import "fmt"

type Node struct {
	prev *Node
	next *Node
	key  interface{}
}

type Queue struct {
	first  *Node
	last   *Node
	length int
}

func (Q *Queue) InitQueue(Key interface{}) {
	newNode := &Node{
		prev: nil,
		next: nil,
		key:  Key,
	}
	Q.first = newNode
	Q.last = newNode
	Q.length = 1
	Q.PrintQueue()
}

func (Q *Queue) Enqueue(Key interface{}) {
	if Q.length == 0 {
		Q.InitQueue(Key)
	} else {
		newNode := &Node{
			prev: Q.last,
			next: nil,
			key:  Key,
		}
		Q.last.next = newNode
		Q.last = newNode
		Q.length++
		Q.PrintQueue()
	}
}

func (Q *Queue) Dequeue() {
	if Q.length <= 1 {
		Q.ResetQueue()
	} else {
		Q.first = Q.first.next
		Q.length--
		Q.PrintQueue()
	}
}

func (Q *Queue) Peek() *Node {
	return Q.first
}

func (Q *Queue) ResetQueue() {
	Q.first = nil
	Q.last = nil
	Q.length = 0
}

func (Q *Queue) PrintQueue() {
	var s []interface{}
	currentNode := Q.first
	for currentNode != nil {
		s = append(s, currentNode.key)
		currentNode = currentNode.next
	}
	fmt.Println(s)
}
