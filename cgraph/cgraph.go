package cgraph

import (
	"context"
	"errors"
	"os"

	"github.com/goccy/go-graphviz/cdt"
	"github.com/goccy/go-graphviz/internal/wasm"
)

type Graph struct {
	wasm *wasm.Graph
}

func toGraph(v *wasm.Graph) *Graph {
	if v == nil {
		return nil
	}
	return &Graph{wasm: v}
}

func toGraphWasm(v *Graph) *wasm.Graph {
	if v == nil {
		return nil
	}
	return v.wasm
}

func (g *Graph) getWasm() *wasm.Graph {
	if g == nil {
		return nil
	}
	return g.wasm
}

type Node struct {
	wasm *wasm.Node
}

func toNode(v *wasm.Node) *Node {
	if v == nil {
		return nil
	}
	return &Node{wasm: v}
}

func (n *Node) getWasm() *wasm.Node {
	if n == nil {
		return nil
	}
	return n.wasm
}

type SubNode struct {
	wasm *wasm.SubNode
}

func toSubNode(v *wasm.SubNode) *SubNode {
	if v == nil {
		return nil
	}
	return &SubNode{wasm: v}
}

func (n *SubNode) getWasm() *wasm.SubNode {
	if n == nil {
		return nil
	}
	return n.wasm
}

type Edge struct {
	wasm *wasm.Edge
}

func toEdge(v *wasm.Edge) *Edge {
	if v == nil {
		return nil
	}
	return &Edge{wasm: v}
}

func (e *Edge) getWasm() *wasm.Edge {
	if e == nil {
		return nil
	}
	return e.wasm
}

type Desc struct {
	wasm *wasm.GraphDescriptor
}

func toDesc(v *wasm.GraphDescriptor) *Desc {
	if v == nil {
		return nil
	}
	return &Desc{wasm: v}
}

func (d *Desc) getWasm() *wasm.GraphDescriptor {
	if d == nil {
		return nil
	}
	return d.wasm
}

type Disc struct {
	wasm *wasm.ClientDiscipline
}

func toDisc(v *wasm.ClientDiscipline) *Disc {
	if v == nil {
		return nil
	}
	return &Disc{wasm: v}
}

func (d *Disc) getWasm() *wasm.ClientDiscipline {
	if d == nil {
		return nil
	}
	return d.wasm
}

// Symbol symbol in one of the above dictionaries.
type Symbol struct {
	wasm *wasm.Sym
}

func toSymbol(v *wasm.Sym) *Symbol {
	if v == nil {
		return nil
	}
	return &Symbol{wasm: v}
}

func (s *Symbol) getWasm() *wasm.Sym {
	if s == nil {
		return nil
	}
	return s.wasm
}

// Record generic runtime record.
type Record struct {
	wasm *wasm.Record
}

func toRecord(v *wasm.Record) *Record {
	if v == nil {
		return nil
	}
	return &Record{wasm: v}
}

func (r *Record) getWasm() *wasm.Record {
	if r == nil {
		return nil
	}
	return r.wasm
}

type Tag struct {
	wasm *wasm.Tag
}

func toTag(v *wasm.Tag) *Tag {
	if v == nil {
		return nil
	}
	return &Tag{wasm: v}
}

func (t *Tag) getWasm() *wasm.Tag {
	if t == nil {
		return nil
	}
	return t.wasm
}

type Object struct {
	wasm *wasm.Object
}

func toObject(v *wasm.Object) *Object {
	if v == nil {
		return nil
	}
	return &Object{wasm: v}
}

func (o *Object) getWasm() *wasm.Object {
	if o == nil {
		return nil
	}
	return o.wasm
}

type CommonFields struct {
	wasm *wasm.CommonFields
}

func toCommonFields(v *wasm.CommonFields) *CommonFields {
	if v == nil {
		return nil
	}
	return &CommonFields{wasm: v}
}

func (c *CommonFields) getWasm() *wasm.CommonFields {
	if c == nil {
		return nil
	}
	return c.wasm
}

type State struct {
	wasm *wasm.State
}

func toState(v *wasm.State) *State {
	if v == nil {
		return nil
	}
	return &State{wasm: v}
}

func (s *State) getWasm() *wasm.State {
	if s == nil {
		return nil
	}
	return s.wasm
}

type CallbackStack struct {
	wasm *wasm.CallbackStack
}

func toCallbackStack(v *wasm.CallbackStack) *CallbackStack {
	if v == nil {
		return nil
	}
	return &CallbackStack{wasm: v}
}

func (c *CallbackStack) getWasm() *wasm.CallbackStack {
	if c == nil {
		return nil
	}
	return c.wasm
}

