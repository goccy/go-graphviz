package gvc

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"

	"github.com/goccy/go-graphviz/cgraph"
	"github.com/goccy/go-graphviz/internal/wasm"
)

type Context struct {
	gvc *wasm.Context
}

func New(ctx context.Context) (*Context, error) {
	plugins, err := DefaultPlugins(ctx)
	if err != nil {
		return nil, err
	}
	return NewWithPlugins(ctx, plugins...)
}

func NewWithPlugins(ctx context.Context, plugins ...Plugin) (*Context, error) {
	plgs, err := newPlugins(ctx, plugins...)
	if err != nil {
		return nil, err
	}
	gvc, err := wasm.GetContextWithPlugins(ctx, plgs, 1)
	if err != nil {
		return nil, err
	}
	if gvc == nil {
		return nil, fmt.Errorf("failed to create graphviz context")
	}
	return &Context{gvc: gvc}, nil
}

func (c *Context) Close() error {
	res, err := c.gvc.FreeContext(context.Background())
	if err != nil {
		return err
	}
	return toError(res)
}

func (c *Context) Layout(ctx context.Context, g *cgraph.Graph, engine string) error {
	if err := c.setupNodeLabelIfEmpty(g); err != nil {
		return err
	}
	res, err := c.gvc.Layout(ctx, toGraphWasm(g), engine)
	if err != nil {
		return err
	}
	return toError(res)
}

func (c *Context) RenderData(ctx context.Context, g *cgraph.Graph, format string, w io.Writer) error {
	var (
		s           string
		renderedLen uint
	)
	if _, err := c.gvc.RenderData(ctx, toGraphWasm(g), format, &s, &renderedLen); err != nil {
		return err
	}
	if _, err := w.Write([]byte(s)); err != nil {
		return err
	}
	return nil
}

func (c *Context) RenderImage(ctx context.Context, g *cgraph.Graph, format string) (image.Image, error) {
	var buf bytes.Buffer
	if err := c.RenderData(ctx, g, format, &buf); err != nil {
		return nil, err
	}
	img, _, err := image.Decode(&buf)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func (c *Context) RenderFilename(ctx context.Context, g *cgraph.Graph, format, filename string) error {
	res, err := c.gvc.RenderFilename(ctx, toGraphWasm(g), format, filename)
	if err != nil {
		return err
	}
	return toError(res)
}

func (c *Context) FreeLayout(ctx context.Context, g *cgraph.Graph) error {
	res, err := c.gvc.FreeLayout(ctx, toGraphWasm(g))
	if err != nil {
		return err
	}
	return toError(res)
}

func (c *Context) Clone(ctx context.Context) (*Context, error) {
	gvc, err := c.gvc.Clone(ctx)
	if err != nil {
		return nil, err
	}
	return &Context{gvc: gvc}, nil
}

func (c *Context) FreeClonedContext(ctx context.Context) error {
	return c.gvc.FreeClonedContext(ctx)
}

func (c *Context) setupNodeLabelIfEmpty(g *cgraph.Graph) error {
	n, err := g.FirstNode()
	if err != nil {
		return err
	}
	if err := c.setLabelIfEmpty(n); err != nil {
		return err
	}
	for {
		n, err = g.NextNode(n)
		if err != nil {
			return err
		}
		if n == nil {
			break
		}
		if err := c.setLabelIfEmpty(n); err != nil {
			return err
		}
	}
	return nil
}

func (c *Context) setLabelIfEmpty(n *cgraph.Node) error {
	label, err := n.Label()
	if err != nil {
		return err
	}
	if label == "" {
		n.SetLabel("\\N")
	}
	return nil
}

func newPlugins(ctx context.Context, plugins ...Plugin) ([]*wasm.SymList, error) {
	defaults, err := wasm.DefaultSymList(ctx)
	if err != nil {
		return nil, err
	}
	if len(plugins) == 0 {
		return defaults, nil
	}
	sym, err := wasm.NewSymList(ctx)
	if err != nil {
		return nil, err
	}
	if err := sym.SetName("gvplugin_go_LTX_library"); err != nil {
		return nil, err
	}
	lib, err := wasm.NewPluginLibrary(ctx)
	if err != nil {
		return nil, err
	}
	if err := lib.SetPackageName("go"); err != nil {
		return nil, err
	}
	var apis []*wasm.PluginAPI
	for _, plg := range plugins {
		apis = append(apis, plg.raw())
	}
	term, err := wasm.PluginAPIZero(ctx)
	if err != nil {
		return nil, err
	}
	apis = append(apis, term)

	if err := lib.SetApis(apis); err != nil {
		return nil, err
	}
	if err := sym.SetAddress(lib); err != nil {
		return nil, err
	}
	symTerm, err := wasm.SymListZero(ctx)
	if err != nil {
		return nil, err
	}
	return append(defaults, sym, symTerm), nil
}

func toError(result int) error {
	if result == 0 {
		return nil
	}
	return lastError()
}

func lastError() error {
	if e, _ := wasm.LastError(context.Background()); e != "" {
		return errors.New(e)
	}
	return nil
}
