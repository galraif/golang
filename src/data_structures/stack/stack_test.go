package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack_InitStack(t *testing.T) {
	stack := Stack{}
	value := 5
	expectedNode := &Node{
		prev: nil,
		next: nil,
		key:  value,
	}
	stack.InitStack(value)
	assert.Equal(t, expectedNode, stack.bottom)
	assert.Equal(t, expectedNode, stack.top)
	assert.Equal(t, 1, stack.length)
}

func TestStack_Push(t *testing.T) {
	stack := Stack{}
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
	stack.Push(firstValue)
	assert.Equal(t, expectedFirstNode, stack.top)
	assert.Equal(t, expectedFirstNode, stack.bottom)
	assert.Equal(t, 1, stack.length)
	stack.Push(secondValue)
	expectedFirstNode.next = expectedSecondNode
	assert.Equal(t, expectedSecondNode, stack.top)
	assert.Equal(t, expectedFirstNode, stack.bottom)
	assert.Equal(t, 2, stack.length)
	stack.Push(thirdValue)
	expectedSecondNode.next = expectedThirdNode
	assert.Equal(t, expectedThirdNode, stack.top)
	assert.Equal(t, expectedFirstNode, stack.bottom)
	assert.Equal(t, 3, stack.length)
	assert.Equal(t, expectedSecondNode, stack.top.prev)
	assert.Equal(t, expectedSecondNode, stack.bottom.next)
}

func TestStack_Pop(t *testing.T) {
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
	stack := Stack{expectedThirdNode, expectedFirstNode, 3}
	resetStack := Stack{nil, nil, 0}
	stack.Pop()
	expectedSecondNode.next = nil
	assert.Equal(t, expectedSecondNode, stack.top)
	assert.Equal(t, expectedFirstNode, stack.bottom)
	assert.Equal(t, 2, stack.length)
	stack.Pop()
	expectedFirstNode.next = nil
	assert.Equal(t, expectedFirstNode, stack.top)
	assert.Equal(t, expectedFirstNode, stack.bottom)
	assert.Equal(t, 1, stack.length)
	stack.Pop()
	assert.Equal(t, resetStack, stack)
	stack.Pop()
	assert.Equal(t, resetStack, stack)
}

func TestStack_Peek(t *testing.T) {
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
	stack := Stack{expectedSecondNode, expectedFirstNode, 2}
	peak := stack.Peek()
	assert.Equal(t, peak, stack.top)
}

func TestStack_ResetStack(t *testing.T) {
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
	stack := Stack{expectedSecondNode, expectedFirstNode, 2}
	resetStack := Stack{nil, nil, 0}
	stack.ResetStack()
	assert.Equal(t, resetStack, stack)
}
