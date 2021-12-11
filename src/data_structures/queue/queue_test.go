package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueue_InitQueue(t *testing.T) {
	queue := Queue{}
	value := 5
	expectedNode := &Node{
		prev: nil,
		next: nil,
		key:  value,
	}
	initQueue := Queue{expectedNode,expectedNode,1}
	queue.InitQueue(value)
	assert.Equal(t, initQueue, queue)
}

func TestQueue_Enqueue(t *testing.T) {
	queue := Queue{}
	firstValue := 5
	secondValue := 7
	thirdValue := 9
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
	queue.Enqueue(firstValue)
	assert.Equal(t, expectedFirstNode, queue.first)
	assert.Equal(t, expectedFirstNode, queue.last)
	assert.Equal(t, 1, queue.length)
	queue.Enqueue(secondValue)
	expectedFirstNode.next = expectedSecondNode
	assert.Equal(t, expectedSecondNode, queue.last)
	assert.Equal(t, expectedFirstNode, queue.first)
	assert.Equal(t, 2, queue.length)
	queue.Enqueue(thirdValue)
	expectedSecondNode.next = expectedThirdNode
	assert.Equal(t, expectedThirdNode, queue.last)
	assert.Equal(t, expectedFirstNode, queue.first)
	assert.Equal(t, 3, queue.length)
	assert.Equal(t, expectedSecondNode, queue.last.prev)
	assert.Equal(t, expectedSecondNode, queue.first.next)
}

func TestQueue_Dequeue(t *testing.T) {
	expectedFirstNode := &Node{
		prev: nil,
		next: nil,
		key:  5,
	}
	expectedSecondNode := &Node{
		prev: expectedFirstNode,
		next: nil,
		key:  7,
	}
	expectedThirdNode := &Node{
		prev: expectedSecondNode,
		next: nil,
		key:  9,
	}
	expectedFirstNode.next = expectedSecondNode
	expectedSecondNode.next = expectedThirdNode
	queue := Queue{expectedFirstNode, expectedThirdNode, 3}
	resetStack := Queue{nil, nil, 0}
	queue.Dequeue()
	expectedSecondNode.prev = nil
	assert.Equal(t, expectedSecondNode, queue.first)
	assert.Equal(t, expectedThirdNode, queue.last)
	assert.Equal(t, 2, queue.length)
	queue.Dequeue()
	expectedThirdNode.prev = nil
	assert.Equal(t, expectedThirdNode, queue.first)
	assert.Equal(t, expectedThirdNode, queue.last)
	assert.Equal(t, 1, queue.length)
	queue.Dequeue()
	assert.Equal(t, resetStack, queue)
	queue.Dequeue()
	assert.Equal(t, resetStack, queue)
}

func TestQueue_ResetQueue(t *testing.T) {
	expectedFirstNode := &Node{
		prev: nil,
		next: nil,
		key:  5,
	}
	expectedSecondNode := &Node{
		prev: expectedFirstNode,
		next: nil,
		key:  7,
	}
	expectedFirstNode.next = expectedSecondNode
	queue := Queue{expectedFirstNode, expectedSecondNode, 2}
	resetStack := Queue{nil, nil, 0}
	queue.ResetQueue()
	assert.Equal(t, resetStack, queue)
}

func TestQueue_Peek(t *testing.T) {
	expectedFirstNode := &Node{
		prev: nil,
		next: nil,
		key:  5,
	}
	expectedSecondNode := &Node{
		prev: expectedFirstNode,
		next: nil,
		key:  7,
	}
	expectedFirstNode.next = expectedSecondNode
	queue := Queue{expectedFirstNode, expectedSecondNode, 2}
	peak := queue.Peek()
	assert.Equal(t, peak, queue.first)
}
