package wasm

import (
	"context"
	"io/fs"
	"sync"
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

var (
	fsMu sync.Mutex
)

func SetWasmFileSystem(fs fs.FS) {
	fsMu.Lock()
	mod.fs.subFS = fs
	fsMu.Unlock()
}

func FileSystem() fs.FS {
	fsMu.Lock()
	defer fsMu.Unlock()
	return mod.fs
}
