package DFSAndBFS

type Node struct {
	Val       int
	Neighbors []*Node
}

var visited = make(map[*Node]*Node)

func cloneGraph(node *Node) *Node {
	if node == nil {
		return node
	}

	if clonedNode, exists := visited[node]; exists {
		return clonedNode
	}

	clonedNode := &Node{
		Val:       node.Val,
		Neighbors: []*Node{},
	}
	visited[node] = clonedNode

	for _, node := range node.Neighbors {
		clonedNode.Neighbors = append(clonedNode.Neighbors, cloneGraph(node))
	}
	return clonedNode
}
