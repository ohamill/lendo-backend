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
	request, err := data.DecodeJson[data.WordInfo](c.Request.Body)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	err = h.Store.AddVertex(request.Word)
	if err != nil {
		c.Status(http.StatusConflict)
		return
	}

	for _, synonym := range request.Synonyms {
		err = h.Store.AddEdge(request.Word, synonym)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}
	}

	c.JSON(http.StatusCreated, data.NewWordInfo(request.Word, request.Synonyms))
}

// AddSynonyms appends one or more synonyms to an existing word
func (h *Handlers) AddSynonyms(c *gin.Context) {
	word := c.Param("word")
	if word == "" {
		c.Status(http.StatusBadRequest)
		return
	}

	synonyms, err := data.DecodeJson[data.SynonymsInfo](c.Request.Body)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	for _, synonym := range synonyms.Synonyms {
		err = h.Store.AddEdge(word, synonym)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}
	}

	wordSynonyms, err := h.Store.GetVertexEdges(word)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusCreated, data.NewWordInfo(word, wordSynonyms))
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
	c.JSON(http.StatusOK, data.NewSynonymsInfo(synonyms))
}

// GetWords fetches all words in the graph
func (h *Handlers) GetWords(c *gin.Context) {
	vertices := h.Store.GetVertexes()
	c.JSON(http.StatusOK, data.NewWordsInfo(vertices))
}
