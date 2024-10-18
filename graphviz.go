package graphviz

import (
	"context"
	"image"
	"io"

	"github.com/goccy/go-graphviz/cgraph"
	"github.com/goccy/go-graphviz/gvc"
)

type Graphviz struct {
	ctx    *gvc.Context
	name   string
	dir    *GraphDescriptor
	layout Layout
}

type Layout string

const (
	CIRCO     Layout = "circo"
	DOT       Layout = "dot"
	FDP       Layout = "fdp"
	NEATO     Layout = "neato"
	NOP       Layout = "nop"
	NOP1      Layout = "nop1"
	NOP2      Layout = "nop2"
	OSAGE     Layout = "osage"
	PATCHWORK Layout = "patchwork"
	SFDP      Layout = "sfdp"
	TWOPI     Layout = "twopi"
)

type Format string

const (
	XDOT Format = "dot"
	SVG  Format = "svg"
	PNG  Format = "png"
	JPG  Format = "jpg"
)

func New(ctx context.Context) (*Graphviz, error) {
	c, err := gvc.New(ctx)
	if err != nil {
		return nil, err
	}
	return &Graphviz{
		ctx:    c,
		dir:    Directed,
		layout: DOT,
	}, nil
}

func NewWithPlugins(ctx context.Context, plugins ...Plugin) (*Graphviz, error) {
	c, err := gvc.NewWithPlugins(ctx, plugins...)
	if err != nil {
		return nil, err
	}
	return &Graphviz{
		ctx:    c,
		dir:    Directed,
		layout: DOT,
	}, nil
}

func (g *Graphviz) Close() error {
	return g.ctx.Close()
}

func (g *Graphviz) SetLayout(layout Layout) *Graphviz {
	g.layout = layout
	return g
}

func (g *Graphviz) Render(ctx context.Context, graph *Graph, format Format, w io.Writer) (e error) {
	defer func() {
		if err := g.ctx.FreeLayout(ctx, graph); err != nil {
			e = err
		}
	}()

	if err := g.ctx.Layout(ctx, graph, string(g.layout)); err != nil {
		return err
	}
	if err := g.ctx.RenderData(ctx, graph, string(format), w); err != nil {
		return err
	}
	return nil
}

func (g *Graphviz) RenderImage(ctx context.Context, graph *Graph) (img image.Image, e error) {
	defer func() {
		if err := g.ctx.FreeLayout(ctx, graph); err != nil {
			e = err
		}
	}()

	if err := g.ctx.Layout(ctx, graph, string(g.layout)); err != nil {
		return nil, err
	}
	image, err := g.ctx.RenderImage(ctx, graph, string(PNG))
	if err != nil {
		return nil, err
	}
	return image, nil
}

func (g *Graphviz) RenderFilename(ctx context.Context, graph *Graph, format Format, path string) (e error) {
	defer func() {
		if err := g.ctx.FreeLayout(ctx, graph); err != nil {
			e = err
		}
	}()

	if err := g.ctx.Layout(ctx, graph, string(g.layout)); err != nil {
		return err
	}
	if err := g.ctx.RenderFilename(ctx, graph, string(format), path); err != nil {
		return err
	}
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
	return graph, nil
}
