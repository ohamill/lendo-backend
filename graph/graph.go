package graph

import (
	"fmt"
	"strings"
)

type Graph struct {
	Edges map[string][]string
}

func New() *Graph {
	return &Graph{
		Edges: make(map[string][]string),
	}
}

func (g *Graph) AddVertex(vertex string) error {
	if g.exists(vertex) {
		return fmt.Errorf("vertex %s already exists", vertex)
	}

	g.Edges[vertex] = make([]string, 0)
	g.print()
	return nil
}

func (g *Graph) AddEdge(vertex, edge string) error {
	if !g.exists(vertex) {
		return fmt.Errorf("vertex %s does not exist", vertex)
	}

	for _, e := range g.Edges[vertex] {
		g.Edges[e] = append(g.Edges[e], edge)
	}

	if !g.exists(edge) {
		g.Edges[edge] = g.Edges[vertex]
	}
	g.Edges[edge] = append(g.Edges[edge], vertex)

	g.Edges[vertex] = append(g.Edges[vertex], edge)
	g.print()
	return nil
}

func (g *Graph) GetVertexEdges(vertex string) ([]string, error) {
	edges, ok := g.Edges[vertex]
	if !ok {
		return nil, fmt.Errorf("vertex %s does not exist", vertex)
	}
	return edges, nil
}

func (g *Graph) GetVertexes() []string {
	vertexes := make([]string, len(g.Edges))

	i := 0
	for vertex := range g.Edges {
		vertexes[i] = vertex
		i++
	}

	return vertexes
}

func (g *Graph) exists(vertex string) bool {
	_, ok := g.Edges[vertex]
	return ok
}

func (g *Graph) print() {
	for vertex, edges := range g.Edges {
		fmt.Printf("%s: %s\n", vertex, strings.Join(edges, " -> "))
	}
}
