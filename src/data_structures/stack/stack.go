package stack

import "fmt"

type Node struct {
	prev *Node
	next *Node
	key  interface{}
}

type Stack struct {
	top    *Node
	bottom *Node
	length int
}

func (S *Stack) InitStack(Key interface{}) {
	newNode := &Node{
		prev: nil,
		next: nil,
		key:  Key,
	}
	S.bottom = newNode
	S.top = newNode
	S.length = 1
	S.PrintStack()
}

func (S *Stack) Push(Key interface{}) {
	if S.length == 0 {
		S.InitStack(Key)
	} else {
		newNode := &Node{
			prev: S.top,
			next: nil,
			key:  Key,
		}
		S.top.next = newNode
		S.top = newNode
		S.length++
		S.PrintStack()
	}
}

func (S *Stack) Pop() {
	if S.length <= 1 {
		S.ResetStack()
	} else {
		S.top = S.top.prev
		S.top.next = nil
		S.length--
		S.PrintStack()
	}
}

func (S *Stack) Peek() *Node {
	return S.top
}

func (S *Stack) ResetStack() {
	S.top = nil
	S.bottom = nil
	S.length = 0
}

func (S *Stack) PrintStack() {
	var s []interface{}
	currentNode := S.bottom
	for currentNode != nil {
		s = append(s, currentNode.key)
		currentNode = currentNode.next
	}
	fmt.Println(s)
}
