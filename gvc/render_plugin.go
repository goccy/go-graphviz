package gvc

import (
	"context"

	"github.com/goccy/go-graphviz/cdt"
	"github.com/goccy/go-graphviz/internal/wasm"
)

type RenderPlugin struct {
	plugin *wasm.PluginAPI
	engine RenderEngine
}

func (p *RenderPlugin) raw() *wasm.PluginAPI {
	return p.plugin
}

func (p *RenderPlugin) RenderEngine() RenderEngine {
	return p.engine
}

type RenderEngine interface {
	BeginJob(ctx context.Context, job *Job) error
	EndJob(ctx context.Context, job *Job) error
	BeginGraph(ctx context.Context, job *Job) error
	EndGraph(ctx context.Context, job *Job) error
	BeginLayer(ctx context.Context, job *Job, layerName string, layerNum, numLayers int) error
	EndLayer(ctx context.Context, job *Job) error
	BeginPage(ctx context.Context, job *Job) error
	EndPage(ctx context.Context, job *Job) error
	BeginCluster(ctx context.Context, job *Job) error
	EndCluster(ctx context.Context, job *Job) error
	BeginNodes(ctx context.Context, job *Job) error
	EndNodes(ctx context.Context, job *Job) error
	BeginEdges(ctx context.Context, job *Job) error
	EndEdges(ctx context.Context, job *Job) error
	BeginNode(ctx context.Context, job *Job) error
	EndNode(ctx context.Context, job *Job) error
	BeginEdge(ctx context.Context, job *Job) error
	EndEdge(ctx context.Context, job *Job) error
	BeginAnchor(ctx context.Context, job *Job, href, tooltip, target, id string) error
	EndAnchor(ctx context.Context, job *Job) error
	BeginLabel(ctx context.Context, job *Job, labelType LabelType) error
	EndLabel(ctx context.Context, job *Job) error
	TextSpan(ctx context.Context, job *Job, point *PointFloat, textSpan *TextSpan) error
	ResolveColor(ctx context.Context, job *Job, color *Color) error
	Ellipse(ctx context.Context, job *Job, points []*PointFloat, filled bool) error
	Polygon(ctx context.Context, job *Job, points []*PointFloat, filled bool) error
	BezierCurve(ctx context.Context, job *Job, points []*PointFloat, filled bool) error
	Polyline(ctx context.Context, job *Job, points []*PointFloat) error
	Comment(ctx context.Context, job *Job, comment string) error
	LibraryShape(ctx context.Context, job *Job, s string, points []*PointFloat, filled bool) error
	LoadImage(ctx context.Context, job *Job, shape *UserShape, box *BoxFloat, filled bool) error
}

type DefaultRenderEngine struct {
}

func (e *DefaultRenderEngine) BeginJob(_ context.Context, _ *Job) error {
	return nil
}

func (e *DefaultRenderEngine) EndJob(_ context.Context, _ *Job) error {
	return nil
}

func (e *DefaultRenderEngine) BeginGraph(_ context.Context, _ *Job) error {
	return nil
}

func (e *DefaultRenderEngine) EndGraph(_ context.Context, _ *Job) error {
	return nil
}

func (e *DefaultRenderEngine) BeginLayer(_ context.Context, _ *Job, layerName string, layerNum, numLayers int) error {
	return nil
}

func (e *DefaultRenderEngine) EndLayer(_ context.Context, _ *Job) error {
	return nil
}

func (e *DefaultRenderEngine) BeginPage(_ context.Context, _ *Job) error {
	return nil
}

func (e *DefaultRenderEngine) EndPage(_ context.Context, _ *Job) error {
	return nil
}

func (e *DefaultRenderEngine) BeginCluster(_ context.Context, _ *Job) error {
	return nil
}

func (e *DefaultRenderEngine) EndCluster(_ context.Context, _ *Job) error {
	return nil
}

func (e *DefaultRenderEngine) BeginNodes(_ context.Context, _ *Job) error {
	return nil
}

func (e *DefaultRenderEngine) EndNodes(_ context.Context, _ *Job) error {
	return nil
}

func (e *DefaultRenderEngine) BeginEdges(_ context.Context, _ *Job) error {
	return nil
}

func (e *DefaultRenderEngine) EndEdges(_ context.Context, _ *Job) error {
	return nil
}

func (e *DefaultRenderEngine) BeginNode(_ context.Context, _ *Job) error {
	return nil
}

func (e *DefaultRenderEngine) EndNode(_ context.Context, _ *Job) error {
	return nil
}

func (e *DefaultRenderEngine) BeginEdge(_ context.Context, _ *Job) error {
	return nil
}

func (e *DefaultRenderEngine) EndEdge(_ context.Context, _ *Job) error {
	return nil
}

func (e *DefaultRenderEngine) BeginAnchor(_ context.Context, _ *Job, _, _, _, _ string) error {
	return nil
}

func (e *DefaultRenderEngine) EndAnchor(_ context.Context, _ *Job) error {
	return nil
}

func (e *DefaultRenderEngine) BeginLabel(_ context.Context, _ *Job, _ LabelType) error {
	return nil
}

