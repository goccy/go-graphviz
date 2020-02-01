package graphviz_test

import (
	"bytes"
	"log"
	"testing"

	"github.com/goccy/go-graphviz"
)

func TestGraphviz_PNG(t *testing.T) {
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
	if err := g.Render(graph, "png", &buf); err != nil {
		log.Fatalf("%+v", err)
	}
	if len(buf.Bytes()) != 4610 {
		t.Fatal("failed to encode png")
	}
}
