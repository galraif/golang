package binary_search_tree

import "fmt"

type TreeNode struct {
	left  *TreeNode
	right *TreeNode
	value int
}

type Tree struct {
	Root *TreeNode
}

func (T *Tree) PrintInOrderTree(node *TreeNode) {
	if node != nil {
		T.PrintInOrderTree(node.left)
		fmt.Println(node.value)
		T.PrintInOrderTree(node.right)
	}
}

func (T *Tree) Insert(value int) {
	newNode := &TreeNode{
		left:  nil,
		right: nil,
		value: value,
	}
	if T.Root == nil {
		T.Root = newNode
	} else {
		currentNode := T.Root
		for {
			if value < currentNode.value {
				//Left
				if currentNode.left == nil {
					currentNode.left = newNode
					break
				}
				currentNode = currentNode.left
			} else {
				//Right
				if currentNode.right == nil {
					currentNode.right = newNode
					break
				}
				currentNode = currentNode.right
			}
		}
	}
}

func (T *Tree) Lookup(value int) *TreeNode {
	if T.Root == nil {
		return &TreeNode{}
	}
	currentNode := T.Root
	for currentNode != nil {
		if value < currentNode.value {
			currentNode = currentNode.left
		} else if value > currentNode.value {
			currentNode = currentNode.right
		} else if value == currentNode.value {
			return currentNode
		}
	}
	return &TreeNode{}
}

func (T *Tree) IsExist(value int) bool {
	checkNode := T.Lookup(value)
	return checkNode.value == value
}

func (T *Tree) IsValueLeaf(value int) bool {
	node := T.Lookup(value)
	return node.left == nil && node.right == nil
}

func (T *Tree) IsTreeNodeLeaf(node *TreeNode) bool {
	return node.left == nil && node.right == nil
}

func (T *Tree) MinValueInTree() int {
	if T.Root == nil {
		return 0
	}
	currentValue := T.Root.value
	currentNode := T.Root
	for currentNode != nil {
		if currentNode.left == nil {
			currentValue = currentNode.value
			break
		}
		if currentNode.left.value < currentNode.value {
			currentNode = currentNode.left
		}
	}
	return currentValue
}

func (T *Tree) MaxValueInTree() int {
	if T.Root == nil {
		return 0
	}
	currentValue := T.Root.value
	currentNode := T.Root
	for currentNode != nil {
		if currentNode.right == nil {
			currentValue = currentNode.value
			break
		}
		if currentNode.right.value > currentNode.value {
			currentNode = currentNode.right
		}
	}
	return currentValue
}

func (T *Tree) MinValueNode(node *TreeNode) *TreeNode {
	currentNode := node
	for currentNode.left != nil {
		currentNode = currentNode.left
	}
	return currentNode
}

func (T *Tree) MaxValueNode(node *TreeNode) *TreeNode {
	currentNode := node
	for currentNode.right != nil {
		currentNode = currentNode.right
	}
	return currentNode
}

func (T *Tree) Remove(value int) {
	if T.Root == nil {
		return
	}
	parentNode := &TreeNode{}
	currentNode := T.Root
	for currentNode != nil {
		if value < currentNode.value {
			parentNode = currentNode
			currentNode = currentNode.left
		} else if value > currentNode.value {
			parentNode = currentNode
			currentNode = currentNode.right
			//value == currentNode.value
		} else {
			//Node without children
			if T.IsTreeNodeLeaf(currentNode) {
				if parentNode.value == 0 {
					T.Root = nil
					break
				} else {
					//if parent > current value, set parent left nil
					if currentNode.value < parentNode.value {
						parentNode.left = nil
						break
						//if parent < current value, set parent right nil
					} else if currentNode.value > parentNode.value {
						parentNode.right = nil
						break
					}
				}
				//Node with one child (without right)
			} else if currentNode.right == nil {
				if currentNode.value < parentNode.value {
					parentNode.left = currentNode.left
					break
				} else if currentNode.value > parentNode.value {
					parentNode.right = currentNode.left
					break
				}
				// Node with one child (without left)
			} else if currentNode.left == nil {
				if currentNode.value < parentNode.value {
					parentNode.left = currentNode.right
					break
				} else if currentNode.value > parentNode.value {
					parentNode.right = currentNode.right
					break
				}
				//Node with 2 children
				//1. Find node with min value from right child
				//2. Remove the node
				//3. Set current value from the deleted min value node
			} else if currentNode.left != nil && currentNode.right != nil {
				minValueNode := T.MinValueNode(currentNode.right)
				T.Remove(minValueNode.value)
				currentNode.value = minValueNode.value
				break
			}
		}
	}
}
