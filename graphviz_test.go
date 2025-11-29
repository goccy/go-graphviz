package graphviz_test

import (
	"bytes"
	"context"
	"embed"
	"io/fs"
	"os"
	"path/filepath"
	"testing"

	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
)

func TestGraphviz_Image(t *testing.T) {
	ctx := context.Background()
	g, err := graphviz.New(ctx)
	if err != nil {
		t.Fatal(err)
	}
	graph, err := g.Graph()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	defer func() {
		graph.Close()
		g.Close()
	}()
	n, err := graph.CreateNodeByName("n")
	if err != nil {
		t.Fatalf("%+v", err)
	}
	m, err := graph.CreateNodeByName("m")
	if err != nil {
		t.Fatalf("%+v", err)
	}
	e, err := graph.CreateEdgeByName("e", n, m)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	e.SetLabel("e")

	t.Run("png", func(t *testing.T) {
		t.Run("Render", func(t *testing.T) {
			var buf bytes.Buffer
			if err := g.Render(ctx, graph, graphviz.PNG, &buf); err != nil {
				t.Fatalf("failed to render: %+v", err)
			}
			if len(buf.Bytes()) == 0 {
				t.Fatal("failed to encode png")
			}
		})
		t.Run("RenderImage", func(t *testing.T) {
			image, err := g.RenderImage(ctx, graph)
			if err != nil {
				t.Fatalf("%+v", err)
			}
			bounds := image.Bounds()
			if bounds.Max.X != 83 {
				t.Fatalf("expected bounds x is %d. but got %d", 83, bounds.Max.X)
			}
			if bounds.Max.Y != 177 {
				t.Fatalf("expected bounds y is %d. but got %d", 177, bounds.Max.Y)
			}
		})
	})
	t.Run("jpg", func(t *testing.T) {
		t.Run("Render", func(t *testing.T) {
			var buf bytes.Buffer
			if err := g.Render(ctx, graph, graphviz.JPG, &buf); err != nil {
				t.Fatalf("%+v", err)
			}
			if len(buf.Bytes()) == 0 {
				t.Fatal("failed to encode jpg")
			}
		})
		t.Run("RenderImage", func(t *testing.T) {
			image, err := g.RenderImage(ctx, graph)
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

func TestParseBytes(t *testing.T) {
	type test struct {
		input       string
		expectedErr bool
	}

	tests := []test{
		{input: "graph test1 { a -- b }"},
		{input: "graph test2 { a -- b", expectedErr: true},
		{input: "graph test3 { a -- b }"},
		{input: "graph test4 { a -- }", expectedErr: true},
		{input: "graph test5 { a -- c }"},
		{input: "graph test6 { a - b }", expectedErr: true},
		{input: "graph test7 { d -- e }"},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			_, err := graphviz.ParseBytes([]byte(test.input))
			if test.expectedErr && err == nil {
				t.Fatal("expected parsing error")
			} else if !test.expectedErr && err != nil {
				t.Fatalf("failed to parse: %v", err)
			}
		})
	}
}

func TestParseFile(t *testing.T) {
	type test struct {
		input       string
		expectedErr bool
	}

	tests := []test{
		{input: "graph test1 { a -- b }"},
		{input: "graph test2 { a -- b", expectedErr: true},
		{input: "graph test3 { a -- b }"},
		{input: "graph test4 { a -- }", expectedErr: true},
		{input: "graph test5 { a -- c }"},
		{input: "graph test6 { a - b }", expectedErr: true},
		{input: "graph test7 { d -- e }"},
	}

	createTempFile := func(t *testing.T, content string) *os.File {
		file, err := os.CreateTemp("", "*")
		if err != nil {
			t.Fatalf("There was an error creating a temporary file. Error: %+v", err)
			return nil
		}
		_, err = file.WriteString(content)
		if err != nil {
			t.Fatalf("There was an error writing '%s' to a temporary file. Error: %+v", content, err)
			return nil
		}
		return file
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			tmpfile := createTempFile(t, test.input)
			defer os.Remove(tmpfile.Name())

			_, err := graphviz.ParseFile(tmpfile.Name())
			if test.expectedErr && err == nil {
				t.Fatal("expected parsing error")
			} else if !test.expectedErr && err != nil {
				t.Fatalf("failed to parse: %v", err)
			}
		})
	}
}

//go:embed testdata/logo.png
var logoFS embed.FS

type imageFS struct{}

func (fs *imageFS) Open(name string) (fs.File, error) {
	return logoFS.Open(filepath.Join("testdata", name))
}

func TestImageRender(t *testing.T) {
	ctx := context.Background()
	graphviz.SetFileSystem(new(imageFS))

	g, err := graphviz.New(ctx)
	if err != nil {
		t.Fatal(err)
	}
	graph, err := g.Graph()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	defer func() {
		graph.Close()
		g.Close()
	}()
	n, err := graph.CreateNodeByName("n")
	if err != nil {
		t.Fatalf("%+v", err)
	}
	n.SetLabel("")

	// specify dummy image path.
	// Normally, a path to `testdata` would be required before `logo.png`,
	// but we confirm that the image can be loaded by appending the path to `testdata` within the `imageFS` specified by graphviz.SetFileSystem function.
	// This test is to verify that images can be loaded relative to the specified FileSystem.
	n.SetImage("logo.png")
	m, err := graph.CreateNodeByName("m")
	if err != nil {
		t.Fatalf("%+v", err)
	}
	if _, err := graph.CreateEdgeByName("e", n, m); err != nil {
		t.Fatalf("%+v", err)
	}
	var buf bytes.Buffer
	if err := g.Render(ctx, graph, "png", &buf); err != nil {
		t.Fatal(err)
	}
	if len(buf.Bytes()) == 0 {
		t.Fatal("failed to render image")
	}
}

func TestNodeDegree(t *testing.T) {
	type test struct {
		nodeName            string
		expectedIndegree    int
		expectedOutdegree   int
		expectedTotalDegree int
	}

	type graphtest struct {
		input string
		tests []test
	}

	graphtests := []graphtest{
		{input: "digraph test { a -> b }", tests: []test{
			{nodeName: "a", expectedIndegree: 0, expectedOutdegree: 1, expectedTotalDegree: 1},
			{nodeName: "b", expectedIndegree: 1, expectedOutdegree: 0, expectedTotalDegree: 1},
		}},
		{input: "digraph test { a -> b; a -> b; a -> a; c -> a }", tests: []test{
			{nodeName: "a", expectedIndegree: 2, expectedOutdegree: 3, expectedTotalDegree: 5},
			{nodeName: "b", expectedIndegree: 2, expectedOutdegree: 0, expectedTotalDegree: 2},
			{nodeName: "c", expectedIndegree: 0, expectedOutdegree: 1, expectedTotalDegree: 1},
		}},
		{input: "graph test { a -- b; a -- b; a -- a; c -- a }", tests: []test{
			{nodeName: "a", expectedIndegree: 2, expectedOutdegree: 3, expectedTotalDegree: 5},
			{nodeName: "b", expectedIndegree: 2, expectedOutdegree: 0, expectedTotalDegree: 2},
			{nodeName: "c", expectedIndegree: 0, expectedOutdegree: 1, expectedTotalDegree: 1},
		}},
		{input: "strict graph test { a -- b; b -- a; a -- a; c -- a }", tests: []test{
			{nodeName: "a", expectedIndegree: 2, expectedOutdegree: 2, expectedTotalDegree: 4},
			{nodeName: "b", expectedIndegree: 1, expectedOutdegree: 0, expectedTotalDegree: 1},
			{nodeName: "c", expectedIndegree: 0, expectedOutdegree: 1, expectedTotalDegree: 1},
		}},
	}

	for _, graphtest := range graphtests {
		input := graphtest.input
		graph, err := graphviz.ParseBytes([]byte(input))
		if err != nil {
			t.Fatalf("Input: %s. Error: %+v", input, err)
		}

		for _, test := range graphtest.tests {
			nodeName := test.nodeName
			node, err := graph.NodeByName(nodeName)
			if err != nil || node == nil {
				t.Fatalf("Unable to retrieve node '%s'. Input: %s. Error: %+v", nodeName, input, err)
			}

			indegree, err := graph.Indegree(node)
			if err != nil {
				t.Fatal(err)
			}
			if test.expectedIndegree != indegree {
				t.Errorf("Unexpected indegree for node '%s'. Input: %s. Expected: %d. Actual: %d.", nodeName, input, test.expectedIndegree, indegree)
			}
			outdegree, err := graph.Outdegree(node)
			if err != nil {
				t.Fatal(err)
			}
			if test.expectedOutdegree != outdegree {
				t.Errorf("Unexpected outdegree for node '%s'. Input: %s. Expected: %d. Actual: %d.", nodeName, input, test.expectedOutdegree, outdegree)
			}
			totalDegree, err := graph.TotalDegree(node)
			if err != nil {
				t.Fatal(err)
			}
			if test.expectedTotalDegree != totalDegree {
				t.Errorf("Unexpected total degree for node '%s'. Input: %s. Expected: %d. Actual: %d.", nodeName, input, test.expectedTotalDegree, totalDegree)
			}
		}
	}
}

func TestEdgeSourceAndTarget(t *testing.T) {
	ctx := context.Background()
	graph, err := graphviz.New(ctx)
	if err != nil {
		t.Fatalf("Error: %+v", err)
	}

	g, err := graph.Graph()
	if err != nil {
		t.Fatalf("Error: %+v", err)
	}

	nodeA, err := g.CreateNodeByName("a")
	if err != nil {
		t.Fatalf("Error: %+v", err)
	}

	nodeB, err := g.CreateNodeByName("b")
	if err != nil {
		t.Fatalf("Error: %+v", err)
	}

	edge, err := g.CreateEdgeByName("edge", nodeA, nodeB)
	if err != nil {
		t.Fatalf("Error: %+v", err)
	}

	head, err := edge.Head()
	if err != nil {
		t.Fatalf("Error: %+v", err)
	}
	if head == nil {
		t.Fatalf("Source is nil")
	}

	headName, err := head.Name()
	if err != nil {
		t.Fatalf("Error: %+v", err)
	}

	if headName != "b" {
		t.Fatalf("Expected source name to be 'b', got '%s'", headName)
	}

	target, err := edge.Tail()
	if err != nil {
		t.Fatalf("Error: %+v", err)
	}
	if target == nil {
		t.Fatalf("Target is nil")
	}

	tailName, err := target.Name()
	if err != nil {
		t.Fatalf("Error: %+v", err)
	}

	if tailName != "a" {
		t.Fatalf("Expected target name to be 'a', got '%s'", tailName)
	}
}

func TestDefaultNodeAttr(t *testing.T) {
	ctx := context.Background()
	graph, err := graphviz.New(ctx)
	if err != nil {
		t.Fatalf("Error: %+v", err)
	}

	g, err := graph.Graph()
	if err != nil {
		t.Fatalf("Error: %+v", err)
	}

	err = g.SetDefaultAttr(cgraph.KindGraph, "bgcolor", "red")
	if err != nil {
		t.Fatalf("Error: %+v", err)
	}

	err = g.SetDefaultAttr(cgraph.KindNode, "shape", "box")
	if err != nil {
		t.Fatalf("Error: %+v", err)
	}

	err = g.SetDefaultAttr(cgraph.KindEdge, "color", "blue")
	if err != nil {
		t.Fatalf("Error: %+v", err)
	}

	bgcolor := g.GetStr("bgcolor")
	if bgcolor != "red" {
		t.Fatalf("Expected bgcolor to be 'red', got '%s'", bgcolor)
	}

	n, err := g.CreateNodeByName("n")
	if err != nil {
		t.Fatalf("Error: %+v", err)
	}
	shape := n.GetStr("shape")
	if shape != "box" {
		t.Fatalf("Expected shape to be 'box', got '%s'", shape)
	}

	e, err := g.CreateEdgeByName("e", n, n)
	if err != nil {
		t.Fatalf("Error: %+v", err)
	}
	color := e.GetStr("color")
	if color != "blue" {
		t.Fatalf("Expected color to be 'blue', got '%s'", color)
	}
}
