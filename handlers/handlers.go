package handlers

import (
	"github.com/gin-gonic/gin"
	"lendo-backend/data"
	"lendo-backend/graph"
	"net/http"
)

var store *graph.Graph

func init() {
	store = graph.New()
}

func AddWord(c *gin.Context) {
	request, err := data.UnmarshalJson[data.CreateWordRequest](c.Request.Body)
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

	c.JSON(http.StatusCreated, data.CreateWordResponse{
		Word:     request.Word,
		Synonyms: request.Synonyms,
	})
}

func AddSynonyms(c *gin.Context) {
	word := c.Param("word")
	if word == "" {
		c.Status(http.StatusBadRequest)
		return
	}

	synonyms, err := data.UnmarshalJson[data.CreateSynonymsRequest](c.Request.Body)
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
	c.JSON(http.StatusCreated, data.CreateSynonymsResponse{
		Word:     word,
		Synonyms: wordSynonyms,
	})
}

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
	c.JSON(http.StatusOK, data.GetSynonymsResponse{
		Synonyms: synonyms,
	})
}

func GetWords(c *gin.Context) {
	c.JSON(http.StatusOK, store.GetVertexes())
}
