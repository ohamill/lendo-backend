package main

import (
	"lendo-backend/graph"
	"lendo-backend/router"
	"log"
)

func main() {
	g := graph.New()
	r := router.Setup(g)
	log.Fatalln(r.Run("localhost:8080")) // listen and serve on 0.0.0.0:8080
}
