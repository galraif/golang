package linked_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDoubleList_InitList(t *testing.T) {
	link := List{}
	value := 5
	expectedNode := &Node{
		prev: nil,
		next: nil,
		key:  value,
	}
	link.InitList(value)
	assert.Equal(t, 1, link.length)
	assert.Equal(t, expectedNode, link.head)
	assert.Equal(t, value, link.head.key)
	assert.Equal(t, expectedNode, link.tail)
	assert.Equal(t, value, link.tail.key)
}

func TestDoubleList_Append(t *testing.T) {
	link := List{}
	firstValue := 5
	expectedFirstNode := &Node{
		prev: nil,
		next: nil,
		key:  firstValue,
	}
	secondValue := 7
	expectedSecondNode := &Node{
		prev: expectedFirstNode,
		next: nil,
		key:  secondValue,
	}
	link.Append(firstValue)
	assert.Equal(t, 1, link.length)
	assert.Equal(t, expectedFirstNode, link.head)
	assert.Equal(t, firstValue, link.head.key)
	assert.Equal(t, expectedFirstNode, link.tail)
	assert.Equal(t, firstValue, link.tail.key)
	link.Append(secondValue)
	expectedFirstNode.next = expectedSecondNode
	assert.Equal(t, 2, link.length)
	assert.Equal(t, expectedFirstNode, link.head)
	assert.Equal(t, expectedSecondNode, link.tail)
}

func TestDoubleList_Prepend(t *testing.T) {
	link := List{}
	firstValue := 5
	expectedFirstNode := &Node{
		prev: nil,
		next: nil,
		key:  firstValue,
	}
	secondValue := 7
	expectedSecondNode := &Node{
		prev: nil,
		next: expectedFirstNode,
		key:  secondValue,
	}
	link.Append(firstValue)
	assert.Equal(t, 1, link.length)
	assert.Equal(t, expectedFirstNode, link.head)
	assert.Equal(t, firstValue, link.head.key)
	assert.Equal(t, expectedFirstNode, link.tail)
	assert.Equal(t, firstValue, link.tail.key)
	link.Prepend(secondValue)
	expectedFirstNode.prev = expectedSecondNode
	assert.Equal(t, 2, link.length)
	assert.Equal(t, expectedFirstNode, link.tail)
	assert.Equal(t, expectedSecondNode, link.head)
}

func TestDoubleList_TraverseToIndex(t *testing.T) {
	firstValue := 3
	secondValue := 5
	thirdValue := 7
	expectedFirstNode := &Node{
		prev: nil,
		next: nil,
		key:  firstValue,
	}
	expectedSecondNode := &Node{
		prev: expectedFirstNode,
		next: nil,
		key:  secondValue,
	}
	expectedThirdNode := &Node{
		prev: expectedSecondNode,
		next: nil,
		key:  thirdValue,
	}
	expectedFirstNode.next = expectedSecondNode
	expectedSecondNode.next = expectedThirdNode
	link := List{expectedFirstNode, expectedThirdNode, 3}
	assert.Equal(t, expectedFirstNode, link.TraverseToIndex(-1))
	assert.Equal(t, expectedFirstNode, link.TraverseToIndex(0))
	assert.Equal(t, expectedSecondNode, link.TraverseToIndex(1))
	assert.Equal(t, expectedThirdNode, link.TraverseToIndex(2))
	assert.Equal(t, expectedThirdNode, link.TraverseToIndex(7))
}

func TestDoubleList_Insert(t *testing.T) {
	firstValue := 3
	secondValue := 5
	insertedFirstValue := 7
	insertedLastValue := 9
	insertedRandomValue := 11
	firstNode := &Node{
		prev: nil,
		next: nil,
		key:  firstValue,
	}
	secondNode := &Node{
		prev: firstNode,
		next: nil,
		key:  secondValue,
	}
	firstNode.next = secondNode
	link := List{firstNode, secondNode, 2}
	link.Insert(-1, insertedFirstValue)
	assert.Equal(t, insertedFirstValue, link.head.key)
	link.Insert(34, insertedLastValue)
	assert.Equal(t, insertedLastValue, link.tail.key)
	link.Insert(1, insertedRandomValue)
	assert.Equal(t, link.head.next.key, insertedRandomValue)
}

