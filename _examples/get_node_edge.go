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

	_, err = graph.CreateEdge("e", n, m)
	if err != nil {
		return err
	}

	_, err = graph.CreateEdge("f", m, n)
	if err != nil {
		return err
	}

	node1, err := graph.GetNode("n")
	if err != nil {
		return err
	}
	if node1 != nil {
		node1.SetLabel("Node 1 (n)")
	}

	node2, err := graph.GetNode("z")
	if err != nil {
		return err
	}
	// should not run
	if node2 != nil {
		node2.SetLabel("Node 2 (z)")
	}

	e1, err := graph.GetEdge("e", n, m)
	if err != nil {
		return err
	}
	if e1 != nil {
		e1.SetLabel("e")
		e1.SetColor("red")
	}

	e2, err := graph.GetEdge("g", n, m)
	if err != nil {
		return err
	}
	// should not run
	if e2 != nil {
		e2.SetLabel("g")
		e1.SetColor("blue")
	}

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
