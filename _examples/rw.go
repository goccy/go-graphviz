package main

import (
	"bytes"
	"log"

	"github.com/goccy/go-graphviz"
)

func renderDOTGraph() ([]byte, error) {
	g := graphviz.New()
	graph := g.Graph()
	defer func() {
		graph.Close()
		g.Close()
	}()
	n := graph.CreateNode("n")
	m := graph.CreateNode("m")
	e := graph.CreateEdge("e", n, m)
	e.SetLabel("e")
	var buf bytes.Buffer
	if err := g.Render(graph, "dot", &buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func main() {
	graphBytes, err := renderDOTGraph()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	graph := graphviz.ParseBytes(graphBytes)
	n := graph.Node("n")
	l := graph.CreateNode("l")
	e2 := graph.CreateEdge("e2", n, l)
	e2.SetLabel("e2")
	g := graphviz.New()
	defer func() {
		graph.Close()
		g.Close()
	}()
	g.RenderFilename(graph, "png", "rw.png")
}
