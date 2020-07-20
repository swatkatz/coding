package main

import (
	"fmt"
	"github.com/golang-collections/collections/queue"
)

type EdgeNode struct {
	nodeID int
	weight int
	next *EdgeNode // linked list of all the nodes
}

type Graph struct {
	edges []*EdgeNode // adjacency info
	degree []int // out-degree of each vertex
	nvertices int // number of vertices
	nedges int // number of edges (for an undirected graph, it is counted only once)
	directed bool
}

func NewGraph(directed bool, vertices int) *Graph {
	return &Graph{
		edges:     make([]*EdgeNode, vertices),
		degree:    make([]int, vertices),
		nvertices: vertices,
		directed:  directed,
	}
}

func (g *Graph) insertEdge(x int, y int, directed bool) {
	yEdgeNode := &EdgeNode{
		nodeID: y,
	}
	yEdgeNode.next = g.edges[x] // attach the new edgeNode at the top of the list
	g.edges[x] = yEdgeNode // make yEdgeNode the head
	g.nedges += 1 // increment the number of edges by 1
	if !directed {
		g.insertEdge(y, x, true) // if undirected insert another edge
	} else {
		g.nedges += 1
	}
}

func (g *Graph) printGraph() {
	for i := 0; i < g.nvertices; i++ {
		fmt.Printf("Vertex: %v => ", i)
		edge := g.edges[i]
		for edge != nil {
			fmt.Printf(" %v ", edge.nodeID)
			edge = edge.next
		}
		fmt.Printf("\n")
	}
}

type GraphSearch struct {
	discovered []bool
	processed []bool
}

func (g *Graph) initializeSearch() *GraphSearch {
	// this is largely useless
	processed := make([]bool, g.nvertices)  // vertices that have been completely processed
	discovered := make([]bool, g.nvertices) // vertices that have been discovered
	return &GraphSearch{
		discovered: discovered,
		processed:  processed, // becomes terribly important in DFS
	}
}

func (g *Graph) bfs(start int, gs *GraphSearch) []int {
	// link from discovered v to discoverer u i.e. parents[v] = u (per bfs)
	parents := make([]int, g.nvertices)
	// Initialize parents to -1
	for i := range parents {
		parents[i] = -1
	}
	// use a queue
	q := queue.New()
	q.Enqueue(start)
	gs.discovered[start] = true
	for q.Len() > 0 {
		u := q.Dequeue().(int)
		edge := g.edges[u]
		for edge != nil {
			v := edge.nodeID
			// process edge
			if g.directed || !gs.processed[v] {
				fmt.Printf("process edge %v -> %v \n", u, v)
			}
			if !gs.discovered[v] {
				// add to the queue
				q.Enqueue(v)
				// mark as discovered
				gs.discovered[v] = true
				// add to parents
				parents[v] = u
			}
			edge = edge.next
		}
		// all nodes for u have been discovered
		gs.processed[u] = true
		// process vertex u
		fmt.Printf("vertex %v has been processed \n", u)
	}
	return parents
}

func (g *Graph) dfs(u int, gs *GraphSearch, parents []int) []int {
	gs.discovered[u] = true
	// potentially could pre_process node here
	edge := g.edges[u]
	for edge != nil {
		v := edge.nodeID
		// if this node has not been discovered
		if !gs.discovered[v] {
			parents[v] = u
			// process edge
			fmt.Printf("process edge %v -> %v \n", u, v)
			g.dfs(v, gs, parents)
		} else if !gs.processed[v] || g.directed {
			// The vertex may have been discovered, but not all of its edges have been processed.
			// This can happen if a vertex was reached from somewhere, but it hasn't itself been
			// processed yet. This is also useful in finding out cycles,
			// as the parent is not assigned here, rather it is assigned only when discovery happens

			// process edge
			fmt.Printf("process edge %v -> %v \n", u, v)
		}
		edge = edge.next
	}
	gs.processed[u] = true
	// process vertex u
	fmt.Printf("vertex %v has been processed \n", u)
	return parents
}

// reverse a stack, go to the end of the stack, print the value, then pop the remaining
func findPaths(start int, end int, parents []int) {
	// reached the top of the tree (root node)
	if start == end || end == -1 {
		fmt.Printf("%v -> ", start)
	} else {
		findPaths(start, parents[end], parents)
		fmt.Printf("%v -> ", end)
	}
}

func main() {
	g := NewGraph(false, 5)
	g.insertEdge(0, 1, g.directed)
	g.insertEdge(0, 2, g.directed)
	g.insertEdge(1, 2, g.directed)
	g.insertEdge(1, 3, g.directed)
	g.insertEdge(2, 3, g.directed)
	// this is a lone vertex
	g.insertEdge(4, 4, g.directed)
	g.printGraph()

	// BFS
	/*search := g.initializeSearch()
	parents := g.bfs(0, search)
	// this has to be from the root node (SHORTEST PATH)
	findPaths(0, 3, parents)
	fmt.Println()
	findPaths(0, 1, parents)

	// connected components
	search = g.initializeSearch()
	c := 0
	for i := 0; i < g.nvertices; i++ {
		if !search.discovered[i] {
			// found root for 1 connected component
			c += 1
			// perform bfs and connect everything that can be connected
			g.bfs(i, search)
		}
	}
	fmt.Printf("# of connected components %v \n", c)*/

	// DFS
	search := g.initializeSearch()
	parents := make([]int, g.nvertices)
	for i := 0; i < g.nvertices; i++ {
		parents[i] = -1
	}
	parents = g.dfs(0, search, parents)
	fmt.Printf("parents: %v \n", parents)
}

