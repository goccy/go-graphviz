package graphviz

import (
	"github.com/goccy/go-graphviz/cdt"
	"github.com/goccy/go-graphviz/cgraph"
	"github.com/goccy/go-graphviz/gvc"
)

// types from cdt package.
type (
	Dict       = cdt.Dict
	DictHold   = cdt.Hold
	DictLink   = cdt.Link
	DictMethod = cdt.Method
	DictData   = cdt.Data
	DictDisc   = cdt.Disc
	DictStat   = cdt.Stat
)

// types from cgraph package.
type (
	Graph            = cgraph.Graph
	Node             = cgraph.Node
	SubNode          = cgraph.SubNode
	Edge             = cgraph.Edge
	GraphDescriptor  = cgraph.Desc
	ClientDiscipline = cgraph.Disc
	Symbol           = cgraph.Symbol
	Record           = cgraph.Record
	Tag              = cgraph.Tag
	Object           = cgraph.Object
	CommonFields     = cgraph.CommonFields
	State            = cgraph.State
	CallbackStack    = cgraph.CallbackStack
	Attribute        = cgraph.Attr
	DataDict         = cgraph.DataDict
	ObjectTag        = cgraph.ObjectTag
	ID               = cgraph.ID
	ArrowType        = cgraph.ArrowType
	ClusterMode      = cgraph.ClusterMode
	DirType          = cgraph.DirType
	ImagePos         = cgraph.ImagePos
	JustType         = cgraph.JustType
	LabelLocation    = cgraph.LabelLocation
	ModeType         = cgraph.ModeType
	ModelType        = cgraph.ModelType
	OrderingType     = cgraph.OrderingType
	OutputMode       = cgraph.OutputMode
	PackMode         = cgraph.PackMode
	PageDir          = cgraph.PageDir
	QuadType         = cgraph.QuadType
	RankDir          = cgraph.RankDir
	RatioType        = cgraph.RatioType
	Shape            = cgraph.Shape
	SmoothType       = cgraph.SmoothType
	StartType        = cgraph.StartType
	GraphStyle       = cgraph.GraphStyle
	NodeStyle        = cgraph.NodeStyle
	EdgeStyle        = cgraph.EdgeStyle
)

// types from gvc package.
type (
	Plugin              = gvc.Plugin
	Context             = gvc.Context
	DevicePlugin        = gvc.DevicePlugin
	DeviceFeature       = gvc.DeviceFeature
	DevicePluginOption  = gvc.DevicePluginOption
	RenderPlugin        = gvc.RenderPlugin
	RenderEngine        = gvc.RenderEngine
	DefaultRenderEngine = gvc.DefaultRenderEngine
	RenderFeature       = gvc.RenderFeature
	RenderPluginOption  = gvc.RenderPluginOption
	ColorType           = gvc.ColorType
	LabelType           = gvc.LabelType
	Job                 = gvc.Job
	PointFloat          = gvc.PointFloat
	TextSpan            = gvc.TextSpan
	TextFont            = gvc.TextFont
	PostScriptAlias     = gvc.PostScriptAlias
	Scale               = gvc.Scale
	Translation         = gvc.Translation
	ObjectState         = gvc.ObjectState
	FillType            = gvc.FillType
	PenType             = gvc.PenType
	Color               = gvc.Color
)

// variables from cgraph package.
var (
	Directed         = cgraph.Directed
	StrictDirected   = cgraph.StrictDirected
	UnDirected       = cgraph.UnDirected
	StrictUnDirected = cgraph.StrictUnDirected
)

// const variables from cgraph package.
const (
	NormalArrow   = cgraph.NormalArrow
	InvArrow      = cgraph.InvArrow
	DotArrow      = cgraph.DotArrow
	InvDotArrow   = cgraph.InvDotArrow
	ODotArrow     = cgraph.ODotArrow
	InvODotArrow  = cgraph.InvODotArrow
	NoneArrow     = cgraph.NoneArrow
	TeeArrow      = cgraph.TeeArrow
	EmptyArrow    = cgraph.EmptyArrow
	InvEmptyArrow = cgraph.InvEmptyArrow
	DiamondArrow  = cgraph.DiamondArrow
	ODiamondArrow = cgraph.ODiamondArrow
	EDiamondArrow = cgraph.EDiamondArrow
	CrowArrow     = cgraph.CrowArrow
	BoxArrow      = cgraph.BoxArrow
	OBoxArrow     = cgraph.OBoxArrow
	OpenArrow     = cgraph.OpenArrow
	HalfOpenArrow = cgraph.HalfOpenArrow
	VeeArrow      = cgraph.VeeArrow
)

