package binary_search_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTree_Insert(t *testing.T) {
	tree := Tree{}
	tree.Insert(9)
	tree.Insert(4)
	tree.Insert(6)
	tree.Insert(20)
	tree.Insert(170)
	tree.Insert(15)
	tree.Insert(1)
	//Expected tree
	//     9
	//  4     20
	//1  6  15  170
	assert.Equal(t, 9, tree.Root.value)
	assert.Equal(t, 4, tree.Root.left.value)
	assert.Equal(t, 20, tree.Root.right.value)
	assert.Equal(t, 1, tree.Root.left.left.value)
	assert.Equal(t, 6, tree.Root.left.right.value)
	assert.Equal(t, 15, tree.Root.right.left.value)
	assert.Equal(t, 170, tree.Root.right.right.value)
}

func TestTree_Lookup(t *testing.T) {
	rightNode := &TreeNode{
		value: 20,
		left:  nil,
		right: nil,
	}
	leftNode := &TreeNode{
		value: 4,
		left:  nil,
		right: nil,
	}
	rootNode := &TreeNode{
		value: 9,
		left:  leftNode,
		right: rightNode,
	}
	tree := Tree{}
	assert.Equal(t, &TreeNode{}, tree.Lookup(12))
	tree.Root = rootNode
	//Expected tree
	//     9
	//  4     20
	assert.Equal(t, rootNode, tree.Lookup(9))
	assert.Equal(t, leftNode, tree.Lookup(4))
	assert.Equal(t, rightNode, tree.Lookup(20))
	assert.Equal(t, &TreeNode{}, tree.Lookup(13))
}

func TestTree_IsExist(t *testing.T) {
	rightNode := &TreeNode{
		value: 20,
		left:  nil,
		right: nil,
	}
	leftNode := &TreeNode{
		value: 4,
		left:  nil,
		right: nil,
	}
	rootNode := &TreeNode{
		value: 9,
		left:  leftNode,
		right: rightNode,
	}
	tree := Tree{}
	assert.Equal(t, false, tree.IsExist(12))
	tree.Root = rootNode
	//Expected tree
	//     9
	//  4     20
	assert.Equal(t, false, tree.IsExist(8))
	assert.Equal(t, true, tree.IsExist(9))
	assert.Equal(t, false, tree.IsExist(10))
	assert.Equal(t, true, tree.IsExist(4))
	assert.Equal(t, false, tree.IsExist(19))
	assert.Equal(t, true, tree.IsExist(20))
	assert.Equal(t, false, tree.IsExist(-1))
}

func TestTree_IsLeaf(t *testing.T) {
	tree := Tree{}
	tree.Insert(9)
	tree.Insert(4)
	tree.Insert(6)
	tree.Insert(20)
	tree.Insert(170)
	tree.Insert(15)
	tree.Insert(1)
	//Expected tree
	//     9
	//  4     20
	//1  6  15  170
	assert.Equal(t, false, tree.IsValueLeaf(9))
	assert.Equal(t, false, tree.IsValueLeaf(4))
	assert.Equal(t, false, tree.IsValueLeaf(20))
	assert.Equal(t, true, tree.IsValueLeaf(1))
	assert.Equal(t, true, tree.IsValueLeaf(6))
	assert.Equal(t, true, tree.IsValueLeaf(15))
	assert.Equal(t, true, tree.IsValueLeaf(170))
}

func TestTree_IsTreeNodeLeaf(t *testing.T) {
	rightNode := &TreeNode{
		value: 20,
		left:  nil,
		right: nil,
	}
	leftNode := &TreeNode{
		value: 4,
		left:  nil,
		right: nil,
	}
	rootNode := &TreeNode{
		value: 9,
		left:  leftNode,
		right: rightNode,
	}
	tree := Tree{rootNode}
	//Expected tree
	//     9
	//  4     20
	assert.Equal(t, false, tree.IsTreeNodeLeaf(rootNode))
	assert.Equal(t, true, tree.IsTreeNodeLeaf(leftNode))
	assert.Equal(t, true, tree.IsTreeNodeLeaf(rightNode))
}

