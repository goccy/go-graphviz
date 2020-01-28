package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/goccy/go-graphviz"
)

func main() {
	g := graphviz.New()
	graph := g.Graph()
	defer func() {
		graph.Close()
		g.Close()
	}()
	n := graph.CreateNode("n")
	m := graph.CreateNode("m")
	graph.CreateEdge("e", n, m).SetLabel("e")
	var buf bytes.Buffer
	if err := g.Render(graph, "dot", &buf); err != nil {
		log.Fatalf("%+v", err)
	}
	fmt.Println(buf.String())
}
