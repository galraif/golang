package linked_list

import "fmt"

type Node struct {
	prev *Node
	next *Node
	key  interface{}
}

type List struct {
	head   *Node
	tail   *Node
	length int
}

func (L *List) InitList(Key interface{}) {
	newNode := &Node{
		prev: nil,
		next: nil,
		key:  Key,
	}
	L.head = newNode
	L.tail = newNode
	L.length = 1
}

func (L *List) Append(Key interface{}) {
	if L.length == 0 {
		L.InitList(Key)
	} else {
		newNode := &Node{
			prev: L.tail,
			next: nil,
			key:  Key,
		}
		L.tail.next = newNode
		L.tail = newNode
		L.length++
		L.PrintList()
	}
}

func (L *List) Prepend(Key interface{}) {
	newNode := &Node{
		prev: nil,
		next: L.head,
		key:  Key,
	}
	L.head.prev = newNode
	L.head = newNode
	L.length++
	L.PrintList()
}

func (L *List) TraverseToIndex(index int) *Node {
	switch {
	case index <=0:
		return L.head
	case index > L.length:
		return L.tail
	default:
		counter := 0
		currentNode := L.head
		for counter != index {
			currentNode = currentNode.next
			counter++
		}
		return currentNode
	}
}

func (L *List) Insert(index int, Key interface{}) {
	switch {
	case index >= L.length:
		L.Append(Key)
	case index <= 0:
		L.Prepend(Key)
	default:
		newNode := &Node{
			prev: nil,
			next: nil,
			key:  Key,
		}
		leader := L.TraverseToIndex(index - 1)
		follower := leader.next
		leader.next = newNode
		newNode.prev = leader
		newNode.next = follower
		follower.prev = newNode
		L.length++
		L.PrintList()
	}
}

func (L *List) RemoveFirstNode() {
	if L.length <= 1 {
		L.RestList()
	} else {
		newHead := L.head
		L.head = newHead.next
		newHead.prev = nil
		L.length--
	}
	L.PrintList()
}

func (L *List) RemoveLastNode() {
	if L.length <= 1 {
		L.RestList()
	} else {
		newTail := L.tail
		L.tail = newTail.prev
		L.tail.next = nil
		L.length--
	}
	L.PrintList()
}

func (L *List) Remove(index int) {
	switch {
	case index <= 0:
		L.RemoveFirstNode()
	case index+1 == L.length, index >= L.length:
		L.RemoveLastNode()
	default:
		leader := L.TraverseToIndex(index - 1)
		remove := leader.next
		follower := remove.next
		leader.next = follower
		follower.prev = remove.prev
		L.length--
		L.PrintList()
	}
}

func (L *List) Reverse() {
	currentNode := L.head
	var nextNode *Node
	L.head, L.tail = L.tail, L.head
	for currentNode != nil {
		nextNode = currentNode.next
		currentNode.next, currentNode.prev = currentNode.prev, currentNode.next
		currentNode = nextNode
	}
	L.PrintList()
}

func (L *List) PrintList() {
	var s []interface{}
	currentNode := L.head
	for currentNode != nil {
		s = append(s, currentNode.key)
		currentNode = currentNode.next
	}
	fmt.Println(s)
}

func (L *List) RestList() {
	L.head = nil
	L.tail = nil
	L.length = 0
}
