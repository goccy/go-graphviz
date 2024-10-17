package gvc

import (
	"context"

	"github.com/goccy/go-graphviz/internal/wasm"
)

type Plugin interface {
	raw() *wasm.PluginAPI
}

func DefaultPlugins(ctx context.Context) ([]Plugin, error) {
	pngRenderPlugin, err := PNGRenderPlugin(ctx)
	if err != nil {
		return nil, err
	}
	pngDevicePlugin, err := PNGDevicePlugin(ctx)
	if err != nil {
		return nil, err
	}
	jpgRenderPlugin, err := JPGRenderPlugin(ctx)
	if err != nil {
		return nil, err
	}
	jpgDevicePlugin, err := JPGDevicePlugin(ctx)
	if err != nil {
		return nil, err
	}
	return []Plugin{
		pngRenderPlugin,
		pngDevicePlugin,
		jpgRenderPlugin,
		jpgDevicePlugin,
	}, nil
}
