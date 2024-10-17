package main

import (
	"bytes"
	"context"
	"fmt"
	"log"

	"github.com/goccy/go-graphviz"
)

func _main(ctx context.Context) error {
	g, err := graphviz.New(ctx)
	if err != nil {
		return err
	}
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
	n, err := graph.CreateNodeByName("n")
	if err != nil {
		return err
	}
	m, err := graph.CreateNodeByName("m")
	if err != nil {
		return err
	}
	e, err := graph.CreateEdgeByName("e", n, m)
	if err != nil {
		return err
	}
	e.SetLabel("e")
	var buf bytes.Buffer
	if err := g.Render(ctx, graph, "dot", &buf); err != nil {
		log.Fatalf("%+v", err)
	}
	fmt.Println(buf.String())
	return nil
}

func main() {
	if err := _main(context.Background()); err != nil {
		log.Fatal(err)
	}
}
