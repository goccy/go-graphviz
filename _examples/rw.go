package main

import (
	"bytes"
	"log"

	"github.com/goccy/go-graphviz"
)

func renderDOTGraph() ([]byte, error) {
	g := graphviz.New()
	graph, err := g.Graph()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := graph.Close(); err != nil {
			log.Fatal(err)
		}
		g.Close()
	}()
	n, err := graph.CreateNode("n")
	if err != nil {
		return nil, err
	}
	m, err := graph.CreateNode("m")
	if err != nil {
		return nil, err
	}
	e, err := graph.CreateEdge("e", n, m)
	if err != nil {
		return nil, err
	}
	e.SetLabel("e")
	var buf bytes.Buffer
	if err := g.Render(graph, "dot", &buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func _main() error {
	graphBytes, err := renderDOTGraph()
	if err != nil {
		return err
	}
	graph, err := graphviz.ParseBytes(graphBytes)
	if err != nil {
		return err
	}
	n, err := graph.Node("n")
	if err != nil {
		return err
	}
	l, err := graph.CreateNode("l")
	if err != nil {
		return err
	}
	e2, err := graph.CreateEdge("e2", n, l)
	if err != nil {
		return err
	}
	e2.SetLabel("e2")
	g := graphviz.New()
	defer func() {
		if err := graph.Close(); err != nil {
			log.Fatal(err)
		}
		g.Close()
	}()
	g.RenderFilename(graph, "png", "rw.png")
	return nil
}

func main() {
	if err := _main(); err != nil {
		log.Fatalf("%+v", err)
	}
}
