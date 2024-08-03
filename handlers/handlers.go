// Package handlers exposes handler functions to be used by the Gin router
package handlers

import (
	"github.com/gin-gonic/gin"
	"lendo-backend/data"
	"lendo-backend/graph"
	"net/http"
)

// The graph data structure to store the words and their synonyms
var store *graph.Graph

func init() {
	store = graph.New()
}

// AddWord adds a new word to the graph. The request may optionally include the word's synonyms - if so, they are also added to the graph.
func AddWord(c *gin.Context) {
	request, err := data.DecodeJson[data.WordInfo](c.Request.Body)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	err = store.AddVertex(request.Word)
	if err != nil {
		c.Status(http.StatusConflict)
		return
	}

	for _, synonym := range request.Synonyms {
		err = store.AddEdge(request.Word, synonym)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}
	}

	c.JSON(http.StatusCreated, data.NewWordInfo(request.Word, request.Synonyms))
}

// AddSynonyms appends one or more synonyms to an existing word
func AddSynonyms(c *gin.Context) {
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
		err = store.AddEdge(word, synonym)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}
	}

	wordSynonyms, err := store.GetVertexEdges(word)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusCreated, data.NewWordInfo(word, wordSynonyms))
}

// GetSynonyms fetches all synonyms for a word
func GetSynonyms(c *gin.Context) {
	word := c.Param("word")
	if word == "" {
		c.Status(http.StatusBadRequest)
		return
	}

	synonyms, err := store.GetVertexEdges(word)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, data.NewSynonymsInfo(synonyms))
}

// GetWords fetches all words in the graph
func GetWords(c *gin.Context) {
	c.JSON(http.StatusOK, store.GetVertexes())
}