type Attr struct {
	wasm *wasm.Attr
}

func toAttr(v *wasm.Attr) *Attr {
	if v == nil {
		return nil
	}
	return &Attr{wasm: v}
}

func (a *Attr) getWasm() *wasm.Attr {
	if a == nil {
		return nil
	}
	return a.wasm
}

type DataDict struct {
	wasm *wasm.DataDict
}

func toDataDict(v *wasm.DataDict) *DataDict {
	if v == nil {
		return nil
	}
	return &DataDict{wasm: v}
}

func (d *DataDict) getWasm() *wasm.DataDict {
	if d == nil {
		return nil
	}
	return d.wasm
}

type ID uint64

func ParseBytes(bytes []byte) (*Graph, error) {
	graph, err := wasm.MemRead(context.Background(), string(bytes))
	if err != nil {
		return nil, err
	}
	if graph == nil {
		return nil, lastError()
	}
	g := toGraph(graph)
	if err := setupNodeLabelIfEmpty(g); err != nil {
		return nil, err
	}
	return g, nil
}

func ParseFile(path string) (*Graph, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return ParseBytes(file)
}

func Open(name string, desc *Desc, disc *Disc) (*Graph, error) {
	graph, err := wasm.Open(context.Background(), name, desc.getWasm(), disc.getWasm())
	if err != nil {
		return nil, err
	}
	if graph == nil {
		return nil, lastError()
	}
	g := toGraph(graph)
	if err := setupNodeLabelIfEmpty(g); err != nil {
		return nil, err
	}
	return g, nil
}

func setupNodeLabelIfEmpty(g *Graph) error {
	n, err := g.FirstNode()
	if err != nil {
		return err
	}
	if n == nil {
		return nil
	}
	if err := setLabelIfEmpty(n); err != nil {
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
		if err := setLabelIfEmpty(n); err != nil {
			return err
		}
	}
	return nil
}

func setLabelIfEmpty(n *Node) error {
	if n.Label() == "" {
		n.SetLabel("\\N")
	}
	return nil
}

type ObjectTag int

var (
	GRAPH   ObjectTag = ObjectTag(wasm.GRAPH)
	NODE    ObjectTag = ObjectTag(wasm.NODE)
	OUTEDGE ObjectTag = ObjectTag(wasm.OUT_EDGE)
	INEDGE  ObjectTag = ObjectTag(wasm.IN_EDGE)
	EDGE    ObjectTag = ObjectTag(wasm.EDGE)
)

func (r *Record) Name() string {
	return r.wasm.GetName()
}

func (r *Record) SetName(v string) {
	r.wasm.SetName(v)
}

func (r *Record) Next() *Record {
	return toRecord(r.wasm.GetNext())
}

func (r *Record) SetNext(v *Record) {
	r.wasm.SetNext(v.getWasm())
}

func (t *Tag) ObjectTag() ObjectTag {
	return ObjectTag(t.wasm.GetObjectType())
}

func (t *Tag) ID() ID {
	return ID(t.wasm.GetId())
}

func (t *Tag) SetID(v ID) {
	t.wasm.SetId(uint64(v))
}

func (o *Object) Tag() *Tag {
	return toTag(o.wasm.GetTag())
}

func (o *Object) SetTag(v *Tag) {
	o.wasm.SetTag(v.getWasm())
}

func (o *Object) Data() *Record {
	return toRecord(o.wasm.GetData())
}

func (o *Object) SetData(v *Record) {
	o.wasm.SetData(v.getWasm())
}

func (o *Object) SafeSet(name, value, def string) error {
	res, err := wasm.SafeSetStr(context.Background(), o.wasm, name, value, def)
	if err != nil {
		return err
	}
	return toError(res)
}

func (n *SubNode) SeqLink() *cdt.Link {
	return toDictLink(n.wasm.GetSeqLink())
}

func (n *SubNode) SetSeqLink(v *cdt.Link) {
	n.wasm.SetSeqLink(toDictLinkWasm(v))
}

func (n *SubNode) IDLink() *cdt.Link {
	return toDictLink(n.wasm.GetIdLink())
}

func (n *SubNode) SetIDLink(v *cdt.Link) {
	n.wasm.SetIdLink(toDictLinkWasm(v))
}

func (n *SubNode) Node() *Node {
	return toNode(n.wasm.GetNode())
}

func (n *SubNode) SetNode(v *Node) {
	n.wasm.SetNode(v.getWasm())
}

func (n *SubNode) InID() *cdt.Link {
	return toDictLink(n.wasm.GetInId())
}

func (n *SubNode) SetInID(v *cdt.Link) {
	n.wasm.SetInId(toDictLinkWasm(v))
}

