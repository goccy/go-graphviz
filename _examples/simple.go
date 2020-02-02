package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/goccy/go-graphviz"
)

func _main() error {
	g := graphviz.New()
	graph, err := g.Graph()
	if err != nil {
		return err
	}
	defer func() {
		if err := graph.Close(); err != nil {
			log.Fatal(err)
		}
		g.Close()
	}()
	n, err := graph.CreateNode("n")
	if err != nil {
		return err
	}
	m, err := graph.CreateNode("m")
	if err != nil {
		return err
	}
	e, err := graph.CreateEdge("e", n, m)
	if err != nil {
		return err
	}
	e.SetLabel("e")
	var buf bytes.Buffer
	if err := g.Render(graph, "dot", &buf); err != nil {
		log.Fatalf("%+v", err)
	}
	fmt.Println(buf.String())
	return nil
}

func main() {
	if err := _main(); err != nil {
		log.Fatal(err)
	}
}
