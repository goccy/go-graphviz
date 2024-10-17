package gvc

import (
	_ "unsafe"

	"github.com/goccy/go-graphviz/cgraph"
	"github.com/goccy/go-graphviz/internal/wasm"
)

//go:linkname toGraphWasm github.com/goccy/go-graphviz/cgraph.toGraphWasm
func toGraphWasm(*cgraph.Graph) *wasm.Graph
