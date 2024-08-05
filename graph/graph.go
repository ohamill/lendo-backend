// Package graph exposes a graph data structure and an API for reading and updating the graph
package graph

import (
	"lendo-backend/lendoErrors"
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

func (g *Graph) exists(vertex string) bool {
	_, ok := g.Edges[vertex]
	return ok
}
