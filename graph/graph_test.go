package graph

import (
	"errors"
	"lendo-backend/lendoErrors"
	"testing"
)

var vertex = "happy"
var edge = "glad"

func TestGraph_New(t *testing.T) {
	g := New()
	if len(g.Edges) != 0 {
		t.Fatalf("graph.New() vertices expected 0, got %d\n", len(g.Edges))
	}
}

func TestGraph_AddVertex(t *testing.T) {
	g := New()

	err := g.AddVertex(vertex)
	if err != nil {
		t.Fatalf("error adding vertex: %v\n", err)
	}

	if len(g.Edges) != 1 {
		t.Fatalf("graph.Edges length expected 1, got %d\n", len(g.Edges))
	}
	if _, ok := g.Edges[vertex]; !ok {
		t.Fatalf("graph.Edges does not contain expected vertex '%s'\n", vertex)
	}
}

func TestGraph_AddVertex_Err(t *testing.T) {
	g := New()

	// First time should succeed
	err := g.AddVertex(vertex)
	if err != nil {
		t.Fatalf("error adding vertex: %v\n", err)
	}
	// Second time should fail
	err = g.AddVertex(vertex)
	if err == nil {
		t.Fatalf("g.AddVertex() expected to fail because vertex already exists, but it did not.\n")
	}

	if !errors.Is(err, lendoErrors.ErrVertexAlreadyExists) {
		t.Fatalf("g.AddVertex() returned unexpected error. Expected ErrVertexAlreadyExists, got %+v\n", err)
	}
}

func TestGraph_AddEdge_EmptyEdges(t *testing.T) {
	g := New()

	err := g.AddVertex(vertex)
	if err != nil {
		t.Fatalf("g.AddVertex failed unexpectedly: %v\n", err)
	}

	err = g.AddEdge(vertex, edge)
	if err != nil {
		t.Fatalf("g.AddEdge failed: %v\n", err)
	}

	// The edge should be added as a vertex
	if len(g.Edges) != 2 {
		t.Fatalf("g.Edges length expected 2, got %d\n", len(g.Edges))
	}
	if _, ok := g.Edges[edge]; !ok {
		t.Fatalf("g.Edges does not contain expected vertex '%s'\n", edge)
	}
	// The vertex and the edge should have each other as edges
	if !edgeExists(edge, g.Edges[vertex]) {
		t.Fatalf("vertex '%s' does not contain expected edge '%s'\n", vertex, edge)
	}
	if !edgeExists(vertex, g.Edges[edge]) {
		t.Fatalf("vertex '%s' does not contain expected edge '%s'\n", edge, vertex)
	}
}

func TestGraph_AddEdge_OneEdge(t *testing.T) {
	g := New()

	err := g.AddVertex(vertex)
	if err != nil {
		t.Fatalf("g.AddVertex failed unexpectedly: %v\n", err)
	}
	err = g.AddEdge(vertex, "joyous")
	if err != nil {
		t.Fatalf("g.AddEdge failed unexpectedly: %v\n", err)
	}

	err = g.AddEdge(vertex, edge)
	if err != nil {
		t.Fatalf("g.AddEdge failed: %v\n", err)
	}

	// The edge should be added as a vertex
	if len(g.Edges) != 3 {
		t.Fatalf("g.Edges length expected 2, got %d\n", len(g.Edges))
	}
	if _, ok := g.Edges[edge]; !ok {
		t.Fatalf("g.Edges does not contain expected vertex '%s'\n", edge)
	}
	// The vertex, its existing edge, and the new edge should all have each other as edges
	if !edgeExists(edge, g.Edges[vertex]) {
		t.Fatalf("vertex '%s' does not contain expected edge '%s'\n", vertex, edge)
	}
	if !edgeExists(edge, g.Edges["joyous"]) {
		t.Fatalf("vertex '%s' does not contain expected edge '%s'\n", vertex, edge)
	}
	if !edgeExists(vertex, g.Edges[edge]) {
		t.Fatalf("vertex '%s' does not contain expected edge '%s'\n", edge, vertex)
	}
	if !edgeExists("joyous", g.Edges[edge]) {
		t.Fatalf("vertex '%s' does not contain expected edge 'joyous'\n", edge)
	}
}

func TestGraph_GetVertexEdges(t *testing.T) {
	g := New()

	err := g.AddVertex(vertex)
	if err != nil {
		t.Fatalf("g.AddVertex failed unexpectedly: %v\n", err)
	}
	err = g.AddEdge(vertex, edge)
	if err != nil {
		t.Fatalf("g.AddEdge failed unexpectedly: %v\n", err)
	}

	edges, err := g.GetVertexEdges(vertex)
	if err != nil {
		t.Fatalf("g.GetVertexEdges failed: %v\n", err)
	}

	if len(edges) != 1 {
		t.Fatalf("g.GetVertexEdges length expected 1, got %d\n", len(edges))
	}
	if !edgeExists(edge, edges) {
		t.Fatalf("g.GetVertexEdges did not contain expected edge '%s'\n", edge)
	}
}

func edgeExists(edge string, edges []string) bool {
	found := false
	for _, e := range edges {
		if e == edge {
			found = true
		}
	}
	return found
}
