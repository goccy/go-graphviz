package gvc

import (
	"context"

	"github.com/goccy/go-graphviz/internal/wasm"
)

type DevicePlugin struct {
	plugin *wasm.PluginAPI
}

func (p *DevicePlugin) raw() *wasm.PluginAPI {
	return p.plugin
}

type DeviceFeature int64

var (
	DeviceDoesPages        DeviceFeature = DeviceFeature(wasm.DEVICE_DOES_PAGES)
	DeviceDoesLayers       DeviceFeature = DeviceFeature(wasm.DEVICE_DOES_LAYERS)
	DeviceEvents           DeviceFeature = DeviceFeature(wasm.DEVICE_EVENTS)
	DeviceDoesTrueColor    DeviceFeature = DeviceFeature(wasm.DEVICE_DOES_TRUECOLOR)
	DeviceBinaryFormat     DeviceFeature = DeviceFeature(wasm.DEVICE_BINARY_FORMAT)
	DeviceCompressedFormat DeviceFeature = DeviceFeature(wasm.DEVICE_COMPRESSED_FORMAT)
	DeviceNoWriter         DeviceFeature = DeviceFeature(wasm.DEVICE_NO_WRITER)
)

type DevicePluginOption func(*deviceConfig)

func WithDeviceQuality(quality int) DevicePluginOption {
	return func(cfg *deviceConfig) {
		cfg.Quality = int64(quality)
	}
}

func WithDeviceFeatures(features ...DeviceFeature) DevicePluginOption {
	return func(cfg *deviceConfig) {
		cfg.Features = features
	}
}

func WithDeviceDPI(x, y float64) DevicePluginOption {
	return func(cfg *deviceConfig) {
		cfg.DPI = deviceDPI{
			X: x,
			Y: y,
		}
	}
}

func NewDevicePlugin(ctx context.Context, typ string, opts ...DevicePluginOption) (*DevicePlugin, error) {
	cfg := defaultDevicePluginConfig(typ)
	for _, opt := range opts {
		opt(cfg)
	}
	return newDevicePlugin(ctx, cfg)
}

func PNGDevicePlugin(ctx context.Context) (*DevicePlugin, error) {
	return newDevicePlugin(ctx, defaultDevicePluginConfig("png:png"))
}

func JPGDevicePlugin(ctx context.Context) (*DevicePlugin, error) {
	return newDevicePlugin(ctx, defaultDevicePluginConfig("jpg:jpg"))
}

type deviceDPI struct {
	X float64
	Y float64
}

type deviceConfig struct {
	Type     string
	Quality  int64
	Features []DeviceFeature
	DPI      deviceDPI
}

func defaultDevicePluginConfig(typ string) *deviceConfig {
	return &deviceConfig{
		Type:    typ,
		Quality: 1,
		Features: []DeviceFeature{
			DeviceBinaryFormat,
			DeviceDoesTrueColor,
		},
		DPI: deviceDPI{
			X: 96,
			Y: 96,
		},
	}
}

func newDevicePlugin(ctx context.Context, cfg *deviceConfig) (*DevicePlugin, error) {
	plg, err := wasm.NewPluginAPI(ctx)
	if err != nil {
		return nil, err
	}
	if err := plg.SetApi(wasm.API_DEVICE); err != nil {
		return nil, err
	}
	types, err := wasm.NewPluginInstalled(ctx)
	if err != nil {
		return nil, err
	}
	if err := types.SetType(cfg.Type); err != nil {
		return nil, err
	}
	if err := types.SetQuality(cfg.Quality); err != nil {
		return nil, err
	}
	features, err := wasm.NewDeviceFeatures(ctx)
	if err != nil {
		return nil, err
	}
	var flags int64
	for _, feature := range cfg.Features {
		flags |= int64(feature)
	}
	features.SetFlags(flags)
	dpi, err := wasm.NewPointFloat(ctx)
	if err != nil {
		return nil, err
	}
	dpi.SetX(cfg.DPI.X)
	dpi.SetY(cfg.DPI.Y)
	features.SetDefaultDpi(dpi)
	if err := types.SetFeatures(features); err != nil {
		return nil, err
	}
	term, err := wasm.PluginInstalledZero(ctx)
	if err != nil {
		return nil, err
	}
	if err := plg.SetTypes([]*wasm.PluginInstalled{types, term}); err != nil {
		return nil, err
	}
	return &DevicePlugin{
		plugin: plg,
	}, nil
}
