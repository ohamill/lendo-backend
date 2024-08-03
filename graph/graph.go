// Package graph exposes a graph data structure and an API for reading and updating the graph
package graph

import (
	"fmt"
	"lendo-backend/lendoErrors"
	"strings"
)

type Graph struct {
	// Edges stores mapping from each word to its list of synonyms
	Edges map[string][]string
}

func New() *Graph {
	return &Graph{
		Edges: make(map[string][]string),
	}
}

// AddVertex adds a new vertex to the graph and assigns it an empty slice of values. The vertex must not already exist, or this method will return an error.
func (g *Graph) AddVertex(vertex string) error {
	if g.exists(vertex) {
		return lendoErrors.ErrVertexAlreadyExists
	}

	g.Edges[vertex] = make([]string, 0)
	g.print()
	return nil
}

// AddEdge adds a new edge to a vertex. The vertex must already exist, or else an error is returned. The new edge is also added to all existing edges of the vertex.
func (g *Graph) AddEdge(vertex, edge string) error {
	if !g.exists(vertex) {
		return lendoErrors.ErrVertexDoesNotExist
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

// GetVertexEdges returns all edges for the provided vertex. To distinguish between vertices with no edges and non-existent vertices, an error is returned if the provided vertex does not exist in the graph.
func (g *Graph) GetVertexEdges(vertex string) ([]string, error) {
	edges, ok := g.Edges[vertex]
	if !ok {
		return nil, lendoErrors.ErrVertexDoesNotExist
	}
	return edges, nil
}

// GetVertexes returns all vertices contained in the graph
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