func (n *SubNode) OutID() *cdt.Link {
	return toDictLink(n.wasm.GetOutId())
}

func (n *SubNode) SetOutID(v *cdt.Link) {
	n.wasm.SetOutId(toDictLinkWasm(v))
}

func (n *SubNode) InSeq() *cdt.Link {
	return toDictLink(n.wasm.GetInSeq())
}

func (n *SubNode) SetInSeq(v *cdt.Link) {
	n.wasm.SetInSeq(toDictLinkWasm(v))
}

func (n *SubNode) OutSeq() *cdt.Link {
	return toDictLink(n.wasm.GetOutSeq())
}

func (n *SubNode) SetOutSeq(v *cdt.Link) {
	n.wasm.SetOutSeq(toDictLinkWasm(v))
}

func (n *Node) Base() *Object {
	return toObject(n.wasm.GetBase())
}

func (n *Node) SetBase(v *Object) {
	n.wasm.SetBase(v.getWasm())
}

func (n *Node) Root() *Graph {
	return toGraph(n.wasm.GetRoot())
}

func (n *Node) SetRootGraph(v *Graph) {
	n.wasm.SetRoot(v.getWasm())
}

func (n *Node) MainSub() *SubNode {
	return toSubNode(n.wasm.GetMainsub())
}

func (n *Node) SetMainSub(v *SubNode) {
	n.wasm.SetMainsub(v.getWasm())
}

func (e *Edge) Base() *Object {
	return toObject(e.wasm.GetBase())
}

func (e *Edge) SetBase(v *Object) {
	e.wasm.SetBase(v.getWasm())
}

func (e *Edge) SeqLink() *cdt.Link {
	return toDictLink(e.wasm.GetSeqLink())
}

func (e *Edge) SetSeqLink(v *cdt.Link) {
	e.wasm.SetSeqLink(toDictLinkWasm(v))
}

func (e *Edge) IDLink() *cdt.Link {
	return toDictLink(e.wasm.GetIdLink())
}

func (e *Edge) SetIDLink(v *cdt.Link) {
	e.wasm.SetIdLink(toDictLinkWasm(v))
}

func (e *Edge) Node() *Node {
	return toNode(e.wasm.GetNode())
}

func (e *Edge) SetNode(v *Node) {
	e.wasm.SetNode(v.getWasm())
}

func (c *CommonFields) Disc() *Disc {
	return toDisc(c.wasm.GetDisc())
}

func (c *CommonFields) SetDisc(v *Disc) {
	c.wasm.SetDisc(v.getWasm())
}

func (c *CommonFields) State() *State {
	return toState(c.wasm.GetState())
}

func (c *CommonFields) SetState(v *State) {
	c.wasm.SetState(v.getWasm())
}

func (c *CommonFields) StrDict() *cdt.Dict {
	return toDict(c.wasm.GetStrdict())
}

func (c *CommonFields) SetStrDict(v *cdt.Dict) {
	c.wasm.SetStrdict(toDictWasm(v))
}

func (c *CommonFields) Seq() [3]uint64 {
	res := c.wasm.GetSeq()
	return [3]uint64{res[0], res[1], res[2]}
}

func (c *CommonFields) SetSeq(v [3]uint64) {
	c.wasm.SetSeq(v[:])
}

func (c *CommonFields) Callback() *CallbackStack {
	return toCallbackStack(c.wasm.GetCb())
}

func (c *CommonFields) SetCallback(v *CallbackStack) {
	c.wasm.SetCb(v.getWasm())
}

func (c *CommonFields) LookupByName() [3]*cdt.Dict {
	res := c.wasm.GetLookupByName()
	return [3]*cdt.Dict{toDict(res[0]), toDict(res[1]), toDict(res[2])}
}

func (c *CommonFields) SetLookupByName(v [3]*cdt.Dict) {
	args := make([]*wasm.Dict, len(v))
	for i := range args {
		args[i] = toDictWasm(v[i])
	}
	c.wasm.SetLookupByName(args)
}

func (c *CommonFields) LookupByID() [3]*cdt.Dict {
	res := c.wasm.GetLookupById()
	return [3]*cdt.Dict{toDict(res[0]), toDict(res[1]), toDict(res[2])}
}

func (c *CommonFields) SetLookupByID(v [3]*cdt.Dict) {
	args := make([]*wasm.Dict, len(v))
	for i := range args {
		args[i] = toDictWasm(v[i])
	}
	c.wasm.SetLookupById(args)
}

func (a *Attr) Header() *Record {
	return toRecord(a.wasm.GetH())
}

func (a *Attr) SetHeader(v *Record) {
	a.wasm.SetH(v.getWasm())
}

