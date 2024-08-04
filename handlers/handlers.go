// Package handlers exposes handler functions to be used by the Gin router
package handlers

import (
	"github.com/gin-gonic/gin"
	"lendo-backend/data"
	"lendo-backend/graph"
	"net/http"
)

type Handlers struct {
	// The graph data structure to store the words and their synonyms
	Store *graph.Graph
}

// AddWord adds a new word to the graph. The request may optionally include the word's synonyms - if so, they are also added to the graph.
func (h *Handlers) AddWord(c *gin.Context) {
	request, err := data.DecodeJson[data.Word](c.Request.Body)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	err = h.Store.AddVertex(request.Word)
	if err != nil {
		c.Status(http.StatusConflict)
		return
	}

	c.JSON(http.StatusCreated, request)
}

// AddSynonym appends one or more synonyms to an existing word
func (h *Handlers) AddSynonym(c *gin.Context) {
	word := c.Param("word")
	if word == "" {
		c.Status(http.StatusBadRequest)
		return
	}

	synonym, err := data.DecodeJson[data.Synonym](c.Request.Body)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	err = h.Store.AddEdge(word, synonym.Synonym)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	wordSynonyms, err := h.Store.GetVertexEdges(word)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusCreated, data.NewCompleteWordInfo(word, wordSynonyms))
}

// GetSynonyms fetches all synonyms for a word
func (h *Handlers) GetSynonyms(c *gin.Context) {
	word := c.Param("word")
	if word == "" {
		c.Status(http.StatusBadRequest)
		return
	}

	synonyms, err := h.Store.GetVertexEdges(word)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, data.Synonyms{Synonyms: synonyms})
}
