package gvc

import (
	_ "unsafe"

	"github.com/goccy/go-graphviz/cdt"
	"github.com/goccy/go-graphviz/cgraph"
	"github.com/goccy/go-graphviz/internal/wasm"
)

//go:linkname toGraph github.com/goccy/go-graphviz/cgraph.toGraph
func toGraph(*wasm.Graph) *cgraph.Graph

//go:linkname toGraphWasm github.com/goccy/go-graphviz/cgraph.toGraphWasm
func toGraphWasm(*cgraph.Graph) *wasm.Graph

//go:linkname toNode github.com/goccy/go-graphviz/cgraph.toNode
func toNode(*wasm.Node) *cgraph.Node

//go:linkname toEdge github.com/goccy/go-graphviz/cgraph.toEdge
func toEdge(*wasm.Edge) *cgraph.Edge

//go:linkname toDictLink github.com/goccy/go-graphviz/cdt.toLink
func toDictLink(*wasm.DictLink) *cdt.Link

//go:linkname toDictLinkWasm github.com/goccy/go-graphviz/cdt.toLinkWasm
func toDictLinkWasm(*cdt.Link) *wasm.DictLink
