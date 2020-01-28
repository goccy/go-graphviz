package gvc

import (
	"io"

	"github.com/goccy/go-graphviz/cgraph"
	"github.com/goccy/go-graphviz/internal/ccall"
	"golang.org/x/xerrors"
)

type Context struct {
	*ccall.GVC
}

type Job struct {
	*ccall.GVJ
}

func New() *Context {
	return &Context{GVC: ccall.GvContext()}
}

func (c *Context) Close() int {
	return ccall.GvFreeContext(c.GVC)
}

func (c *Context) Layout(g *cgraph.Graph, engine string) int {
	return ccall.GvLayout(c.GVC, g.Agraph, engine)
}

func (c *Context) RenderData(g *cgraph.Graph, format string, w io.Writer) error {
	if err := ccall.GvRenderData(c.GVC, g.Agraph, format, w); err != nil {
		return xerrors.Errorf("failed to GvRenderData: %w", err)
	}
	return nil
}

func (c *Context) RenderFilename(g *cgraph.Graph, format, filename string) int {
	return ccall.GvRenderFilename(c.GVC, g.Agraph, format, filename)
}

func (c *Context) FreeLayout(g *cgraph.Graph) int {
	return ccall.GvFreeLayout(c.GVC, g.Agraph)
}