func (e *DefaultRenderEngine) EndLabel(_ context.Context, _ *Job) error {
	return nil
}

func (e *DefaultRenderEngine) TextSpan(_ context.Context, _ *Job, _ *PointFloat, _ *TextSpan) error {
	return nil
}

func (e *DefaultRenderEngine) ResolveColor(_ context.Context, _ *Job, _ *Color) error {
	return nil
}

func (e *DefaultRenderEngine) Ellipse(_ context.Context, _ *Job, _ []*PointFloat, _ bool) error {
	return nil
}

func (e *DefaultRenderEngine) Polygon(_ context.Context, _ *Job, _ []*PointFloat, _ bool) error {
	return nil
}

func (e *DefaultRenderEngine) BezierCurve(_ context.Context, _ *Job, _ []*PointFloat, _ bool) error {
	return nil
}

func (e *DefaultRenderEngine) Polyline(_ context.Context, _ *Job, _ []*PointFloat) error {
	return nil
}

func (e *DefaultRenderEngine) Comment(_ context.Context, _ *Job, _ string) error {
	return nil
}

func (e *DefaultRenderEngine) LibraryShape(_ context.Context, _ *Job, _ string, _ []*PointFloat, _ bool) error {
	return nil
}

func (e *DefaultRenderEngine) LoadImage(_ context.Context, _ *Job, _ *UserShape, _ *BoxFloat, _ bool) error {
	return nil
}

type RenderFeature int64

var (
	RenderYGoesDown        RenderFeature = RenderFeature(wasm.RENDER_Y_GOES_DOWN)
	RenderDoesTransform    RenderFeature = RenderFeature(wasm.RENDER_DOES_TRANSFORM)
	RenderDoesLabels       RenderFeature = RenderFeature(wasm.RENDER_DOES_LABELS)
	RenderDoesMaps         RenderFeature = RenderFeature(wasm.RENDER_DOES_MAPS)
	RenderDoesMapRectangle RenderFeature = RenderFeature(wasm.RENDER_DOES_MAP_RECTANGLE)
	RenderDoesMapCircle    RenderFeature = RenderFeature(wasm.RENDER_DOES_MAP_CIRCLE)
	RenderDoesMapPolygon   RenderFeature = RenderFeature(wasm.RENDER_DOES_MAP_POLYGON)
	RenderDoesMapEllipse   RenderFeature = RenderFeature(wasm.RENDER_DOES_MAP_ELLIPSE)
	RenderDoesMapBspline   RenderFeature = RenderFeature(wasm.RENDER_DOES_MAP_BSPLINE)
	RenderDoesTooltips     RenderFeature = RenderFeature(wasm.RENDER_DOES_TOOLTIPS)
	RenderDoesTargets      RenderFeature = RenderFeature(wasm.RENDER_DOES_TARGETS)
	RenderDoesZ            RenderFeature = RenderFeature(wasm.RENDER_DOES_Z)
	RenderNoWhiteBg        RenderFeature = RenderFeature(wasm.RENDER_NO_WHITE_BG)
)

type ColorType int64

var (
	HSVADouble  ColorType = ColorType(wasm.HSVA_DOUBLE)
	RGBAByte    ColorType = ColorType(wasm.RGBA_BYTE)
	RGBAWord    ColorType = ColorType(wasm.RGBA_WORD)
	RGBADouble  ColorType = ColorType(wasm.RGBA_DOUBLE)
	ColorString ColorType = ColorType(wasm.COLOR_STRING)
	ColorIndex  ColorType = ColorType(wasm.COLOR_INDEX)
)

type RenderPluginOption func(*renderConfig)

func WithRenderQuality(quality int) RenderPluginOption {
	return func(cfg *renderConfig) {
		cfg.Quality = int64(quality)
	}
}

func WithRenderFeatures(features ...RenderFeature) RenderPluginOption {
	return func(cfg *renderConfig) {
		cfg.Features = features
	}
}

func WithRenderColorType(typ ColorType) RenderPluginOption {
	return func(cfg *renderConfig) {
		cfg.ColorType = typ
	}
}

func WithRenderPAD(pad float64) RenderPluginOption {
	return func(cfg *renderConfig) {
		cfg.PAD = pad
	}
}

func NewRenderPlugin(ctx context.Context, typ string, engine RenderEngine, opts ...RenderPluginOption) (*RenderPlugin, error) {
	cfg := defaultRenderPluginConfig(typ, engine)
	for _, opt := range opts {
		opt(cfg)
	}
	return newRenderPlugin(ctx, cfg)
}

func PNGRenderPlugin(ctx context.Context) (*RenderPlugin, error) {
	return newRenderPlugin(ctx, defaultRenderPluginConfig("png", newPNGRenderEngine()))
}

func JPGRenderPlugin(ctx context.Context) (*RenderPlugin, error) {
	return newRenderPlugin(ctx, defaultRenderPluginConfig("jpg", newJPGRenderEngine()))
}

func defaultRenderPluginConfig(typ string, engine RenderEngine) *renderConfig {
	return &renderConfig{
		Type:    typ,
		Quality: 1,
		Features: []RenderFeature{
			RenderYGoesDown,
			RenderDoesTransform,
		},
		ColorType:    RGBAByte,
		PAD:          4,
		RenderEngine: engine,
	}
}

