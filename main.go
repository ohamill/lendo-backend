package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"lendo-backend/handlers"
	"log"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	r.POST("/word", handlers.AddWord)
	r.POST("/synonyms/:word", handlers.AddSynonyms)
	r.GET("/synonyms/:word", handlers.GetSynonyms)
	r.GET("/words", handlers.GetWords)

	log.Fatalln(r.Run("localhost:8080")) // listen and serve on 0.0.0.0:8080
}
