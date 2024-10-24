package cgraph

import (
	_ "unsafe"

	"github.com/goccy/go-graphviz/cdt"
	"github.com/goccy/go-graphviz/internal/wasm"
)

//go:linkname toDict github.com/goccy/go-graphviz/cdt.toDict
func toDict(*wasm.Dict) *cdt.Dict

//go:linkname toDictWasm github.com/goccy/go-graphviz/cdt.toDictWasm
func toDictWasm(*cdt.Dict) *wasm.Dict

//go:linkname toDictLink github.com/goccy/go-graphviz/cdt.toLink
func toDictLink(*wasm.DictLink) *cdt.Link

//go:linkname toDictLinkWasm github.com/goccy/go-graphviz/cdt.toLinkWasm
func toDictLinkWasm(*cdt.Link) *wasm.DictLink

func toGraphWasm(v *Graph) *wasm.Graph {
	if v == nil {
		return nil
	}
	return v.wasm
}
