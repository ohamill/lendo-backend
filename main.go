package main

import (
	"lendo-backend/graph"
	"lendo-backend/router"
	"log"
)

func main() {
	g := graph.New()
	r := router.Setup(g)
	log.Fatalln(r.Run())
}
