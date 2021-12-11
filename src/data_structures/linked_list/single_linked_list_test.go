package linked_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList_InitList(t *testing.T) {
	link := SingleList{}
	value := 5
	expectedNode := &SingleNode{
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

func TestList_Append(t *testing.T) {
	link := SingleList{}
	firstValue := 5
	expectedFirstNode := &SingleNode{
		next: nil,
		key:  firstValue,
	}
	secondValue := 7
	expectedSecondNode := &SingleNode{
		next: nil,
		key:  secondValue,
	}
	link.Append(firstValue)
	assert.Equal(t, 1, link.length)
	assert.Equal(t, expectedFirstNode, link.head)
	assert.Equal(t, expectedFirstNode, link.tail)
	link.Append(secondValue)
	expectedFirstNode.next = expectedSecondNode
	assert.Equal(t, 2, link.length)
	assert.Equal(t, expectedFirstNode, link.head)
	assert.Equal(t, expectedSecondNode, link.tail)
}

func TestList_Prepend(t *testing.T) {
	link := SingleList{}
	firstValue := 5
	expectedFirstNode := &SingleNode{
		next: nil,
		key:  firstValue,
	}
	secondValue := 7
	expectedSecondNode := &SingleNode{
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
	assert.Equal(t, 2, link.length)
	assert.Equal(t, expectedFirstNode, link.tail)
	assert.Equal(t, expectedSecondNode, link.head)
}

func TestList_TraverseToIndex(t *testing.T) {
	expectedThirdNode := &SingleNode{
		next: nil,
		key:  7,
	}
	expectedSecondNode := &SingleNode{
		next: expectedThirdNode,
		key:  5,
	}
	expectedFirstNode := &SingleNode{
		next: expectedSecondNode,
		key:  3,
	}
	link := SingleList{expectedFirstNode, expectedThirdNode, 3}
	assert.Equal(t, expectedFirstNode, link.TraverseToIndex(-1))
	assert.Equal(t, expectedFirstNode, link.TraverseToIndex(0))
	assert.Equal(t, expectedSecondNode, link.TraverseToIndex(1))
	assert.Equal(t, expectedThirdNode, link.TraverseToIndex(2))
	assert.Equal(t, expectedThirdNode, link.TraverseToIndex(7))
}

func TestList_Insert(t *testing.T) {
	firstValue := 3
	secondValue := 5
	insertedFirstValue := 7
	insertedLastValue := 9
	insertedRandomValue := 11
	secondNode := &SingleNode{
		next: nil,
		key:  secondValue,
	}
	firstNode := &SingleNode{
		next: secondNode,
		key:  firstValue,
	}
	link := SingleList{firstNode, secondNode, 2}
	link.Insert(-1, insertedFirstValue)
	assert.Equal(t, insertedFirstValue, link.head.key)
	link.Insert(34, insertedLastValue)
	assert.Equal(t, insertedLastValue, link.tail.key)
	link.Insert(1, insertedRandomValue)
	assert.Equal(t, link.head.next.key, insertedRandomValue)
}

func TestList_RemoveFirstNode(t *testing.T) {
	firstValue := 3
	secondValue := 5
	secondNode := &SingleNode{
		next: nil,
		key:  secondValue,
	}
	firstNode := &SingleNode{
		next: secondNode,
		key:  firstValue,
	}
	link := SingleList{firstNode, secondNode, 2}
	restLink := SingleList{nil, nil, 0}
	assert.Equal(t, firstNode, link.head)
	assert.Equal(t, secondNode, link.tail)
	link.RemoveFirstNode()
	assert.Equal(t, secondNode, link.head)
	assert.Equal(t, secondNode, link.tail)
	link.RemoveFirstNode()
	assert.Equal(t, restLink, link)
	link.RemoveFirstNode()
	assert.Equal(t, restLink, link)
}

func TestList_RemoveLastNode(t *testing.T) {
	expectedThirdNode := &SingleNode{
		next: nil,
		key:  7,
	}
	expectedSecondNode := &SingleNode{
		next: expectedThirdNode,
		key:  5,
	}
	expectedFirstNode := &SingleNode{
		next: expectedSecondNode,
		key:  3,
	}
	link := SingleList{expectedFirstNode, expectedThirdNode, 3}
	restLink := SingleList{nil, nil, 0}
	link.RemoveLastNode()
	assert.Equal(t, 2, link.length)
	assert.Equal(t, expectedFirstNode, link.head)
	assert.Equal(t, expectedSecondNode, link.tail)
	link.RemoveLastNode()
	assert.Equal(t, 1,link.length)
	assert.Equal(t, expectedFirstNode, link.head)
	assert.Equal(t, expectedFirstNode, link.tail)
	link.RemoveLastNode()
	assert.Equal(t, restLink, link)
	link.RemoveLastNode()
	assert.Equal(t, restLink, link)
}

func TestList_Remove(t *testing.T) {
	expectedThirdNode := &SingleNode{
		next: nil,
		key:  7,
	}
	expectedSecondNode := &SingleNode{
		next: expectedThirdNode,
		key:  5,
	}
	expectedFirstNode := &SingleNode{
		next: expectedSecondNode,
		key:  3,
	}
	link := SingleList{expectedFirstNode, expectedThirdNode, 3}
	restLink := SingleList{nil, nil, 0}
	link.Remove(1)
	expectedFirstNode.next = expectedThirdNode
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

func TestList_Reverse(t *testing.T) {
	expectedThirdNode := &SingleNode{
		next: nil,
		key:  7,
	}
	expectedSecondNode := &SingleNode{
		next: expectedThirdNode,
		key:  5,
	}
	expectedFirstNode := &SingleNode{
		next: expectedSecondNode,
		key:  3,
	}
	link := SingleList{expectedFirstNode, expectedThirdNode, 3}
	reverseLink := SingleList{expectedThirdNode, expectedFirstNode, 3}
	link.Reverse()
	expectedThirdNode.next = expectedSecondNode
	expectedSecondNode.next = expectedFirstNode
	expectedFirstNode.next = nil
	assert.Equal(t, reverseLink, link)
	assert.Equal(t, expectedSecondNode, link.head.next)
	assert.Equal(t, expectedFirstNode, link.head.next.next)
}

func TestList_RestList(t *testing.T) {
	expectedThirdNode := &SingleNode{
		next: nil,
		key:  7,
	}
	expectedSecondNode := &SingleNode{
		next: expectedThirdNode,
		key:  5,
	}
	expectedFirstNode := &SingleNode{
		next: expectedSecondNode,
		key:  3,
	}
	link := SingleList{expectedFirstNode, expectedThirdNode, 3}
	restLink := SingleList{nil, nil, 0}
	link.RestList()
	assert.Equal(t, restLink, link)
}

