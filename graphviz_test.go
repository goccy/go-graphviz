package graphviz_test

import (
	"bytes"
	"testing"

	"github.com/goccy/go-graphviz"
)

func TestGraphviz_Image(t *testing.T) {
	g := graphviz.New()
	graph, err := g.Graph()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	defer func() {
		graph.Close()
		g.Close()
	}()
	n, err := graph.CreateNode("n")
	if err != nil {
		t.Fatalf("%+v", err)
	}
	m, err := graph.CreateNode("m")
	if err != nil {
		t.Fatalf("%+v", err)
	}
	e, err := graph.CreateEdge("e", n, m)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	e.SetLabel("e")

	t.Run("png", func(t *testing.T) {
		t.Run("Render", func(t *testing.T) {
			var buf bytes.Buffer
			if err := g.Render(graph, graphviz.PNG, &buf); err != nil {
				t.Fatalf("%+v", err)
			}
			if len(buf.Bytes()) != 4602 {
				t.Fatalf("failed to encode png: bytes length is %d", len(buf.Bytes()))
			}
		})
		t.Run("RenderImage", func(t *testing.T) {
			image, err := g.RenderImage(graph)
			if err != nil {
				t.Fatalf("%+v", err)
			}
			bounds := image.Bounds()
			if bounds.Max.X != 83 {
				t.Fatal("failed to get image")
			}
			if bounds.Max.Y != 177 {
				t.Fatal("failed to get image")
			}
		})
	})
	t.Run("jpg", func(t *testing.T) {
		t.Run("Render", func(t *testing.T) {
			var buf bytes.Buffer
			if err := g.Render(graph, graphviz.JPG, &buf); err != nil {
				t.Fatalf("%+v", err)
			}
			if len(buf.Bytes()) != 3296 {
				t.Fatalf("failed to encode jpg: bytes length is %d", len(buf.Bytes()))
			}
		})
		t.Run("RenderImage", func(t *testing.T) {
			image, err := g.RenderImage(graph)
			if err != nil {
				t.Fatalf("%+v", err)
			}
			bounds := image.Bounds()
			if bounds.Max.X != 83 {
				t.Fatal("failed to get image")
			}
			if bounds.Max.Y != 177 {
				t.Fatal("failed to get image")
			}
		})
	})
}

func TestGetNode(t *testing.T) {
	g := graphviz.New()
	graph, err := g.Graph()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	defer func() {
		graph.Close()
		g.Close()
	}()

	_, err = graph.CreateNode("n")
	if err != nil {
		t.Fatalf("%+v", err)
	}

	n, err := graph.GetNode("n")
	if err != nil {
		t.Fatalf("%+v", err)
	}
	if n == nil {
		t.Fatal("failed to get node 'n'. Got `nil` instead.")
	}
	if name := n.Name(); name != "n" {
		t.Fatalf("failed to get node 'n'. Got node '%s' instead.", name)
	}

	m, err := graph.GetNode("m")
	if err != nil {
		t.Fatalf("%+v", err)
	}
	if m != nil {
		t.Fatal("got unexpected node 'm'")
	}
	if nodes := graph.NumberNodes(); nodes != 1 {
		t.Fatalf("expected graph to have 1 node, but found %d", nodes)
	}
}

func TestGetEdge(t *testing.T) {
	g := graphviz.New()
	graph, err := g.Graph()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	defer func() {
		graph.Close()
		g.Close()
	}()

	n, err := graph.CreateNode("n")
	if err != nil {
		t.Fatalf("%+v", err)
	}
	m, err := graph.CreateNode("m")
	if err != nil {
		t.Fatalf("%+v", err)
	}

	_, err = graph.CreateEdge("e1", n, m)
	if err != nil {
		t.Fatalf("%+v", err)
	}

	e1, err := graph.GetEdge("e1", n, m)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	if e1 == nil {
		t.Fatal("failed to get edge 'e1'. Got `nil` instead.")
	}
	if name := e1.Name(); name != "e1" {
		t.Fatalf("failed to get node 'n'. Got node '%s' instead.", name)
	}

	e2, err := graph.GetEdge("e2", m, n)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	if e2 != nil {
		t.Fatal("got unexpected edge 'e2'")
	}
	if edges := graph.NumberEdges(); edges != 1 {
		t.Fatalf("expected graph to have 1 edge, but found %d", edges)
	}
}