func (a *Attr) Dict() *cdt.Dict {
	return toDict(a.wasm.GetDict())
}

func (a *Attr) SetDict(v *cdt.Dict) {
	a.wasm.SetDict(toDictWasm(v))
}

func (a *Attr) Str() []string {
	return a.wasm.GetStr()
}

func (a *Attr) SetStr(v []string) {
	a.wasm.SetStr(v)
}

func (s *Symbol) Link() *cdt.Link {
	return toDictLink(s.wasm.GetLink())
}

func (s *Symbol) SetLink(v *cdt.Link) {
	s.wasm.SetLink(toDictLinkWasm(v))
}

func (s *Symbol) Name() string {
	return s.wasm.GetName()
}

func (s *Symbol) SetName(v string) {
	s.wasm.SetName(v)
}

func (s *Symbol) DefaultValue() string {
	return s.wasm.GetDefval()
}

func (s *Symbol) SetDefaultValue(v string) {
	s.wasm.SetDefval(v)
}

func (s *Symbol) ID() int {
	return int(s.wasm.GetId())
}

func (s *Symbol) SetID(v int) {
	s.wasm.SetId(int32(v))
}

func (s *Symbol) Kind() uint {
	return uint(s.wasm.GetKind())
}

func (s *Symbol) SetKind(v uint) {
	s.wasm.SetKind(uint32(v))
}

func (s *Symbol) Fixed() uint {
	return uint(s.wasm.GetFixed())
}

func (s *Symbol) SetFixed(v uint) {
	s.wasm.SetFixed(uint32(v))
}

func (s *Symbol) Print() uint {
	return uint(s.wasm.GetPrint())
}

func (s *Symbol) SetPrint(v uint) {
	s.wasm.SetPrint(uint32(v))
}

func (d *DataDict) Header() *Record {
	return toRecord(d.wasm.GetH())
}

func (d *DataDict) SetHeader(v *Record) {
	d.wasm.SetH(v.getWasm())
}

func (g *Graph) Base() *Object {
	return toObject(g.wasm.GetBase())
}

func (g *Graph) SetBase(v *Object) {
	g.wasm.SetBase(v.getWasm())
}

func (g *Graph) Desc() *Desc {
	return toDesc(g.wasm.GetDesc())
}

func (g *Graph) SetDesc(v *Desc) {
	g.wasm.SetDesc(v.getWasm())
}

func (g *Graph) SeqLink() *cdt.Link {
	return toDictLink(g.wasm.GetSeqLink())
}

func (g *Graph) SetSeqLink(v *cdt.Link) {
	g.wasm.SetSeqLink(toDictLinkWasm(v))
}

func (g *Graph) IDLink() *cdt.Link {
	return toDictLink(g.wasm.GetIdLink())
}

func (g *Graph) SetIDLink(v *cdt.Link) {
	g.wasm.SetIdLink(toDictLinkWasm(v))
}

func (g *Graph) NSeq() *cdt.Dict {
	return toDict(g.wasm.GetNSeq())
}

func (g *Graph) SetNSeq(v *cdt.Dict) {
	g.wasm.SetNSeq(toDictWasm(v))
}

func (g *Graph) ESeq() *cdt.Dict {
	return toDict(g.wasm.GetESeq())
}

func (g *Graph) SetESeq(v *cdt.Dict) {
	g.wasm.SetESeq(toDictWasm(v))
}

func (g *Graph) EID() *cdt.Dict {
	return toDict(g.wasm.GetEId())
}

func (g *Graph) SetEID(v *cdt.Dict) {
	g.wasm.SetEId(toDictWasm(v))
}

func (g *Graph) GSeq() *cdt.Dict {
	return toDict(g.wasm.GetGSeq())
}

func (g *Graph) SetGSeq(v *cdt.Dict) {
	g.wasm.SetGSeq(toDictWasm(v))
}

func (g *Graph) GID() *cdt.Dict {
	return toDict(g.wasm.GetGId())
}

func (g *Graph) SetGID(v *cdt.Dict) {
	g.wasm.SetGId(toDictWasm(v))
}

func (g *Graph) Parent() *Graph {
	return toGraph(g.wasm.GetParent())
}

func (g *Graph) SetParent(v *Graph) {
	g.wasm.SetParent(v.getWasm())
}

func (g *Graph) GraphRoot() *Graph {
	return toGraph(g.wasm.GetRoot())
}

func (g *Graph) SetGraphRoot(v *Graph) {
	g.wasm.SetRoot(v.getWasm())
}

func (g *Graph) CommonFields() *CommonFields {
	return toCommonFields(g.wasm.GetClos())
}

func (g *Graph) SetCommonFields(v *CommonFields) {
	g.wasm.SetClos(v.getWasm())
}

