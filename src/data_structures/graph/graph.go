package graph

type Node struct {
	Value int
}

type Graph struct {
	Nodes []*Node
	Edges map[Node][]*Node
}

func (G *Graph) AddNode(value int) {
	newNode := &Node{
		Value: value,
	}
	G.Nodes = append(G.Nodes, newNode)
}

func (G *Graph) AddEdge(firstVertex, secondVertex *Node) {
	if G.Edges == nil {
		G.Edges = make(map[Node][]*Node)
	}
	G.Edges[*firstVertex] = append(G.Edges[*firstVertex], secondVertex)
	G.Edges[*secondVertex] = append(G.Edges[*secondVertex], firstVertex)
}