func TestTree_MinValueInTree(t *testing.T) {
	tree := Tree{}
	tree.Insert(9)
	tree.Insert(4)
	tree.Insert(6)
	tree.Insert(20)
	tree.Insert(170)
	tree.Insert(15)
	tree.Insert(1)
	//Expected tree
	//     9
	//  4     20
	//1  6  15  170
	assert.Equal(t, 1, tree.MinValueInTree())
}

func TestTree_MaxValueInTree(t *testing.T) {
	tree := Tree{}
	tree.Insert(9)
	tree.Insert(4)
	tree.Insert(6)
	tree.Insert(20)
	tree.Insert(170)
	tree.Insert(15)
	tree.Insert(1)
	//Expected tree
	//     9
	//  4     20
	//1  6  15  170
	assert.Equal(t, 170, tree.MaxValueInTree())
}

func TestTree_RemoveNodeWithoutChildren(t *testing.T) {
	var emptyNode *TreeNode
	tree := Tree{}
	rootValue := 9
	leftValue := 4
	rightValue := 20
	tree.Insert(rootValue)
	tree.Insert(leftValue)
	tree.Insert(rightValue)
	//Expected tree
	//     9
	//  4     20
	assert.Equal(t, rootValue, tree.Root.value)
	assert.Equal(t, leftValue, tree.Root.left.value)
	assert.Equal(t, rightValue, tree.Root.right.value)
	tree.Remove(4)
	//Expected tree
	//     9
	//        20
	assert.Equal(t, emptyNode, tree.Root.left)
	tree.Remove(20)
	//Expected tree
	//     9
	assert.Equal(t, emptyNode, tree.Root.right)
	tree.Remove(9)
	assert.Equal(t, emptyNode, tree.Root)
}

func TestTree_RemoveNodeWithOneLeftChild(t *testing.T) {
	originLeftValue := 4
	originRightValue := 20
	expectedLeftValue := 1
	expectedRightValue := 15
	tree := Tree{}
	tree.Insert(9)
	tree.Insert(originLeftValue)
	tree.Insert(originRightValue)
	tree.Insert(expectedRightValue)
	tree.Insert(expectedLeftValue)
	//Expected tree
	//     9
	//  4     20
	//1     15
	tree.Remove(originLeftValue)
	//Expected tree
	//     9
	//  1     20
	//      15
	assert.Equal(t, expectedLeftValue, tree.Root.left.value)
	tree.Remove(originRightValue)
	//Expected tree
	//     9
	//  1     15
	assert.Equal(t, expectedRightValue, tree.Root.right.value)
}

func TestTree_RemoveNodeWithOneRightChild(t *testing.T) {
	tree := Tree{}
	originLeftValue := 4
	originRightValue := 20
	expectedLeftValue := 6
	expectedRightValue := 170
	tree.Insert(9)
	tree.Insert(originLeftValue)
	tree.Insert(expectedLeftValue)
	tree.Insert(originRightValue)
	tree.Insert(expectedRightValue)
	//Expected tree
	//     9
	//  4     20
	//   6      170
	tree.Remove(originLeftValue)
	//Expected tree
	//     9
	//  6     20
	//          170
	assert.Equal(t, expectedLeftValue, tree.Root.left.value)
	tree.Remove(originRightValue)
	//Expected tree
	//     9
	//  6     170
	assert.Equal(t, expectedRightValue, tree.Root.right.value)
}

func TestTree_RemoveWithTwoChildren(t *testing.T) {
	tree := Tree{}
	originRoot := 9
	expectedRoot := 14
	tree.Insert(originRoot)
	tree.Insert(4)
	tree.Insert(6)
	tree.Insert(20)
	tree.Insert(170)
	tree.Insert(15)
	tree.Insert(expectedRoot)
	tree.Insert(1)
	//Expected tree
	//         9
	//     4        20
	//   1   6    15  170
	//           14
	tree.Remove(originRoot)
	//Expected tree
	//         14
	//     4        20
	//   1   6    15  170
	assert.Equal(t, expectedRoot, tree.Root.value)
}