const (
	LocalCluster  = cgraph.LocalCluster
	GlobalCluster = cgraph.GlobalCluster
	NoneCluster   = cgraph.NoneCluster
)

const (
	ForwardDir = cgraph.ForwardDir
	BackDir    = cgraph.BackDir
	BothDir    = cgraph.BothDir
	NoneDir    = cgraph.NoneDir
)

const (
	TopLeftPos        = cgraph.TopLeftPos
	TopCenteredPos    = cgraph.TopCenteredPos
	TopRightPos       = cgraph.TopRightPos
	MiddleLeftPos     = cgraph.MiddleLeftPos
	MiddleCenteredPos = cgraph.MiddleCenteredPos
	BottomLeftPos     = cgraph.BottomLeftPos
	BottomCenteredPos = cgraph.BottomCenteredPos
	BottomRightPos    = cgraph.BottomRightPos
)

const (
	LeftJust     = cgraph.LeftJust
	CenteredJust = cgraph.CenteredJust
	RightJust    = cgraph.RightJust
)

const (
	TopLocation      = cgraph.TopLocation
	CenteredLocation = cgraph.CenteredLocation
	BottomLocation   = cgraph.BottomLocation
)

const (
	MajorMode  = cgraph.MajorMode
	KKMode     = cgraph.KKMode
	HierMode   = cgraph.HierMode
	IpsepMode  = cgraph.IpsepMode
	SpringMode = cgraph.SpringMode
	MaxentMode = cgraph.MaxentMode
)

const (
	ShortPathModel = cgraph.ShortPathModel
	CircuitModel   = cgraph.CircuitModel
	SubsetModel    = cgraph.SubsetModel
	MdsModel       = cgraph.MdsModel
)

const (
	OutOrdering = cgraph.OutOrdering
	InOrdering  = cgraph.InOrdering
)

const (
	BreadthFirst = cgraph.BreadthFirst
	NodesFirst   = cgraph.NodesFirst
	EdgesFirst   = cgraph.EdgesFirst
)

const (
	NodePack    = cgraph.NodePack
	ClusterPack = cgraph.ClusterPack
	GraphPack   = cgraph.GraphPack
)

const (
	BLDir = cgraph.BLDir
	BRDir = cgraph.BRDir
	TLDir = cgraph.TLDir
	TRDir = cgraph.TRDir
	RBDir = cgraph.RBDir
	RTDir = cgraph.RTDir
	LBDir = cgraph.LBDir
	LTDir = cgraph.LTDir
)

const (
	NormalQuad = cgraph.NormalQuad
	FastQuad   = cgraph.FastQuad
	NoneQuad   = cgraph.NoneQuad
)

const (
	TBRank = cgraph.TBRank
	LRRank = cgraph.LRRank
	BTRank = cgraph.BTRank
	RLRank = cgraph.RLRank
)

const (
	FillRatio     = cgraph.FillRatio
	CompressRatio = cgraph.CompressRatio
	ExpandRatio   = cgraph.ExpandRatio
	AutoRatio     = cgraph.AutoRatio
)

const (
	BoxShape             = cgraph.BoxShape
	PolygonShape         = cgraph.PolygonShape
	EllipseShape         = cgraph.EllipseShape
	OvalShape            = cgraph.OvalShape
	CircleShape          = cgraph.CircleShape
	PointShape           = cgraph.PointShape
	EggShape             = cgraph.EggShape
	TriangleShape        = cgraph.TriangleShape
	PlainTextShape       = cgraph.PlainTextShape
	PlainShape           = cgraph.PlainShape
	DiamondShape         = cgraph.DiamondShape
	TrapeziumShape       = cgraph.TrapeziumShape
	ParallelogramShape   = cgraph.ParallelogramShape
	HouseShape           = cgraph.HouseShape
	PentagonShape        = cgraph.PentagonShape
	HexagonShape         = cgraph.HexagonShape
	SeptagonShape        = cgraph.SeptagonShape
	OctagonShape         = cgraph.OctagonShape
	DoubleCircleShape    = cgraph.DoubleCircleShape
	DoubleOctagonShape   = cgraph.DoubleOctagonShape
	TripleOctagonShape   = cgraph.TripleOctagonShape
	InvTriangleShape     = cgraph.InvTriangleShape
	InvTrapeziumShape    = cgraph.InvTrapeziumShape
	InvHouseShape        = cgraph.InvHouseShape
	MdiamondShape        = cgraph.MdiamondShape
	MsquareShape         = cgraph.MsquareShape
	McircleShape         = cgraph.McircleShape
	RectShape            = cgraph.RectShape
	RectangleShape       = cgraph.RectangleShape
	SquareShape          = cgraph.SquareShape
	StarShape            = cgraph.StarShape
	NoneShape            = cgraph.NoneShape
	UnderlineShape       = cgraph.UnderlineShape
	CylinderShape        = cgraph.CylinderShape
	NoteShape            = cgraph.NoteShape
	TabShape             = cgraph.TabShape
	FolderShape          = cgraph.FolderShape
	Box3DShape           = cgraph.Box3DShape
	ComponentShape       = cgraph.ComponentShape
	PromoterShape        = cgraph.PromoterShape
	CdsShape             = cgraph.CdsShape
	TerminatorShape      = cgraph.TerminatorShape
	UtrShape             = cgraph.UtrShape
	PrimersiteShape      = cgraph.PrimersiteShape
	RestrictionSiteShape = cgraph.RestrictionSiteShape
	FivePoverHangShape   = cgraph.FivePoverHangShape
	ThreePoverHangShape  = cgraph.ThreePoverHangShape
	NoverHangShape       = cgraph.NoverHangShape
	AssemblyShape        = cgraph.AssemblyShape
	SignatureShape       = cgraph.SignatureShape
	InsulatorShape       = cgraph.InsulatorShape
	RibositeShape        = cgraph.RibositeShape
	RnastabShape         = cgraph.RnastabShape
	ProteasesiteShape    = cgraph.ProteasesiteShape
	ProteinstabShape     = cgraph.ProteinstabShape
	RPromoterShape       = cgraph.RPromoterShape
	RArrowShape          = cgraph.RArrowShape
	LArrowShape          = cgraph.LArrowShape
	LPromoterShape       = cgraph.LPromoterShape
)

