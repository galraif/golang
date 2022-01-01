package graph

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func getGraphWithNodes(firstValue, secondValue, thirdValue int) *Graph {
	g := Graph{}
	g.AddNode(firstValue)
	g.AddNode(secondValue)
	g.AddNode(thirdValue)
	return &g
}

func TestGraph_AddNode(t *testing.T) {
	firstValue := 12
	secondValue := 14
	thirdValue := 17
	g := getGraphWithNodes(firstValue, secondValue, thirdValue)
	assert.Equal(t, firstValue, g.Nodes[0].Value)
	assert.Equal(t, secondValue, g.Nodes[1].Value)
	assert.Equal(t, thirdValue, g.Nodes[2].Value)
	assert.Equal(t, 3, len(g.Nodes))
	assert.Equal(t, 0, len(g.Edges))
}

func TestGraph_AddEdge(t *testing.T) {
	firstValue := 12
	secondValue := 14
	thirdValue := 17
	firstNode := &Node{
		Value: firstValue,
	}
	secondNode := &Node{
		Value: secondValue,
	}
	thirdNode := &Node{
		Value: thirdValue,
	}
	g := getGraphWithNodes(firstValue, secondValue, thirdValue)
	g.AddEdge(firstNode, secondNode)
	g.AddEdge(firstNode, thirdNode)
	g.AddEdge(secondNode, thirdNode)
	//Expected graph
	//     12
	//   /    \
	//  4   -  20
	assert.Equal(t, secondNode, g.Edges[*firstNode][0])
	assert.Equal(t, firstNode, g.Edges[*secondNode][0])
	assert.Equal(t, thirdNode, g.Edges[*firstNode][1])
	assert.Equal(t, firstNode, g.Edges[*thirdNode][0])
	assert.Equal(t, thirdNode, g.Edges[*secondNode][1])
	assert.Equal(t, secondNode, g.Edges[*thirdNode][1])
}
