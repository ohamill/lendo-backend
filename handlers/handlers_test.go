package handlers_test

import (
	"github.com/gin-gonic/gin"
	"lendo-backend/data"
	"lendo-backend/graph"
	"lendo-backend/router"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAddWord(t *testing.T) {
	g := graph.New()
	req, _ := http.NewRequest(http.MethodPost, "/word", strings.NewReader(`{"word": "happy"}`))
	rr := executeRequest(g, req)

	if rr.Code != http.StatusCreated {
		t.Fatalf("AddWord response status code expected %d, got %d", http.StatusCreated, rr.Code)
	}
	respBody, err := data.DecodeJson[data.Word](rr.Body)
	if err != nil {
		t.Fatalf("error decoding response body: %s", err)
	}
	if respBody.Word != "happy" {
		t.Fatalf("respBody.Word expected 'happy', got %s", respBody)
	}
}

func TestAddSynonym(t *testing.T) {
	g := graph.New()
	err := g.AddVertex("happy")
	if err != nil {
		t.Fatalf("error adding vertex: %s", err)
	}
	req, _ := http.NewRequest(http.MethodPost, "/synonym/happy", strings.NewReader(`{"synonym": "glad"}`))
	rr := executeRequest(g, req)

	if rr.Code != http.StatusCreated {
		t.Fatalf("AddSynonym response status code expected %d, got %d", http.StatusCreated, rr.Code)
	}
}

func TestAddSynonyms_MalformedRequest(t *testing.T) {
	g := graph.New()
	err := g.AddVertex("happy")
	if err != nil {
		t.Fatalf("error adding vertex: %s", err)
	}
	req, _ := http.NewRequest(http.MethodPost, "/synonym/happy", strings.NewReader(`{}`))
	rr := executeRequest(g, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("AddSynonym response status code expected %d, got %d", http.StatusBadRequest, rr.Code)
	}
}

func TestGetSynonyms_ZeroSynonyms(t *testing.T) {
	g := graph.New()
	err := g.AddVertex("happy")
	if err != nil {
		t.Fatalf("error adding vertex: %s", err)
	}
	req, _ := http.NewRequest(http.MethodGet, "/synonyms/happy", nil)
	rr := executeRequest(g, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("GetSynonyms response status code expected %d, got %d", http.StatusOK, rr.Code)
	}
	respBody, err := data.DecodeJson[data.Synonyms](rr.Body)
	if err != nil {
		t.Fatalf("error decoding response body: %s", err)
	}
	if len(respBody.Synonyms) != 0 {
		t.Fatalf("respBody.Synonyms length expected 0, got %d", len(respBody.Synonyms))
	}
}

func TestGetSynonyms(t *testing.T) {
	g := graph.New()
	err := g.AddVertex("happy")
	if err != nil {
		t.Fatalf("error adding vertex: %s", err)
	}
	err = g.AddEdge("happy", "glad")
	if err != nil {
		t.Fatalf("error adding edge: %s", err)
	}
	req, _ := http.NewRequest(http.MethodGet, "/synonyms/happy", nil)
	rr := executeRequest(g, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("GetSynonyms response status code expected %d, got %d", http.StatusOK, rr.Code)
	}
	respBody, err := data.DecodeJson[data.Synonyms](rr.Body)
	if err != nil {
		t.Fatalf("error decoding response body: %s", err)
	}
	if len(respBody.Synonyms) != 1 {
		t.Fatalf("respBody.Synonyms length expected 0, got %d", len(respBody.Synonyms))
	}
	if respBody.Synonyms[0] != "glad" {
		t.Fatalf("respBody.Synonyms[0] expected 'glad', got '%s'", respBody.Synonyms[0])
	}
}

func executeRequest(g *graph.Graph, req *http.Request) *httptest.ResponseRecorder {
	gin.SetMode(gin.TestMode)
	r := router.Setup(g)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