func (g *Graph) CopyAttr(t *Graph) error {
	res, err := wasm.CopyAttr(context.Background(), g.wasm, t.getWasm())
	if err != nil {
		return err
	}
	return toError(res)
}

// BindRecord attach a new record of the given size to the object.
func (g *Graph) BindRecord(name string, size uint, moveToFront int) error {
	if _, err := wasm.BindRecord(context.Background(), g.wasm, name, size, moveToFront); err != nil {
		return err
	}
	return nil
}

func (g *Graph) Record(name string, moveToFront int) (*Record, error) {
	res, err := wasm.GetRecord(context.Background(), g.wasm, name, moveToFront)
	if err != nil {
		return nil, err
	}
	return toRecord(res), nil
}

func (g *Graph) DeleteRecord(name string) error {
	res, err := wasm.DeleteRecord(context.Background(), g.wasm, name)
	if err != nil {
		return err
	}
	return toError(res)
}

func (g *Graph) GetStr(name string) string {
	v, _ := wasm.GetStr(context.Background(), g.wasm, name)
	return v
}

func (g *Graph) SymbolName(sym *Symbol) (string, error) {
	return wasm.GetSymName(context.Background(), g.wasm, sym.getWasm())
}

func (g *Graph) Set(name, value string) error {
	res, err := wasm.SetStr(context.Background(), g.wasm, name, value)
	if err != nil {
		return err
	}
	return toError(res)
}

func (g *Graph) SetSymbolName(sym *Symbol, value string) error {
	res, err := wasm.SetSymName(context.Background(), g.wasm, sym.getWasm(), value)
	if err != nil {
		return err
	}
	return toError(res)
}

func (g *Graph) SafeSet(name, value, def string) error {
	res, err := wasm.SafeSetStr(context.Background(), g.wasm, name, value, def)
	if err != nil {
		return err
	}
	return toError(res)
}

func (g *Graph) Close() error {
	res, err := g.wasm.Close(context.Background())
	if err != nil {
		return err
	}
	return toError(res)
}

func (g *Graph) IsSimple() (bool, error) {
	res, err := g.wasm.IsSimple(context.Background())
	if err != nil {
		return false, err
	}
	return res == 1, nil
}

func (g *Graph) CreateNodeByName(name string) (*Node, error) {
	res, err := g.wasm.Node(context.Background(), name, 1)
	if err != nil {
		return nil, err
	}
	return toNode(res), nil
}

func (g *Graph) NodeByName(name string) (*Node, error) {
	res, err := g.wasm.Node(context.Background(), name, 0)
	if err != nil {
		return nil, err
	}
	return toNode(res), nil
}

func (g *Graph) CreateNodeByID(id ID) (*Node, error) {
	res, err := g.wasm.IdNode(context.Background(), uint64(id), 1)
	if err != nil {
		return nil, err
	}
	return toNode(res), nil
}

func (g *Graph) NodeByID(id ID) (*Node, error) {
	res, err := g.wasm.IdNode(context.Background(), uint64(id), 0)
	if err != nil {
		return nil, err
	}
	return toNode(res), nil
}

func (g *Graph) CreateSubNode(n *Node) (*Node, error) {
	res, err := g.wasm.SubNode(context.Background(), n.getWasm(), 1)
	if err != nil {
		return nil, err
	}
	return toNode(res), nil
}

func (g *Graph) SubNode(n *Node) (*Node, error) {
	res, err := g.wasm.SubNode(context.Background(), n.getWasm(), 0)
	if err != nil {
		return nil, err
	}
	return toNode(res), nil
}

func (g *Graph) FirstNode() (*Node, error) {
	res, err := g.wasm.FirstNode(context.Background())
	if err != nil {
		return nil, err
	}
	return toNode(res), nil
}

func (g *Graph) NextNode(n *Node) (*Node, error) {
	res, err := g.wasm.NextNode(context.Background(), n.getWasm())
	if err != nil {
		return nil, err
	}
	return toNode(res), nil
}

func (g *Graph) LastNode() (*Node, error) {
	res, err := g.wasm.LastNode(context.Background())
	if err != nil {
		return nil, err
	}
	return toNode(res), nil
}

func (g *Graph) PreviousNode(n *Node) (*Node, error) {
	res, err := g.wasm.PrevNode(context.Background(), n.getWasm())
	if err != nil {
		return nil, err
	}
	return toNode(res), nil
}

func (g *Graph) SubRep(n *Node) (*SubNode, error) {
	res, err := g.wasm.SubRep(context.Background(), n.getWasm())
	if err != nil {
		return nil, err
	}
	return toSubNode(res), nil
}

