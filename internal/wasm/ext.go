package wasm

import (
	"context"
)

func DefaultSymList(ctx context.Context) ([]*SymList, error) {
	p, err := mod.NewPtr(ctx)
	if err != nil {
		return nil, err
	}
	if _, err := mod.ExportedFunction("wasm_bridge_SymList_default").Call(ctx, p); err != nil {
		return nil, err
	}
	ptr, err := mod.readU32(p)
	if err != nil {
		return nil, err
	}
	slice, err := mod.toSlice(ctx, ptr)
	if err != nil {
		return nil, err
	}
	return newSymListSlice(slice), nil
}

func PluginAPIZero(ctx context.Context) (*PluginAPI, error) {
	p, err := mod.NewPtr(ctx)
	if err != nil {
		return nil, err
	}
	if _, err := mod.ExportedFunction("wasm_bridge_PluginAPI_zero").Call(ctx, p); err != nil {
		return nil, err
	}
	ptr, err := mod.readU32(p)
	if err != nil {
		return nil, err
	}
	return newPluginAPI(ptr), nil
}

func PluginInstalledZero(ctx context.Context) (*PluginInstalled, error) {
	p, err := mod.NewPtr(ctx)
	if err != nil {
		return nil, err
	}
	if _, err := mod.ExportedFunction("wasm_bridge_PluginInstalled_zero").Call(ctx, p); err != nil {
		return nil, err
	}
	ptr, err := mod.readU32(p)
	if err != nil {
		return nil, err
	}
	return newPluginInstalled(ptr), nil
}

func SymListZero(ctx context.Context) (*SymList, error) {
	p, err := mod.NewPtr(ctx)
	if err != nil {
		return nil, err
	}
	if _, err := mod.ExportedFunction("wasm_bridge_SymList_zero").Call(ctx, p); err != nil {
		return nil, err
	}
	ptr, err := mod.readU32(p)
	if err != nil {
		return nil, err
	}
	return newSymList(ptr), nil
}

type RenderEngineInterface interface {
	BeginJob(context.Context, *Job) error
	EndJob(context.Context, *Job) error
	BeginGraph(context.Context, *Job) error
	EndGraph(context.Context, *Job) error
	BeginLayer(context.Context, *Job, string, int, int) error
	EndLayer(context.Context, *Job) error
	BeginPage(context.Context, *Job) error
	EndPage(context.Context, *Job) error
	BeginCluster(context.Context, *Job) error
	EndCluster(context.Context, *Job) error
	BeginNodes(context.Context, *Job) error
	EndNodes(context.Context, *Job) error
	BeginEdges(context.Context, *Job) error
	EndEdges(context.Context, *Job) error
	BeginNode(context.Context, *Job) error
	EndNode(context.Context, *Job) error
	BeginEdge(context.Context, *Job) error
	EndEdge(context.Context, *Job) error
	BeginAnchor(context.Context, *Job, string, string, string, string) error
	EndAnchor(context.Context, *Job) error
	BeginLabel(context.Context, *Job, LabelType) error
	EndLabel(context.Context, *Job) error
	Textspan(context.Context, *Job, *PointFloat, *Textspan) error
	ResolveColor(context.Context, *Job, *Color) error
	Ellipse(context.Context, *Job, []*PointFloat, int) error
	Polygon(context.Context, *Job, []*PointFloat, int) error
	Beziercurve(context.Context, *Job, []*PointFloat, int, int) error
	Polyline(context.Context, *Job, []*PointFloat) error
	Comment(context.Context, *Job, string) error
	LibraryShape(context.Context, *Job, string, *PointFloat, int, int) error
}