func newPNGRenderEngine() *ImageRenderer {
	return &ImageRenderer{DefaultRenderEngine: new(DefaultRenderEngine)}
}

func newJPGRenderEngine() *ImageRenderer {
	return &ImageRenderer{DefaultRenderEngine: new(DefaultRenderEngine)}
}

type renderConfig struct {
	Type         string
	Quality      int64
	PAD          float64
	Features     []RenderFeature
	ColorType    ColorType
	RenderEngine RenderEngine
}

func newRenderPlugin(ctx context.Context, cfg *renderConfig) (*RenderPlugin, error) {
	plg, err := wasm.NewPluginAPI(ctx)
	if err != nil {
		return nil, err
	}
	if err := plg.SetApi(wasm.API_RENDER); err != nil {
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
	features, err := wasm.NewRenderFeatures(ctx)
	if err != nil {
		return nil, err
	}
	var flags int64
	for _, feature := range cfg.Features {
		flags |= int64(feature)
	}
	features.SetFlags(flags)
	features.SetDefaultPad(cfg.PAD)
	features.SetColorType(wasm.ColorType(cfg.ColorType))
	if err := types.SetFeatures(features); err != nil {
		return nil, err
	}
	engine, err := newRenderEngine(ctx, cfg.RenderEngine)
	if err != nil {
		return nil, err
	}
	if err := types.SetEngine(engine); err != nil {
		return nil, err
	}
	term, err := wasm.PluginInstalledZero(ctx)
	if err != nil {
		return nil, err
	}
	if err := plg.SetTypes([]*wasm.PluginInstalled{types, term}); err != nil {
		return nil, err
	}
	return &RenderPlugin{
		plugin: plg,
		engine: cfg.RenderEngine,
	}, nil
}

func newRenderEngine(ctx context.Context, engine RenderEngine) (*wasm.RenderEngine, error) {
	e, err := wasm.NewRenderEngine(ctx)
	if err != nil {
		return nil, err
	}
	ptr := wasm.WasmPtr(e)
	if err := e.SetBeginJob(ctx, wasm.CreateCallbackFunc(func(ctx context.Context, job *wasm.Job) error {
		return engine.BeginJob(ctx, toJob(job))
	}, ptr)); err != nil {
		return nil, err
	}
	if err := e.SetEndJob(ctx, wasm.CreateCallbackFunc(func(ctx context.Context, job *wasm.Job) error {
		return engine.EndJob(ctx, toJob(job))
	}, ptr)); err != nil {
		return nil, err
	}
	if err := e.SetBeginGraph(ctx, wasm.CreateCallbackFunc(func(ctx context.Context, job *wasm.Job) error {
		return engine.BeginGraph(ctx, toJob(job))
	}, ptr)); err != nil {
		return nil, err
	}
	if err := e.SetEndGraph(ctx, wasm.CreateCallbackFunc(func(ctx context.Context, job *wasm.Job) error {
		return engine.EndGraph(ctx, toJob(job))
	}, ptr)); err != nil {
		return nil, err
	}
	if err := e.SetBeginLayer(ctx, wasm.CreateCallbackFunc(func(ctx context.Context, job *wasm.Job, layerName string, layerNum int, numLayers int) error {
		return engine.BeginLayer(ctx, toJob(job), layerName, layerNum, numLayers)
	}, ptr)); err != nil {
		return nil, err
	}
	if err := e.SetEndLayer(ctx, wasm.CreateCallbackFunc(func(ctx context.Context, job *wasm.Job) error {
		return engine.EndLayer(ctx, toJob(job))
	}, ptr)); err != nil {
		return nil, err
	}
	if err := e.SetBeginPage(ctx, wasm.CreateCallbackFunc(func(ctx context.Context, job *wasm.Job) error {
		return engine.BeginPage(ctx, toJob(job))
	}, ptr)); err != nil {
		return nil, err
	}
	if err := e.SetEndPage(ctx, wasm.CreateCallbackFunc(func(ctx context.Context, job *wasm.Job) error {
		return engine.EndPage(ctx, toJob(job))
	}, ptr)); err != nil {
		return nil, err
	}
	if err := e.SetBeginCluster(ctx, wasm.CreateCallbackFunc(func(ctx context.Context, job *wasm.Job) error {
		return engine.BeginCluster(ctx, toJob(job))
	}, ptr)); err != nil {
		return nil, err
	}
	if err := e.SetEndCluster(ctx, wasm.CreateCallbackFunc(func(ctx context.Context, job *wasm.Job) error {
		return engine.EndCluster(ctx, toJob(job))
	}, ptr)); err != nil {
		return nil, err
	}
	if err := e.SetBeginNodes(ctx, wasm.CreateCallbackFunc(func(ctx context.Context, job *wasm.Job) error {
		return engine.BeginNodes(ctx, toJob(job))
	}, ptr)); err != nil {
		return nil, err
	}
	if err := e.SetEndNodes(ctx, wasm.CreateCallbackFunc(func(ctx context.Context, job *wasm.Job) error {
		return engine.EndNodes(ctx, toJob(job))
	}, ptr)); err != nil {
		return nil, err
	}
	if err := e.SetBeginEdges(ctx, wasm.CreateCallbackFunc(func(ctx context.Context, job *wasm.Job) error {
		return engine.BeginEdges(ctx, toJob(job))
	}, ptr)); err != nil {
		return nil, err
	}
	if err := e.SetEndEdges(ctx, wasm.CreateCallbackFunc(func(ctx context.Context, job *wasm.Job) error {
		return engine.EndEdges(ctx, toJob(job))
	}, ptr)); err != nil {
		return nil, err
	}
	if err := e.SetBeginNode(ctx, wasm.CreateCallbackFunc(func(ctx context.Context, job *wasm.Job) error {
		return engine.BeginNode(ctx, toJob(job))
	}, ptr)); err != nil {
		return nil, err
	}
	if err := e.SetEndNode(ctx, wasm.CreateCallbackFunc(func(ctx context.Context, job *wasm.Job) error {
		return engine.EndNode(ctx, toJob(job))
	}, ptr)); err != nil {
		return nil, err
	}
	if err := e.SetBeginEdge(ctx, wasm.CreateCallbackFunc(func(ctx context.Context, job *wasm.Job) error {
		return engine.BeginEdge(ctx, toJob(job))
	}, ptr)); err != nil {
		return nil, err
	}
	if err := e.SetEndEdge(ctx, wasm.CreateCallbackFunc(func(ctx context.Context, job *wasm.Job) error {
		return engine.EndEdge(ctx, toJob(job))
	}, ptr)); err != nil {
		return nil, err
	}
	if err := e.SetBeginAnchor(ctx, wasm.CreateCallbackFunc(func(ctx context.Context, job *wasm.Job, href string, tooltip string, target string, id string) error {
		return engine.BeginAnchor(ctx, toJob(job), href, tooltip, target, id)
	}, ptr)); err != nil {
		return nil, err
	}
	if err := e.SetEndAnchor(ctx, wasm.CreateCallbackFunc(func(ctx context.Context, job *wasm.Job) error {
		return engine.EndAnchor(ctx, toJob(job))
	}, ptr)); err != nil {
		return nil, err
	}
	if err := e.SetBeginLabel(ctx, wasm.CreateCallbackFunc(func(ctx context.Context, job *wasm.Job, typ wasm.LabelType) error {
		return engine.BeginLabel(ctx, toJob(job), LabelType(typ))
	}, ptr)); err != nil {
		return nil, err
	}
	if err := e.SetEndLabel(ctx, wasm.CreateCallbackFunc(func(ctx context.Context, job *wasm.Job) error {
		return engine.EndLabel(ctx, toJob(job))
	}, ptr)); err != nil {
		return nil, err
	}
	if err := e.SetTextspan(ctx, wasm.CreateCallbackFunc(func(ctx context.Context, job *wasm.Job, p *wasm.PointFloat, span *wasm.Textspan) error {
		return engine.TextSpan(ctx, toJob(job), toPointFloat(p), toTextSpan(span))
	}, ptr)); err != nil {
		return nil, err
	}
	if err := e.SetResolveColor(ctx, wasm.CreateCallbackFunc(func(ctx context.Context, job *wasm.Job, c *wasm.Color) error {
		return engine.ResolveColor(ctx, toJob(job), toColor(c))
	}, ptr)); err != nil {
		return nil, err
	}
	if err := e.SetEllipse(ctx, wasm.CreateCallbackFunc(func(ctx context.Context, job *wasm.Job, p []*wasm.PointFloat, filled int) error {
		points := make([]*PointFloat, len(p))
		for i := range p {
			points[i] = toPointFloat(p[i])
		}
		return engine.Ellipse(ctx, toJob(job), points, filled > 0)
	}, ptr)); err != nil {
		return nil, err
	}
	if err := e.SetPolygon(ctx, wasm.CreateCallbackFunc(func(ctx context.Context, job *wasm.Job, p []*wasm.PointFloat, _ uint32, filled int) error {
		points := make([]*PointFloat, len(p))
		for i := range p {
			points[i] = toPointFloat(p[i])
		}
		return engine.Polygon(ctx, toJob(job), points, filled > 0)
	}, ptr)); err != nil {
		return nil, err
	}
	if err := e.SetBeziercurve(ctx, wasm.CreateCallbackFunc(func(ctx context.Context, job *wasm.Job, p []*wasm.PointFloat, _ uint32, filled int) error {
		points := make([]*PointFloat, len(p))
		for i := range p {
			points[i] = toPointFloat(p[i])
		}
		return engine.BezierCurve(ctx, toJob(job), points, filled > 0)
	}, ptr)); err != nil {
		return nil, err
	}
	if err := e.SetPolyline(ctx, wasm.CreateCallbackFunc(func(ctx context.Context, job *wasm.Job, p []*wasm.PointFloat, _ uint32) error {
		points := make([]*PointFloat, len(p))
		for i := range p {
			points[i] = toPointFloat(p[i])
		}
		return engine.Polyline(ctx, toJob(job), points)
	}, ptr)); err != nil {
		return nil, err
	}
	if err := e.SetComment(ctx, wasm.CreateCallbackFunc(func(ctx context.Context, job *wasm.Job, comment string) error {
		return engine.Comment(ctx, toJob(job), comment)
	}, ptr)); err != nil {
		return nil, err
	}
	if err := e.SetLibraryShape(ctx, wasm.CreateCallbackFunc(func(ctx context.Context, job *wasm.Job, s string, p []*wasm.PointFloat, _ uint32, filled int) error {
		points := make([]*PointFloat, len(p))
		for i := range p {
			points[i] = toPointFloat(p[i])
		}
		return engine.LibraryShape(ctx, toJob(job), s, points, filled > 0)
	}, ptr)); err != nil {
		return nil, err
	}
	return e, nil
}

type LabelType int

var (
	LabelPlain LabelType = LabelType(wasm.LABEL_PLAIN)
	LabelHTML  LabelType = LabelType(wasm.LABEL_HTML)
)

type Job struct {
	wasm *wasm.Job
}

func toJob(v *wasm.Job) *Job {
	if v == nil {
		return nil
	}
	return &Job{wasm: v}
}

func (j *Job) getWasm() *wasm.Job {
	if j == nil {
		return nil
	}
	return j.wasm
}

type Point struct {
	wasm *wasm.Point
}

func toPoint(v *wasm.Point) *Point {
	if v == nil {
		return nil
	}
	return &Point{wasm: v}
}

func (p *Point) getWasm() *wasm.Point {
	if p == nil {
		return nil
	}
	return p.wasm
}

func (p *Point) X() int {
	return int(p.wasm.GetX())
}

func (p *Point) SetX(x int) {
	p.wasm.SetX(int64(x))
}

func (p *Point) Y() int {
	return int(p.wasm.GetY())
}

func (p *Point) SetY(y int) {
	p.wasm.SetY(int64(y))
}

type PointFloat struct {
	wasm *wasm.PointFloat
}

func toPointFloat(v *wasm.PointFloat) *PointFloat {
	if v == nil {
		return nil
	}
	return &PointFloat{wasm: v}
}

func (p *PointFloat) getWasm() *wasm.PointFloat {
	if p == nil {
		return nil
	}
	return p.wasm
}

func (p *PointFloat) X() float64 {
	return p.wasm.GetX()
}

func (p *PointFloat) SetX(x float64) {
	p.wasm.SetX(x)
}

func (p *PointFloat) Y() float64 {
	return p.wasm.GetY()
}

func (p *PointFloat) SetY(y float64) {
	p.wasm.SetY(y)
}

type TextSpan struct {
	wasm *wasm.Textspan
}

func toTextSpan(v *wasm.Textspan) *TextSpan {
	if v == nil {
		return nil
	}
	return &TextSpan{wasm: v}
}

func (s *TextSpan) getWasm() *wasm.Textspan {
	if s == nil {
		return nil
	}
	return s.wasm
}

func (s *TextSpan) Text() string {
	return s.wasm.GetStr()
}

func (s *TextSpan) SetText(v string) {
	s.wasm.SetStr(v)
}

func (s *TextSpan) Font() *TextFont {
	return toTextFont(s.wasm.GetFont())
}

func (s *TextSpan) SetFont(v *TextFont) {
	s.wasm.SetFont(v.getWasm())
}

func (s *TextSpan) YOffsetLayout() float64 {
	return s.wasm.GetYOffsetLayout()
}

func (s *TextSpan) SetYOffsetLayout(v float64) {
	s.wasm.SetYOffsetLayout(v)
}

func (s *TextSpan) YOffsetCenterLine() float64 {
	return s.wasm.GetYOffsetCenterLine()
}

func (s *TextSpan) SetYOffsetCenterLine(v float64) {
	s.wasm.SetYOffsetCenterLine(v)
}

func (s *TextSpan) Size() *PointFloat {
	return toPointFloat(s.wasm.GetSize())
}

func (s *TextSpan) SetSize(v *PointFloat) {
	s.wasm.SetSize(v.getWasm())
}

func (s *TextSpan) Just() int {
	return int(s.wasm.GetJust())
}

func (s *TextSpan) SetJust(v int) {
	s.wasm.SetJust(int64(v))
}

type TextFont struct {
	wasm *wasm.TextFont
}

func toTextFont(v *wasm.TextFont) *TextFont {
	if v == nil {
		return nil
	}
	return &TextFont{wasm: v}
}

func (f *TextFont) getWasm() *wasm.TextFont {
	if f == nil {
		return nil
	}
	return f.wasm
}

func (f *TextFont) Name() string {
	return f.wasm.GetName()
}

func (f *TextFont) SetName(v string) {
	f.wasm.SetName(v)
}

func (f *TextFont) Color() string {
	return f.wasm.GetColor()
}

func (f *TextFont) SetColor(v string) {
	f.wasm.SetColor(v)
}

func (f *TextFont) PostScriptAlias() *PostScriptAlias {
	return toPostScriptAlias(f.wasm.GetPostscriptAlias())
}

func (f *TextFont) SetPostScriptAlias(v *PostScriptAlias) {
	f.wasm.SetPostscriptAlias(v.getWasm())
}

func (f *TextFont) Size() float64 {
	return f.wasm.GetSize()
}

func (f *TextFont) SetSize(v float64) {
	f.wasm.SetSize(v)
}

func (f *TextFont) Flags() uint {
	return uint(f.wasm.GetFlags())
}

func (f *TextFont) SetFlags(v uint) {
	f.wasm.SetFlags(uint64(v))
}

func (f *TextFont) Count() uint {
	return uint(f.wasm.GetCount())
}

func (f *TextFont) SetCount(v uint) {
	f.wasm.SetCount(uint64(v))
}

type PostScriptAlias struct {
	wasm *wasm.PostscriptAlias
}

func toPostScriptAlias(v *wasm.PostscriptAlias) *PostScriptAlias {
	if v == nil {
		return nil
	}
	return &PostScriptAlias{wasm: v}
}

func (a *PostScriptAlias) getWasm() *wasm.PostscriptAlias {
	if a == nil {
		return nil
	}
	return a.wasm
}

func (s *PostScriptAlias) Name() string {
	return s.wasm.GetName()
}

func (s *PostScriptAlias) SetName(v string) {
	s.wasm.SetName(v)
}

func (s *PostScriptAlias) Family() string {
	return s.wasm.GetFamily()
}

func (s *PostScriptAlias) SetFamily(v string) {
	s.wasm.SetFamily(v)
}

func (s *PostScriptAlias) Weight() string {
	return s.wasm.GetWeight()
}

func (s *PostScriptAlias) SetWeight(v string) {
	s.wasm.SetWeight(v)
}

func (s *PostScriptAlias) Stretch() string {
	return s.wasm.GetStretch()
}

func (s *PostScriptAlias) SetStretch(v string) {
	s.wasm.SetStretch(v)
}

func (s *PostScriptAlias) Style() string {
	return s.wasm.GetStyle()
}

func (s *PostScriptAlias) SetStyle(v string) {
	s.wasm.SetStyle(v)
}

func (s *PostScriptAlias) XFigCode() int {
	return int(s.wasm.GetXfigCode())
}

func (s *PostScriptAlias) SetXFigCode(v int) {
	s.wasm.SetXfigCode(int64(v))
}

func (s *PostScriptAlias) SVGFontFamily() string {
	return s.wasm.GetSvgFontFamily()
}

func (s *PostScriptAlias) SetSVGFontFamily(v string) {
	s.wasm.SetSvgFontFamily(v)
}

func (s *PostScriptAlias) SVGFontWeight() string {
	return s.wasm.GetSvgFontWeight()
}

func (s *PostScriptAlias) SetSVGFontWeight(v string) {
	s.wasm.SetSvgFontWeight(v)
}

func (s *PostScriptAlias) SVGFontStyle() string {
	return s.wasm.GetSvgFontStyle()
}

func (s *PostScriptAlias) SetSVGFontStyle(v string) {
	s.wasm.SetSvgFontStyle(v)
}

type Scale = PointFloat

func (j *Job) Zoom() float64 {
	return j.wasm.GetZoom()
}

func (j *Job) SetZoom(v float64) {
	j.wasm.SetZoom(v)
}

func (j *Job) Scale() *Scale {
	return toPointFloat(j.wasm.GetScale())
}

func (j *Job) SetScale(v *Scale) {
	j.wasm.SetScale(v.getWasm())
}

func (j *Job) Width() uint64 {
	return j.wasm.GetWidth()
}

func (j *Job) SetWidth(v uint64) {
	j.wasm.SetWidth(v)
}

func (j *Job) Height() uint64 {
	return j.wasm.GetHeight()
}

func (j *Job) SetHeight(v uint64) {
	j.wasm.SetHeight(v)
}

type Translation = PointFloat

func (j *Job) Translation() *Translation {
	return toPointFloat(j.wasm.GetTranslation())
}

func (j *Job) SetTranslation(v *Translation) {
	j.wasm.SetTranslation(v.getWasm())
}

func (j *Job) OutputData() []byte {
	return []byte(j.wasm.GetOutputData())
}

func (j *Job) SetOutputData(v []byte) {
	j.wasm.SetOutputData(string(v))
}

func (j *Job) ExternalContext() bool {
	return j.wasm.GetExternalContext()
}

func (j *Job) OutputLangName() string {
	return j.wasm.GetOutputLangname()
}

func (j *Job) SetOutputLangName(v string) {
	j.wasm.SetOutputLangname(v)
}

func (j *Job) OutputDataPosition() uint {
	return uint(j.wasm.GetOutputDataPosition())
}

func (j *Job) SetOutputDataPosition(v uint) {
	j.wasm.SetOutputDataPosition(v)
}

func (j *Job) OutputFileName() string {
	return j.wasm.GetOutputFilename()
}

func (j *Job) SetOutputFileName(v string) {
	j.wasm.SetOutputFilename(v)
}

func (j *Job) Object() *ObjectState {
	return toObjectState(j.wasm.GetObj())
}

func (j *Job) DPI() *PointFloat {
	return toPointFloat(j.wasm.GetDpi())
}

func (j *Job) SetDPI(v *PointFloat) {
	j.wasm.SetDpi(v.getWasm())
}

type ObjectState struct {
	wasm *wasm.ObjectState
}

func toObjectState(v *wasm.ObjectState) *ObjectState {
	if v == nil {
		return nil
	}
	return &ObjectState{wasm: v}
}

func (s *ObjectState) getWasm() *wasm.ObjectState {
	if s == nil {
		return nil
	}
	return s.wasm
}

type FillType int64

var (
	FillNone   FillType = FillType(wasm.FILL_NONE)
	FillSolid  FillType = FillType(wasm.FILL_SOLID)
	FillLinear FillType = FillType(wasm.FILL_LINEAR)
	FillRadial FillType = FillType(wasm.FILL_RADIAL)
)

type PenType int64

var (
	PenNone   PenType = PenType(wasm.PEN_NONE)
	PenDashed PenType = PenType(wasm.PEN_DASHED)
	PenDotted PenType = PenType(wasm.PEN_DOTTED)
	PenSolid  PenType = PenType(wasm.PEN_SOLID)
)

func (s *ObjectState) Pen() PenType {
	return PenType(s.wasm.GetPen())
}

func (s *ObjectState) SetPen(v PenType) {
	s.wasm.SetPen(wasm.PenType(v))
}

func (s *ObjectState) PenWidth() float64 {
	return s.wasm.GetPenwidth()
}

func (s *ObjectState) SetPenWidth(v float64) {
	s.wasm.SetPenwidth(v)
}

func (s *ObjectState) Fill() FillType {
	return FillType(s.wasm.GetFill())
}

func (s *ObjectState) SetFill(v FillType) {
	s.wasm.SetFill(wasm.FillType(v))
}

func (s *ObjectState) PenColor() *Color {
	return toColor(s.wasm.GetPencolor())
}

func (s *ObjectState) SetPenColor(v *Color) {
	s.wasm.SetPencolor(v.getWasm())
}

func (s *ObjectState) FillColor() *Color {
	return toColor(s.wasm.GetFillcolor())
}

func (s *ObjectState) SetFillColor(v *Color) {
	s.wasm.SetFillcolor(v.getWasm())
}

func (s *ObjectState) StopColor() *Color {
	return toColor(s.wasm.GetStopcolor())
}

func (s *ObjectState) RawStyle() []string {
	return s.wasm.GetRawstyle()
}

type Color struct {
	wasm *wasm.Color
}

func toColor(v *wasm.Color) *Color {
	if v == nil {
		return nil
	}
	return &Color{wasm: v}
}

func (c *Color) getWasm() *wasm.Color {
	if c == nil {
		return nil
	}
	return c.wasm
}

func (c *Color) RGBADouble() [4]float64 {
	res := c.wasm.GetRgbaDouble()
	return [4]float64{res[0], res[1], res[2], res[3]}
}

func (c *Color) SetRGBADouble(v [4]float64) {
	c.wasm.SetRgbaDouble(v[:])
}

func (c *Color) RGBAUint() [4]uint {
	res := c.wasm.GetRgbaUint()
	return [4]uint{res[0], res[1], res[2], res[3]}
}

func (c *Color) SetRGBAUint(v [4]uint) {
	c.wasm.SetRgbaUint(v[:])
}

func (c *Color) RGBAInt() [4]int {
	res := c.wasm.GetRgbaInt()
	return [4]int{res[0], res[1], res[2], res[3]}
}

func (c *Color) SetRGBAInt(v [4]int) {
	c.wasm.SetRgbaInt(v[:])
}

func (c *Color) HSVA() [4]float64 {
	res := c.wasm.GetHsva()
	return [4]float64{res[0], res[1], res[2], res[3]}
}

func (c *Color) SetHSVA(v [4]float64) {
	c.wasm.SetHsva(v[:])
}

func (c *Color) String() string {
	return c.wasm.GetString()
}

func (c *Color) SetString(v string) {
	c.wasm.SetString(v)
}

func (c *Color) Index() int {
	return int(c.wasm.GetIndex())
}

func (c *Color) SetIndex(v int) {
	c.wasm.SetIndex(int64(v))
}

func (c *Color) Type() ColorType {
	return ColorType(c.wasm.GetType())
}

func (c *Color) SetType(v ColorType) {
	c.wasm.SetType(wasm.ColorType(v))
}

type UserShape struct {
	wasm *wasm.UserShape
}

func toUserShape(v *wasm.UserShape) *UserShape {
	if v == nil {
		return nil
	}
	return &UserShape{wasm: v}
}

func (s *UserShape) getWasm() *wasm.UserShape {
	if s == nil {
		return nil
	}
	return s.wasm
}

func (s *UserShape) Link() *cdt.Link {
	return toDictLink(s.wasm.GetLink())
}

func (s *UserShape) SetLink(v *cdt.Link) {
	s.wasm.SetLink(toDictLinkWasm(v))
}

func (s *UserShape) Name() string {
	return s.wasm.GetName()
}

func (s *UserShape) SetName(v string) {
	s.wasm.SetName(v)
}

func (s *UserShape) MacroID() int {
	return int(s.wasm.GetMacroId())
}

func (s *UserShape) SetMacroID(v int) {
	s.wasm.SetMacroId(int64(v))
}

func (s *UserShape) MustInline() bool {
	return s.wasm.GetMustInline()
}

func (s *UserShape) SetMustInline(v bool) {
	s.wasm.SetMustInline(v)
}

func (s *UserShape) NoCache() bool {
	return s.wasm.GetNocache()
}

func (s *UserShape) SetNoCache(v bool) {
	s.wasm.SetNocache(v)
}

func (s *UserShape) ImageType() ImageType {
	return ImageType(s.wasm.GetType())
}

func (s *UserShape) SetImageType(v ImageType) {
	s.wasm.SetType(wasm.ImageType(v))
}

func (s *UserShape) StringType() string {
	return s.wasm.GetStringtype()
}

func (s *UserShape) SetStringType(v string) {
	s.wasm.SetStringtype(v)
}

func (s *UserShape) X() int {
	return int(s.wasm.GetX())
}

func (s *UserShape) SetX(v int) {
	s.wasm.SetX(int64(v))
}

func (s *UserShape) Y() int {
	return int(s.wasm.GetY())
}

func (s *UserShape) SetY(v int) {
	s.wasm.SetY(int64(v))
}

func (s *UserShape) Width() int {
	return int(s.wasm.GetW())
}

func (s *UserShape) SetWidth(v int) {
	s.wasm.SetW(int64(v))
}

func (s *UserShape) Height() int {
	return int(s.wasm.GetH())
}

func (s *UserShape) SetHeight(v int) {
	s.wasm.SetH(int64(v))
}

func (s *UserShape) DPI() int {
	return int(s.wasm.GetDpi())
}

func (s *UserShape) SetDPI(v int) {
	s.wasm.SetDpi(int64(v))
}

func (s *UserShape) Data() []byte {
	return s.wasm.GetData().([]byte)
}

func (s *UserShape) SetData(v []byte) {
	s.wasm.SetData(v)
}

func (s *UserShape) DataSize() uint {
	return uint(s.wasm.GetDatasize())
}

func (s *UserShape) SetDataSize(v uint) {
	s.wasm.SetDatasize(uint64(v))
}

type ImageType int

var (
	ImageTypeNull ImageType = ImageType(wasm.IMAGE_TYPE_NULL)
	ImageTypeBMP  ImageType = ImageType(wasm.IMAGE_TYPE_BMP)
	ImageTypeGIF  ImageType = ImageType(wasm.IMAGE_TYPE_GIF)
	ImageTypePNG  ImageType = ImageType(wasm.IMAGE_TYPE_PNG)
	ImageTypeJPEG ImageType = ImageType(wasm.IMAGE_TYPE_JPEG)
	ImageTypePDF  ImageType = ImageType(wasm.IMAGE_TYPE_PDF)
	ImageTypePS   ImageType = ImageType(wasm.IMAGE_TYPE_PS)
	ImageTypeEPS  ImageType = ImageType(wasm.IMAGE_TYPE_EPS)
	ImageTypeSVG  ImageType = ImageType(wasm.IMAGE_TYPE_SVG)
	ImageTypeXML  ImageType = ImageType(wasm.IMAGE_TYPE_XML)
	ImageTypeRIFF ImageType = ImageType(wasm.IMAGE_TYPE_RIFF)
	ImageTypeWEBP ImageType = ImageType(wasm.IMAGE_TYPE_WEBP)
	ImageTypeICO  ImageType = ImageType(wasm.IMAGE_TYPE_ICO)
	ImageTypeTIFF ImageType = ImageType(wasm.IMAGE_TYPE_TIFF)
)

type Box struct {
	wasm *wasm.Box
}

func toBox(v *wasm.Box) *Box {
	if v == nil {
		return nil
	}
	return &Box{wasm: v}
}

func (b *Box) getWasm() *wasm.Box {
	if b == nil {
		return nil
	}
	return b.wasm
}

func (b *Box) LL() *Point {
	return toPoint(b.wasm.GetLl())
}

func (b *Box) SetLL(v *Point) {
	b.wasm.SetLl(v.getWasm())
}

func (b *Box) UR() *Point {
	return toPoint(b.wasm.GetUr())
}

func (b *Box) SetUR(v *Point) {
	b.wasm.SetUr(v.getWasm())
}

type BoxFloat struct {
	wasm *wasm.BoxFloat
}

func toBoxFloat(v *wasm.BoxFloat) *BoxFloat {
	if v == nil {
		return nil
	}
	return &BoxFloat{wasm: v}
}

func (f *BoxFloat) getWasm() *wasm.BoxFloat {
	if f == nil {
		return nil
	}
	return f.wasm
}

func (f *BoxFloat) LL() *PointFloat {
	return toPointFloat(f.wasm.GetLl())
}

func (f *BoxFloat) SetLL(v *PointFloat) {
	f.wasm.SetLl(v.getWasm())
}

func (f *BoxFloat) UR() *PointFloat {
	return toPointFloat(f.wasm.GetUr())
}

func (f *BoxFloat) SetUR(v *PointFloat) {
	f.wasm.SetUr(v.getWasm())
}
