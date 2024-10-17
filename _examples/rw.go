package main

import (
	"bytes"
	"context"
	"log"

	"github.com/goccy/go-graphviz"
)

func renderDOTGraph(ctx context.Context) ([]byte, error) {
	g, err := graphviz.New(ctx)
	if err != nil {
		return nil, err
	}
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
	n, err := graph.CreateNodeByName("n")
	if err != nil {
		return nil, err
	}
	m, err := graph.CreateNodeByName("m")
	if err != nil {
		return nil, err
	}
	e, err := graph.CreateEdgeByName("e", n, m)
	if err != nil {
		return nil, err
	}
	e.SetLabel("e")
	var buf bytes.Buffer
	if err := g.Render(ctx, graph, "dot", &buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func _main(ctx context.Context) error {
	graphBytes, err := renderDOTGraph(ctx)
	if err != nil {
		return err
	}
	graph, err := graphviz.ParseBytes(graphBytes)
	if err != nil {
		return err
	}
	n, err := graph.NodeByName("n")
	if err != nil {
		return err
	}
	l, err := graph.CreateNodeByName("l")
	if err != nil {
		return err
	}
	e2, err := graph.CreateEdgeByName("e2", n, l)
	if err != nil {
		return err
	}
	e2.SetLabel("e2")
	g, err := graphviz.New(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err := graph.Close(); err != nil {
			log.Fatal(err)
		}
		g.Close()
	}()
	g.RenderFilename(ctx, graph, "png", "rw.png")
	return nil
}

func main() {
	if err := _main(context.Background()); err != nil {
		log.Fatalf("%+v", err)
	}
}