func TestDoubleList_RemoveFirstNode(t *testing.T) {
	firstValue := 3
	secondValue := 5
	firstNode := &Node{
		prev: nil,
		next: nil,
		key:  firstValue,
	}
	secondNode := &Node{
		prev: firstNode,
		next: nil,
		key:  secondValue,
	}
	firstNode.next = secondNode
	link := List{firstNode, secondNode, 2}
	restLink := List{nil, nil, 0}
	assert.Equal(t, firstNode, link.head)
	assert.Equal(t, secondNode, link.tail)
	link.RemoveFirstNode()
	secondNode.prev = nil
	assert.Equal(t, secondNode, link.head)
	assert.Equal(t, secondNode, link.tail)
	link.RemoveFirstNode()
	assert.Equal(t, restLink, link)
	link.RemoveFirstNode()
	assert.Equal(t, restLink, link)
}

func TestDoubleList_RemoveLastNode(t *testing.T) {
	firstValue := 3
	secondValue := 5
	firstNode := &Node{
		prev: nil,
		next: nil,
		key:  firstValue,
	}
	secondNode := &Node{
		prev: firstNode,
		next: nil,
		key:  secondValue,
	}
	firstNode.next = secondNode
	link := List{firstNode, secondNode, 2}
	restLink := List{nil, nil, 0}
	assert.Equal(t, firstNode, link.head)
	assert.Equal(t, secondNode, link.tail)
	link.RemoveLastNode()
	assert.Equal(t, firstNode, link.head)
	assert.Equal(t, firstNode, link.tail)
	link.RemoveLastNode()
	assert.Equal(t, restLink, link)
	link.RemoveLastNode()
	assert.Equal(t, restLink, link)
}

func TestDoubleList_Remove(t *testing.T) {
	expectedFirstNode := &Node{
		prev: nil,
		next: nil,
		key:  3,
	}
	expectedSecondNode := &Node{
		prev: expectedFirstNode,
		next: nil,
		key:  5,
	}
	expectedThirdNode := &Node{
		prev: expectedSecondNode,
		next: nil,
		key:  7,
	}
	expectedFirstNode.next = expectedSecondNode
	expectedSecondNode.next = expectedThirdNode
	link := List{expectedFirstNode, expectedThirdNode, 3}
	restLink := List{nil, nil, 0}
	link.Remove(1)
	expectedFirstNode.next = expectedThirdNode
	expectedThirdNode.prev = expectedFirstNode
	assert.Equal(t, expectedFirstNode, link.head)
	assert.Equal(t, expectedThirdNode, link.tail)
	assert.Equal(t, 2, link.length)
	link.Remove(7)
	expectedFirstNode.next = nil
	assert.Equal(t, expectedFirstNode, link.head)
	assert.Equal(t, expectedFirstNode, link.tail)
	assert.Equal(t, 1, link.length)
	link.Remove(-1)
	assert.Equal(t, restLink, link)
	link.Remove(0)
	assert.Equal(t, restLink, link)
}

func TestDoubleList_Reverse(t *testing.T) {
	expectedFirstNode := &Node{
		prev: nil,
		next: nil,
		key:  3,
	}
	expectedSecondNode := &Node{
		prev: expectedFirstNode,
		next: nil,
		key:  5,
	}
	expectedThirdNode := &Node{
		prev: expectedSecondNode,
		next: nil,
		key:  7,
	}
	expectedFirstNode.next = expectedSecondNode
	expectedSecondNode.next = expectedThirdNode
	link := List{expectedFirstNode, expectedThirdNode, 3}
	link.Reverse()
	expectedThirdNode.prev = nil
	expectedThirdNode.next = expectedSecondNode
	expectedSecondNode.prev = expectedThirdNode
	expectedSecondNode.next = expectedFirstNode
	expectedFirstNode.prev = expectedSecondNode
	expectedFirstNode.next = nil
	assert.Equal(t, expectedThirdNode, link.head)
	assert.Equal(t, expectedFirstNode, link.tail)
	assert.Equal(t, expectedSecondNode, link.tail.prev)
	assert.Equal(t, expectedSecondNode, link.head.next)
}

func TestDoubleList_RestList(t *testing.T) {
	expectedFirstNode := &Node{
		prev: nil,
		next: nil,
		key:  3,
	}
	expectedSecondNode := &Node{
		prev: expectedFirstNode,
		next: nil,
		key:  5,
	}
	expectedThirdNode := &Node{
		prev: expectedSecondNode,
		next: nil,
		key:  7,
	}
	expectedFirstNode.next = expectedSecondNode
	expectedSecondNode.next = expectedThirdNode
	link := List{expectedFirstNode, expectedThirdNode, 3}
	restLink := List{nil, nil, 0}
	link.RestList()
	assert.Equal(t, restLink, link)
}