func (g *Graph) CreateEdgeByName(name string, start *Node, end *Node) (*Edge, error) {
	res, err := g.wasm.Edge(context.Background(), start.getWasm(), end.getWasm(), name, 1)
	if err != nil {
		return nil, err
	}
	return toEdge(res), nil
}

func (g *Graph) EdgeByName(name string, start *Node, end *Node) (*Edge, error) {
	res, err := g.wasm.Edge(context.Background(), start.getWasm(), end.getWasm(), name, 0)
	if err != nil {
		return nil, err
	}
	return toEdge(res), nil
}

func (g *Graph) CreateEdgeByID(id ID, start *Node, end *Node) (*Edge, error) {
	res, err := g.wasm.IdEdge(context.Background(), start.getWasm(), end.getWasm(), uint64(id), 1)
	if err != nil {
		return nil, err
	}
	return toEdge(res), nil
}

func (g *Graph) EdgeByID(id ID, start *Node, end *Node) (*Edge, error) {
	res, err := g.wasm.IdEdge(context.Background(), start.getWasm(), end.getWasm(), uint64(id), 0)
	if err != nil {
		return nil, err
	}
	return toEdge(res), nil
}

func (g *Graph) CreateSubEdge(e *Edge) (*Edge, error) {
	res, err := g.wasm.SubEdge(context.Background(), e.getWasm(), 1)
	if err != nil {
		return nil, err
	}
	return toEdge(res), nil
}

func (g *Graph) SubEdge(e *Edge) (*Edge, error) {
	res, err := g.wasm.SubEdge(context.Background(), e.getWasm(), 0)
	if err != nil {
		return nil, err
	}
	return toEdge(res), nil
}

func (g *Graph) FirstIn(n *Node) (*Edge, error) {
	res, err := g.wasm.FirstIn(context.Background(), n.getWasm())
	if err != nil {
		return nil, err
	}
	return toEdge(res), nil
}

func (g *Graph) FirstOut(n *Node) (*Edge, error) {
	res, err := g.wasm.FirstOut(context.Background(), n.getWasm())
	if err != nil {
		return nil, err
	}
	return toEdge(res), nil
}

func (g *Graph) NextIn(e *Edge) (*Edge, error) {
	res, err := g.wasm.NextIn(context.Background(), e.getWasm())
	if err != nil {
		return nil, err
	}
	return toEdge(res), nil
}

func (g *Graph) NextOut(e *Edge) (*Edge, error) {
	res, err := g.wasm.NextOut(context.Background(), e.getWasm())
	if err != nil {
		return nil, err
	}
	return toEdge(res), nil
}

func (g *Graph) FirstEdge(n *Node) (*Edge, error) {
	res, err := g.wasm.FirstEdge(context.Background(), n.getWasm())
	if err != nil {
		return nil, err
	}
	return toEdge(res), nil
}

func (g *Graph) NextEdge(e *Edge, n *Node) (*Edge, error) {
	res, err := g.wasm.NextEdge(context.Background(), e.getWasm(), n.getWasm())
	if err != nil {
		return nil, err
	}
	return toEdge(res), nil
}

func (g *Graph) Contains(o any) (bool, error) {
	res, err := g.wasm.Contains(context.Background(), o)
	if err != nil {
		return false, err
	}
	return res == 1, nil
}

func (g *Graph) Name() (string, error) {
	return wasm.GraphNameOf(context.Background(), g.wasm)
}

func (g *Graph) Delete(obj any) error {
	res, err := g.wasm.Delete(context.Background(), obj)
	if err != nil {
		return err
	}
	return toError(res)
}

func (g *Graph) DeleteSubGraph(sub *Graph) error {
	res, err := g.wasm.DeleteSubGraph(context.Background(), sub.getWasm())
	if err != nil {
		return err
	}
	return toError(res)
}

func (g *Graph) DeleteNode(n *Node) (bool, error) {
	res, err := g.wasm.DeleteNode(context.Background(), n.getWasm())
	if err != nil {
		return false, err
	}
	return res == 1, nil
}

func (g *Graph) DeleteEdge(e *Edge) (bool, error) {
	res, err := g.wasm.DeleteEdge(context.Background(), e.getWasm())
	if err != nil {
		return false, err
	}
	return res == 1, nil
}

func (g *Graph) Strdup(s string) (string, error) {
	return g.wasm.Strdup(context.Background(), s)
}

func (g *Graph) StrdupHTML(s string) (string, error) {
	return g.wasm.StrdupHTML(context.Background(), s)
}

func (g *Graph) StrBind(s string) (string, error) {
	return g.wasm.StrBind(context.Background(), s)
}

