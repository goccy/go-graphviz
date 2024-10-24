package cdt

import (
	"github.com/goccy/go-graphviz/internal/wasm"
)

func toLinkWasm(v *Link) *wasm.DictLink {
	if v == nil {
		return nil
	}
	return v.wasm
}

func toDictWasm(v *Dict) *wasm.Dict {
	if v == nil {
		return nil
	}
	return v.wasm
}
