package main

import "fmt"

// Graph struct
type Graph struct {
	vertexes []*Vertex
}

// Vertex struct
type Vertex struct {
	key      int
	adjacent []*Vertex
}

// add Vertex to graph
func (g *Graph) addVertex(key int) {
	if contains(g.vertexes, key) {
		err := fmt.Errorf("key is in graph verteses: %v", key)
		fmt.Println(err.Error())
	} else {
		g.vertexes = append(g.vertexes, &Vertex{key: key})
	}
}

func (g *Graph) addEdge(from, to int) {
	// getVertex from graph
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	// check error
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("content can not be nil %v -> %v", from, to)
		fmt.Println(err.Error())
	} else if contains(fromVertex.adjacent, to) {
		err := fmt.Errorf("already exists %v -> %v", from, to)
		fmt.Println(err.Error())
	} else {
		// add edge to vertexes
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	}
}

// get Vertex
func (g *Graph) getVertex(key int) *Vertex {
	for i, v := range g.vertexes {
		if v.key == key {
			return g.vertexes[i]
		}
	}
	return nil
}

// contains
func contains(s []*Vertex, key int) bool {
	for _, v := range s {
		if key == v.key {
			return true
		}
	}
	return false
}

// Print vertexes
func (g *Graph) PrintVertex() {
	for _, v := range g.vertexes {
		fmt.Printf("\nVertex %d:", v.key)
		for _, v := range v.adjacent {
			fmt.Printf(" %v", v.key)
		}
	}
	fmt.Println()
}

func main() {
	fmt.Println("Graph")

	g := &Graph{}
	for i := 1; i < 6; i++ {
		g.addVertex(i)
	}
	g.addEdge(1, 2)
	g.addEdge(2, 2)
	g.addEdge(3, 4)
	g.addEdge(4, 1)
	g.addEdge(4, 2)
	g.PrintVertex()

}
