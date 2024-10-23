package gvc

import (
	"github.com/goccy/go-graphviz/internal/wasm"
)

func init() {
	getRenderEnginePtr := func(job *wasm.Job) uint64 {
		return job.GetGvc().GetApi()[wasm.API_RENDER].GetTypeptr().GetEngine().(uint64)
	}
	getLoadImageEnginePtr := func(job *wasm.Job) uint64 {
		return job.GetGvc().GetApi()[wasm.API_LOADIMAGE].GetTypeptr().GetEngine().(uint64)
	}

	wasm.Register_RenderEngine_BeginJob(func(job *wasm.Job) (uint64, error) { return getRenderEnginePtr(job), nil })
	wasm.Register_RenderEngine_EndJob(func(job *wasm.Job) (uint64, error) { return getRenderEnginePtr(job), nil })
	wasm.Register_RenderEngine_BeginGraph(func(job *wasm.Job) (uint64, error) { return getRenderEnginePtr(job), nil })
	wasm.Register_RenderEngine_EndGraph(func(job *wasm.Job) (uint64, error) { return getRenderEnginePtr(job), nil })
	wasm.Register_RenderEngine_BeginLayer(func(job *wasm.Job, _ string, _ int, _ int) (uint64, error) { return getRenderEnginePtr(job), nil })
	wasm.Register_RenderEngine_EndLayer(func(job *wasm.Job) (uint64, error) { return getRenderEnginePtr(job), nil })
	wasm.Register_RenderEngine_BeginPage(func(job *wasm.Job) (uint64, error) { return getRenderEnginePtr(job), nil })
	wasm.Register_RenderEngine_EndPage(func(job *wasm.Job) (uint64, error) { return getRenderEnginePtr(job), nil })
	wasm.Register_RenderEngine_BeginCluster(func(job *wasm.Job) (uint64, error) { return getRenderEnginePtr(job), nil })
	wasm.Register_RenderEngine_EndCluster(func(job *wasm.Job) (uint64, error) { return getRenderEnginePtr(job), nil })
	wasm.Register_RenderEngine_BeginNodes(func(job *wasm.Job) (uint64, error) { return getRenderEnginePtr(job), nil })
	wasm.Register_RenderEngine_EndNodes(func(job *wasm.Job) (uint64, error) { return getRenderEnginePtr(job), nil })
	wasm.Register_RenderEngine_BeginEdges(func(job *wasm.Job) (uint64, error) { return getRenderEnginePtr(job), nil })
	wasm.Register_RenderEngine_EndEdges(func(job *wasm.Job) (uint64, error) { return getRenderEnginePtr(job), nil })
	wasm.Register_RenderEngine_BeginNode(func(job *wasm.Job) (uint64, error) { return getRenderEnginePtr(job), nil })
	wasm.Register_RenderEngine_EndNode(func(job *wasm.Job) (uint64, error) { return getRenderEnginePtr(job), nil })
	wasm.Register_RenderEngine_BeginEdge(func(job *wasm.Job) (uint64, error) { return getRenderEnginePtr(job), nil })
	wasm.Register_RenderEngine_EndEdge(func(job *wasm.Job) (uint64, error) { return getRenderEnginePtr(job), nil })
	wasm.Register_RenderEngine_BeginAnchor(func(job *wasm.Job, _ string, _ string, _ string, _ string) (uint64, error) {
		return getRenderEnginePtr(job), nil
	})
	wasm.Register_RenderEngine_EndAnchor(func(job *wasm.Job) (uint64, error) { return getRenderEnginePtr(job), nil })
	wasm.Register_RenderEngine_BeginLabel(func(job *wasm.Job, _ wasm.LabelType) (uint64, error) { return getRenderEnginePtr(job), nil })
	wasm.Register_RenderEngine_EndLabel(func(job *wasm.Job) (uint64, error) { return getRenderEnginePtr(job), nil })
	wasm.Register_RenderEngine_Textspan(func(job *wasm.Job, _ *wasm.PointFloat, _ *wasm.Textspan) (uint64, error) {
		return getRenderEnginePtr(job), nil
	})
	wasm.Register_RenderEngine_ResolveColor(func(job *wasm.Job, _ *wasm.Color) (uint64, error) { return getRenderEnginePtr(job), nil })
	wasm.Register_RenderEngine_Ellipse(func(job *wasm.Job, _ []*wasm.PointFloat, _ int) (uint64, error) { return getRenderEnginePtr(job), nil })
	wasm.Register_RenderEngine_Polygon(func(job *wasm.Job, _ []*wasm.PointFloat, _ uint32, _ int) (uint64, error) {
		return getRenderEnginePtr(job), nil
	})
	wasm.Register_RenderEngine_Beziercurve(func(job *wasm.Job, _ []*wasm.PointFloat, _ uint32, _ int) (uint64, error) {
		return getRenderEnginePtr(job), nil
	})
	wasm.Register_RenderEngine_Polyline(func(job *wasm.Job, _ []*wasm.PointFloat, _ uint32) (uint64, error) {
		return getRenderEnginePtr(job), nil
	})
	wasm.Register_RenderEngine_Comment(func(job *wasm.Job, _ string) (uint64, error) { return getRenderEnginePtr(job), nil })
	wasm.Register_RenderEngine_LibraryShape(func(job *wasm.Job, _ string, _ []*wasm.PointFloat, _ uint32, _ int) (uint64, error) {
		return getRenderEnginePtr(job), nil
	})

	wasm.Register_LoadImageEngine_LoadImage(func(job *wasm.Job, shape *wasm.UserShape, bf *wasm.BoxFloat, filled bool) (uint64, error) {
		return getLoadImageEnginePtr(job), nil
	})
}
