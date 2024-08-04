package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"lendo-backend/graph"
	"lendo-backend/handlers"
)

func Setup(store *graph.Graph) *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())

	env := &handlers.Handlers{Store: store}

	r.POST("/word", env.AddWord)
	r.POST("/synonym/:word", env.AddSynonym)
	r.GET("/synonyms/:word", env.GetSynonyms)

	return r
}