const (
	NoneSmooth      = cgraph.NoneSmooth
	AvgDistSmooth   = cgraph.AvgDistSmooth
	GraphDistSmooth = cgraph.GraphDistSmooth
	PowerDistSmooth = cgraph.PowerDistSmooth
	RngSmooth       = cgraph.RngSmooth
	SprintSmooth    = cgraph.SprintSmooth
	TriangleSmooth  = cgraph.TriangleSmooth
)

const (
	RegularStart = cgraph.RegularStart
	SelfStart    = cgraph.SelfStart
	RandomStart  = cgraph.RandomStart
)

const (
	SolidGraphStyle   = cgraph.SolidGraphStyle
	DashedGraphStyle  = cgraph.DashedGraphStyle
	DottedGraphStyle  = cgraph.DottedGraphStyle
	BoldGraphStyle    = cgraph.BoldGraphStyle
	RoundedGraphStyle = cgraph.RoundedGraphStyle
	FilledGraphStyle  = cgraph.FilledGraphStyle
	StripedGraphStyle = cgraph.StripedGraphStyle
)

const (
	SolidNodeStyle     = cgraph.SolidNodeStyle
	DashedNodeStyle    = cgraph.DashedNodeStyle
	DottedNodeStyle    = cgraph.DottedNodeStyle
	BoldNodeStyle      = cgraph.BoldNodeStyle
	RoundedNodeStyle   = cgraph.RoundedNodeStyle
	DiagonalsNodeStyle = cgraph.DiagonalsNodeStyle
	FilledNodeStyle    = cgraph.FilledNodeStyle
	StripedNodeStyle   = cgraph.StripedNodeStyle
	WedgedNodeStyle    = cgraph.WedgedNodeStyle
)

const (
	SolidEdgeStyle  = cgraph.SolidEdgeStyle
	DashedEdgeStyle = cgraph.DashedEdgeStyle
	DottedEdgeStyle = cgraph.DottedEdgeStyle
	BoldEdgeStyle   = cgraph.BoldEdgeStyle
)

// functions from cgraph package.
var (
	ParseFile  = cgraph.ParseFile
	ParseBytes = cgraph.ParseBytes
)

// functions from gvc package.
var (
	DefaultPlugins  = gvc.DefaultPlugins
	DeviceQuality   = gvc.WithDeviceQuality
	DeviceFeatures  = gvc.WithDeviceFeatures
	DeviceDPI       = gvc.WithDeviceDPI
	NewDevicePlugin = gvc.NewDevicePlugin
	PNGDevicePlugin = gvc.PNGDevicePlugin
	JPGDevicePlugin = gvc.JPGDevicePlugin
	RenderQuality   = gvc.WithRenderQuality
	RenderFeatures  = gvc.WithRenderFeatures
	RenderColorType = gvc.WithRenderColorType
	RenderPAD       = gvc.WithRenderPAD
	NewRenderPlugin = gvc.NewRenderPlugin
	PNGRenderPlugin = gvc.PNGRenderPlugin
	JPGRenderPlugin = gvc.JPGRenderPlugin
)