func (g *Graph) StrFree(s string) error {
	res, err := g.wasm.StrFree(context.Background(), s)
	if err != nil {
		return err
	}
	return toError(res)
}

func (g *Graph) Attr(kind int, name, value string) (*Symbol, error) {
	res, err := g.wasm.Attr(context.Background(), kind, name, value)
	if err != nil {
		return nil, err
	}
	return toSymbol(res), nil
}

func (g *Graph) NextAttr(kind int, attr *Symbol) (*Symbol, error) {
	res, err := g.wasm.NextAttr(context.Background(), kind, attr.getWasm())
	if err != nil {
		return nil, err
	}
	return toSymbol(res), nil
}

func (g *Graph) Init(kind int, recName string, recSize int, moveToFront int) error {
	return g.wasm.Init(context.Background(), kind, recName, recSize, moveToFront)
}

func (g *Graph) Clean(kind int, recName string) error {
	return g.wasm.Clean(context.Background(), kind, recName)
}

func (g *Graph) CreateSubGraphByName(name string) (*Graph, error) {
	res, err := g.wasm.SubGraph(context.Background(), name, 1)
	if err != nil {
		return nil, err
	}
	return toGraph(res), nil
}

func (g *Graph) SubGraphByName(name string) (*Graph, error) {
	res, err := g.wasm.SubGraph(context.Background(), name, 0)
	if err != nil {
		return nil, err
	}
	return toGraph(res), nil
}

func (g *Graph) CreateSubGraphByID(id ID) (*Graph, error) {
	res, err := g.wasm.IdSubGraph(context.Background(), uint64(id), 1)
	if err != nil {
		return nil, err
	}
	return toGraph(res), nil
}

func (g *Graph) SubGraphByID(id ID) (*Graph, error) {
	res, err := g.wasm.IdSubGraph(context.Background(), uint64(id), 0)
	if err != nil {
		return nil, err
	}
	return toGraph(res), nil
}

func (g *Graph) FirstSubGraph() (*Graph, error) {
	res, err := g.wasm.FirstSubGraph(context.Background())
	if err != nil {
		return nil, err
	}
	return toGraph(res), nil
}

func (g *Graph) NextSubGraph() (*Graph, error) {
	res, err := g.wasm.NextSubGraph(context.Background())
	if err != nil {
		return nil, err
	}
	return toGraph(res), nil
}

func (g *Graph) NodeNum() (int, error) {
	return g.wasm.NodeNum(context.Background())
}

func (g *Graph) EdgeNum() (int, error) {
	return g.wasm.EdgeNum(context.Background())
}

func (g *Graph) SubGraphNum() (int, error) {
	return g.wasm.SubGraphNum(context.Background())
}

// Degree returns the degree of the given node in the graph, where arguments "in" and
// "out" are C-like booleans that select which edge sets to query.
//
// g.Degree(node, 0, 0) // always returns 0
// g.Degree(node, 0, 1) // returns the node's outdegree
// g.Degree(node, 1, 0) // returns the node's indegree
// g.Degree(node, 1, 1) // returns the node's total degree (indegree + outdegree).
func (g *Graph) Degree(n *Node, in, out int) (int, error) {
	return g.wasm.Degree(context.Background(), n.getWasm(), in, out)
}

// Indegree returns the indegree of the given node in the graph.
//
// Note: While undirected graphs don't normally have a
// notion of indegrees, calling this method on an
// undirected graph will treat it as if it's directed.
// As a result, it's best to avoid calling this method
// on an undirected graph.
func (g *Graph) Indegree(n *Node) (int, error) {
	return g.wasm.Degree(context.Background(), n.getWasm(), 1, 0)
}

// Outdegree returns the outdegree of the given node in the graph.
//
// Note: While undirected graphs don't normally have a
// notion of outdegrees, calling this method on an
// undirected graph will treat it as if it's directed.
// As a result, it's best to avoid calling this method
// on an undirected graph.
func (g *Graph) Outdegree(n *Node) (int, error) {
	return g.wasm.Degree(context.Background(), n.getWasm(), 0, 1)
}

// TotalDegree returns the total degree of the given node in the graph.
// This can be thought of as the total number of edges coming
// in and out of a node.
func (g *Graph) TotalDegree(n *Node) (int, error) {
	return g.wasm.Degree(context.Background(), n.getWasm(), 1, 1)
}

func (g *Graph) CountUniqueEdges(n *Node, in, out int) (int, error) {
	return g.wasm.CountUniqueEdges(context.Background(), n.getWasm(), in, out)
}

func (n *Node) Name() (string, error) {
	return wasm.GraphNameOf(context.Background(), n.wasm)
}

