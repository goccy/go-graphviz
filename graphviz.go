package graphviz

import (
	"image"
	"io"

	"github.com/goccy/go-graphviz/cgraph"
	"github.com/goccy/go-graphviz/gvc"
)

type Graphviz struct {
	ctx    *gvc.Context
	name   string
	dir    *cgraph.Desc
	layout Layout
}

type Graph struct {
	graph *cgraph.Graph
}

type Node struct {
	node *cgraph.Node
}

type Edge struct {
	edge *cgraph.Edge
}

type Layout string

const (
	CIRCO     Layout = "circo"
	DOT       Layout = "dot"
	FDP       Layout = "fdp"
	NEATO     Layout = "neato"
	OSAGE     Layout = "osage"
	PATCHWORK Layout = "patchwork"
	SFDP      Layout = "sfdp"
	TWOPI     Layout = "twopi"
)

func ParseFile(path string) (*Graph, error) {
	graph, err := cgraph.ParseFile(path)
	if err != nil {
		return nil, err
	}
	return toGraph(graph), nil
}

func ParseBytes(bytes []byte) (*Graph, error) {
	graph, err := cgraph.ParseBytes(bytes)
	if err != nil {
		return nil, err
	}
	return toGraph(graph), nil
}

func New() *Graphviz {
	return &Graphviz{
		ctx:    gvc.New(),
		dir:    cgraph.Directed,
		layout: DOT,
	}
}

func toGraph(g *cgraph.Graph) *Graph {
	if g == nil {
		return nil
	}
	return &Graph{graph: g}
}

func toNode(n *cgraph.Node) *Node {
	if n == nil {
		return nil
	}
	return &Node{node: n}
}

func toEdge(e *cgraph.Edge) *Edge {
	if e == nil {
		return nil
	}
	return &Edge{edge: e}
}

func (g *Graphviz) Close() {
	g.ctx.Close()
}

func (g *Graphviz) SetLayout(layout Layout) *Graphviz {
	g.layout = layout
	return g
}

func (g *Graphviz) Render(graph *Graph, format string, w io.Writer) error {
	g.ctx.Layout(graph.graph, string(g.layout))
	defer g.ctx.FreeLayout(graph.graph)

	if err := g.ctx.RenderData(graph.graph, format, w); err != nil {
		return err
	}
	return nil
}

func (g *Graphviz) RenderImage(graph *Graph, format string) image.Image {
	g.ctx.Layout(graph.graph, string(g.layout))
	defer g.ctx.FreeLayout(graph.graph)

	return g.ctx.RenderImage(graph.graph, format)
}

func (g *Graphviz) RenderFilename(graph *Graph, format, path string) error {
	g.ctx.Layout(graph.graph, string(g.layout))
	defer g.ctx.FreeLayout(graph.graph)

	g.ctx.RenderFilename(graph.graph, format, path)
	return nil
}

func (g *Graphviz) Graph(option ...GraphOption) (*Graph, error) {
	for _, opt := range option {
		opt(g)
	}
	graph, err := cgraph.Open(g.name, g.dir, nil)
	if err != nil {
		return nil, err
	}
	return toGraph(graph), nil
}

func (g *Graph) Close() error {
	return g.graph.Close()
}

func (g *Graph) Node(id string) (*Node, error) {
	node, err := g.graph.Node(id, 0)
	if err != nil {
		return nil, err
	}
	return toNode(node), nil
}

func (g *Graph) CreateNode(id string) (*Node, error) {
	node, err := g.graph.Node(id, 1)
	if err != nil {
		return nil, err
	}
	return toNode(node), nil
}

func (g *Graph) CreateEdge(id string, start *Node, end *Node) (*Edge, error) {
	edge, err := g.graph.Edge(start.node, end.node, id, 1)
	if err != nil {
		return nil, err
	}
	return toEdge(edge).SetLabel(""), nil
}
