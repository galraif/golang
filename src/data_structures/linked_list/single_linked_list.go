package linked_list

import "fmt"

type SingleNode struct {
	next *SingleNode
	key  interface{}
}

type SingleList struct {
	head   *SingleNode
	tail   *SingleNode
	length int
}

func (L *SingleList) InitList(Key interface{}) {
	newNode := &SingleNode{
		next: nil,
		key:  Key,
	}
	L.head = newNode
	L.tail = newNode
	L.length = 1
}

func (L *SingleList) Append(Key interface{}) {
	if L.length == 0 {
		L.InitList(Key)
	} else {
		newNode := &SingleNode{
			next: nil,
			key:  Key,
		}
		L.tail.next = newNode
		L.tail = newNode
		L.length++
		L.PrintList()
	}
}

func (L *SingleList) Prepend(Key interface{}) {
	newNode := &SingleNode{
		next: L.head,
		key:  Key,
	}
	L.head = newNode
	L.length++
	L.PrintList()
}

func (L *SingleList) TraverseToIndex(index int) *SingleNode {
	switch {
	case index <= 0:
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

func (L *SingleList) Insert(index int, Key interface{}) {
	switch {
	case index >= L.length:
		L.Append(Key)
	case index <= 0:
		L.Prepend(Key)
	default:
		newNode := &SingleNode{
			next: nil,
			key:  Key,
		}
		leader := L.TraverseToIndex(index - 1)
		follower := leader.next
		leader.next = newNode
		newNode.next = follower
		L.length++
		L.PrintList()
	}
}

func (L *SingleList) RemoveFirstNode() {
	if L.length <= 1 {
		L.RestList()
	} else {
		newHead := L.head
		L.head = newHead.next
		L.length--
	}
	L.PrintList()
}

func (L *SingleList) RemoveLastNode() {
	if L.length <= 1 {
		L.RestList()
	} else {
		L.tail = L.TraverseToIndex(L.length - 2)
		L.tail.next = nil
		L.length--
	}
	L.PrintList()
}

func (L *SingleList) Remove(index int) {
	switch {
	case index <= 0:
		L.RemoveFirstNode()
	case index+1 == L.length, index >= L.length:
		L.RemoveLastNode()
	default:
		leader := L.TraverseToIndex(index - 1)
		removeNode := leader.next
		leader.next = removeNode.next
		L.length--
		L.PrintList()
	}
}

func (L *SingleList) Reverse() {
	currentNode := L.head
	var nextNode *SingleNode
	var prevNode *SingleNode
	L.tail = L.head
	for currentNode != nil {
		nextNode, currentNode.next = currentNode.next, prevNode
		prevNode, currentNode = currentNode, nextNode
	}
	L.head = prevNode
	L.PrintList()
}

func (L *SingleList) RestList() {
	L.head = nil
	L.tail = nil
	L.length = 0
}

func (L *SingleList) PrintList() {
	var s []interface{}
	currentNode := L.head
	for currentNode != nil {
		s = append(s, currentNode.key)
		currentNode = currentNode.next
	}
	fmt.Println(s)
}