func (n *Node) CopyAttr(t *Node) error {
	res, err := wasm.CopyAttr(context.Background(), n.wasm, t.getWasm())
	if err != nil {
		return err
	}
	return toError(res)
}

func (n *Node) BindRecord(name string, size uint, moveToFront int) error {
	if _, err := wasm.BindRecord(context.Background(), n.wasm, name, size, moveToFront); err != nil {
		return err
	}
	return nil
}

func (n *Node) Record(name string, moveToFront int) (*Record, error) {
	res, err := wasm.GetRecord(context.Background(), n.wasm, name, moveToFront)
	if err != nil {
		return nil, err
	}
	return toRecord(res), nil
}

func (n *Node) DeleteRecord(name string) error {
	res, err := wasm.DeleteRecord(context.Background(), n.wasm, name)
	if err != nil {
		return err
	}
	return toError(res)
}

func (n *Node) GetStr(name string) string {
	v, _ := wasm.GetStr(context.Background(), n.wasm, name)
	return v
}

func (n *Node) SymbolName(sym *Symbol) (string, error) {
	return wasm.GetSymName(context.Background(), n.wasm, sym.getWasm())
}

func (n *Node) Set(name, value string) error {
	res, err := wasm.SetStr(context.Background(), n.wasm, name, value)
	if err != nil {
		return err
	}
	return toError(res)
}

func (n *Node) SetSymbolName(sym *Symbol, value string) error {
	res, err := wasm.SetSymName(context.Background(), n.wasm, sym.getWasm(), value)
	if err != nil {
		return err
	}
	return toError(res)
}

func (n *Node) SafeSet(name, value, def string) error {
	res, err := wasm.SafeSetStr(context.Background(), n.wasm, name, value, def)
	if err != nil {
		return err
	}
	return toError(res)
}

func (n *Node) ReLabel(newname string) error {
	res, err := n.wasm.ReLabel(context.Background(), newname)
	if err != nil {
		return err
	}
	return toError(res)
}

func (n *Node) Before(v *Node) error {
	res, err := n.wasm.Before(context.Background(), v.getWasm())
	if err != nil {
		return err
	}
	return toError(res)
}

func (e *Edge) Name() (string, error) {
	return wasm.GraphNameOf(context.Background(), e.wasm)
}

func (e *Edge) CopyAttr(t *Edge) error {
	res, err := wasm.CopyAttr(context.Background(), e.wasm, t.getWasm())
	if err != nil {
		return err
	}
	return toError(res)
}

func (e *Edge) BindRecord(name string, size uint, moveToFront int) error {
	if _, err := wasm.BindRecord(context.Background(), e.wasm, name, size, moveToFront); err != nil {
		return err
	}
	return nil
}

func (e *Edge) Record(name string, moveToFront int) (*Record, error) {
	res, err := wasm.GetRecord(context.Background(), e.wasm, name, moveToFront)
	if err != nil {
		return nil, err
	}
	return toRecord(res), nil
}

func (e *Edge) DeleteRecord(name string) error {
	res, err := wasm.DeleteRecord(context.Background(), e.wasm, name)
	if err != nil {
		return err
	}
	return toError(res)
}

func (e *Edge) GetStr(name string) string {
	v, _ := wasm.GetStr(context.Background(), e.wasm, name)
	return v
}

func (e *Edge) SymbolName(sym *Symbol) (string, error) {
	return wasm.GetSymName(context.Background(), e.wasm, sym.getWasm())
}

func (e *Edge) Set(name, value string) error {
	res, err := wasm.SetStr(context.Background(), e.wasm, name, value)
	if err != nil {
		return err
	}
	return toError(res)
}

func (e *Edge) SetSymbolName(sym *Symbol, value string) error {
	res, err := wasm.SetSymName(context.Background(), e.wasm, sym.getWasm(), value)
	if err != nil {
		return err
	}
	return toError(res)
}

func (e *Edge) SafeSet(name, value, def string) error {
	res, err := wasm.SafeSetStr(context.Background(), e.wasm, name, value, def)
	if err != nil {
		return err
	}
	return toError(res)
}

func HTMLStr(s string) (bool, error) {
	return wasm.HtmlStr(context.Background(), s)
}

func Canon(s string, i int) (string, error) {
	return wasm.Canon(context.Background(), s, i)
}

func StrCanon(a0 string, a1 string) (string, error) {
	return wasm.StrCanon(context.Background(), a0, a1)
}

func CanonStr(str string) (string, error) {
	return wasm.CanonStr(context.Background(), str)
}

func AttrSym(obj *Object, name string) (*Symbol, error) {
	sym, err := wasm.AttrSym(context.Background(), obj.getWasm(), name)
	if err != nil {
		return nil, err
	}
	return toSymbol(sym), nil
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
