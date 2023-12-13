package graphviz_test

import (
	"bytes"
	"io/ioutil"
	"os"
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

func TestParseBytes(t *testing.T) {
	type test struct {
		input          string
		expected_valid bool
	}

	tests := []test{
		{input: "graph test { a -- b }", expected_valid: true},
		{input: "graph test { a -- b", expected_valid: false},
		{input: "graph test { a -- b }", expected_valid: true},
		{input: "graph test { a -- }", expected_valid: false},
		{input: "graph test { a -- c }", expected_valid: true},
		{input: "graph test { a - b }", expected_valid: false},
		{input: "graph test { d -- e }", expected_valid: true},
	}

	for i, test := range tests {
		_, err := graphviz.ParseBytes([]byte(test.input))
		actual_valid := err == nil
		if actual_valid != test.expected_valid {
			t.Errorf("Test %d of TestParseBytes failed. Parsing error: %+v", i+1, err)
		}
	}
}

func TestParseFile(t *testing.T) {
	type test struct {
		input          string
		expected_valid bool
	}

	tests := []test{
		{input: "graph test { a -- b }", expected_valid: true},
		{input: "graph test { a -- b", expected_valid: false},
		{input: "graph test { a -- b }", expected_valid: true},
		{input: "graph test { a -- }", expected_valid: false},
		{input: "graph test { a -- c }", expected_valid: true},
		{input: "graph test { a - b }", expected_valid: false},
		{input: "graph test { d -- e }", expected_valid: true},
	}

	createTempFile := func(t *testing.T, content string) *os.File {
		file, err := ioutil.TempFile("", "*")
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

	for i, test := range tests {
		tmpfile := createTempFile(t, test.input)
		defer os.Remove(tmpfile.Name())

		_, err := graphviz.ParseFile(tmpfile.Name())
		actual_valid := err == nil
		if actual_valid != test.expected_valid {
			t.Errorf("Test %d of TestParseFile failed. Parsing error: %+v", i+1, err)
		}
	}
}

func TestNodeDegree(t *testing.T) {
	type test struct {
		node_name             string
		expected_indegree     int
		expected_outdegree    int
		expected_total_degree int
	}

	type graphtest struct {
		input string
		tests []test
	}

	graphtests := []graphtest{
		{input: "digraph test { a -> b }", tests: []test{
			{node_name: "a", expected_indegree: 0, expected_outdegree: 1, expected_total_degree: 1},
			{node_name: "b", expected_indegree: 1, expected_outdegree: 0, expected_total_degree: 1},
		}},
		{input: "digraph test { a -> b; a -> b; a -> a; c -> a }", tests: []test{
			{node_name: "a", expected_indegree: 2, expected_outdegree: 3, expected_total_degree: 5},
			{node_name: "b", expected_indegree: 2, expected_outdegree: 0, expected_total_degree: 2},
			{node_name: "c", expected_indegree: 0, expected_outdegree: 1, expected_total_degree: 1},
		}},
		{input: "graph test { a -- b; a -- b; a -- a; c -- a }", tests: []test{
			{node_name: "a", expected_indegree: 2, expected_outdegree: 3, expected_total_degree: 5},
			{node_name: "b", expected_indegree: 2, expected_outdegree: 0, expected_total_degree: 2},
			{node_name: "c", expected_indegree: 0, expected_outdegree: 1, expected_total_degree: 1},
		}},
		{input: "strict graph test { a -- b; b -- a; a -- a; c -- a }", tests: []test{
			{node_name: "a", expected_indegree: 2, expected_outdegree: 2, expected_total_degree: 4},
			{node_name: "b", expected_indegree: 1, expected_outdegree: 0, expected_total_degree: 1},
			{node_name: "c", expected_indegree: 0, expected_outdegree: 1, expected_total_degree: 1},
		}},
	}

	for _, graphtest := range graphtests {
		input := graphtest.input
		graph, err := graphviz.ParseBytes([]byte(input))
		if err != nil {
			t.Fatalf("Input: %s. Error: %+v", input, err)
		}

		for _, test := range graphtest.tests {
			node_name := test.node_name
			node, err := graph.Node(node_name)
			if err != nil || node == nil {
				t.Fatalf("Unable to retrieve node '%s'. Input: %s. Error: %+v", node_name, input, err)
			}

			indegree := graph.Indegree(node)
			if test.expected_indegree != indegree {
				t.Errorf("Unexpected indegree for node '%s'. Input: %s. Expected: %d. Actual: %d.", node_name, input, test.expected_indegree, indegree)
			}
			outdegree := graph.Outdegree(node)
			if test.expected_outdegree != outdegree {
				t.Errorf("Unexpected outdegree for node '%s'. Input: %s. Expected: %d. Actual: %d.", node_name, input, test.expected_outdegree, outdegree)
			}
			total_degree := graph.TotalDegree(node)
			if test.expected_total_degree != total_degree {
				t.Errorf("Unexpected total degree for node '%s'. Input: %s. Expected: %d. Actual: %d.", node_name, input, test.expected_total_degree, total_degree)
			}
		}
	}
}
