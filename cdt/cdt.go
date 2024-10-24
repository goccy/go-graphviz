package cdt

import (
	"context"
	"errors"

	"github.com/goccy/go-graphviz/internal/wasm"
)

type Dict struct {
	wasm *wasm.Dict
}

func toDict(v *wasm.Dict) *Dict {
	if v == nil {
		return nil
	}
	return &Dict{wasm: v}
}

func (d *Dict) getWasm() *wasm.Dict {
	return d.wasm
}

type Hold struct {
	wasm *wasm.DictHold
}

func toHold(v *wasm.DictHold) *Hold {
	if v == nil {
		return nil
	}
	return &Hold{wasm: v}
}

func (h *Hold) getWasm() *wasm.DictHold {
	return h.wasm
}

type Link struct {
	wasm *wasm.DictLink
}

func toLink(v *wasm.DictLink) *Link {
	if v == nil {
		return nil
	}
	return &Link{wasm: v}
}

func (l *Link) getWasm() *wasm.DictLink {
	return l.wasm
}

type Method struct {
	wasm *wasm.DictMethod
}

func toMethod(v *wasm.DictMethod) *Method {
	if v == nil {
		return nil
	}
	return &Method{wasm: v}
}

func (m *Method) getWasm() *wasm.DictMethod {
	return m.wasm
}

type Data struct {
	wasm *wasm.DictData
}

func toData(v *wasm.DictData) *Data {
	if v == nil {
		return nil
	}
	return &Data{wasm: v}
}

func (d *Data) getWasm() *wasm.DictData {
	return d.wasm
}

type Disc struct {
	wasm *wasm.DictDisc
}

func toDisc(v *wasm.DictDisc) *Disc {
	if v == nil {
		return nil
	}
	return &Disc{wasm: v}
}

func (d *Disc) getWasm() *wasm.DictDisc {
	return d.wasm
}

type Stat struct {
	wasm *wasm.DictStat
}

func toStat(v *wasm.DictStat) *Stat {
	if v == nil {
		return nil
	}
	return &Stat{wasm: v}
}

func (s *Stat) getWasm() *wasm.DictStat {
	return s.wasm
}

type Search func(*Dict, any, int) any
type Make func(*Dict, any, *Disc) any
type Memory func(*Dict, any, uint, *Disc) any
type Free func(*Dict, any, *Disc)
type Compare func(*Dict, any, any, *Disc) int
type Hash func(*Dict, any, *Disc) uint
type Event func(*Dict, int, any, *Disc) int

func StrHash(a1 any, a2 int) (uint, error) {
	return wasm.StrHash(context.Background(), a1, a2)
}

func Open(disc *Disc, mtd *Method) (*Dict, error) {
	res, err := wasm.NewDictWithDisc(context.Background(), disc.getWasm(), mtd.getWasm())
	if err != nil {
		return nil, err
	}
	return toDict(res), nil
}

func (d *Dict) Close() error {
	res, err := d.wasm.Close(context.Background())
	if err != nil {
		return err
	}
	return toError(res)
}

func (d *Dict) View(dict *Dict) (*Dict, error) {
	res, err := d.wasm.View(context.Background(), dict.getWasm())
	if err != nil {
		return nil, err
	}
	return toDict(res), nil
}

func (d *Dict) Disc(disc *Disc) (*Disc, error) {
	res, err := d.wasm.Disc(context.Background(), disc.getWasm())
	if err != nil {
		return nil, err
	}
	return toDisc(res), nil
}

func (d *Dict) Method(mtd *Method) (*Method, error) {
	res, err := d.wasm.Method(context.Background(), mtd.getWasm())
	if err != nil {
		return nil, err
	}
	return toMethod(res), nil
}

func (d *Dict) Flatten() (*Link, error) {
	res, err := d.wasm.Flatten(context.Background())
	if err != nil {
		return nil, err
	}
	return toLink(res), nil
}

func (d *Dict) Extract() (*Link, error) {
	res, err := d.wasm.Extract(context.Background())
	if err != nil {
		return nil, err
	}
	return toLink(res), nil
}

func (d *Dict) Restore(link *Link) error {
	res, err := d.wasm.Restore(context.Background(), link.getWasm())
	if err != nil {
		return err
	}
	return toError(res)
}

func (d *Dict) Walk(fn func(context.Context, *Dict, any, any) error, data any) error {
	// TODO
	res, err := d.wasm.Walk(context.Background(), wasm.CreateCallbackFunc(func(ctx context.Context, a1 any, a2 any) (int, error) {
		if err := fn(ctx, d, a1, a2); err != nil {
			return 0, err
		}
		return 0, nil
	}, wasm.WasmPtr(d.wasm)), data)
	if err != nil {
		return err
	}
	return toError(res)
}

func (d *Dict) Renew(a0 any) (any, error) {
	res, err := d.wasm.Renew(context.Background(), a0)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (d *Dict) Size() (int, error) {
	res, err := d.wasm.Size(context.Background())
	if err != nil {
		return 0, err
	}
	return res, nil
}

func (d *Dict) Stat(a0 *Stat, a1 int) (int, error) {
	res, err := d.wasm.Stat(context.Background(), a0.getWasm(), a1)
	if err != nil {
		return 0, err
	}
	return res, nil
}

func (l *Link) Right() *Link {
	return toLink(l.wasm.GetRight())
}

func (l *Link) SetRight(v *Link) {
	l.wasm.SetRight(v.getWasm())
}

func (l *Link) Left() *Link {
	return toLink(l.wasm.GetLeft())
}

func (l *Link) SetLeft(v *Link) {
	l.wasm.SetLeft(v.getWasm())
}

func (l *Link) Hash() uint {
	return uint(l.wasm.GetHash())
}

func (l *Link) SetHash(v uint) {
	l.wasm.SetHash(uint32(v))
}

func toError(result int) error {
	if result == 0 {
		return nil
	}
	if e, _ := wasm.LastError(context.Background()); e != "" {
		return errors.New(e)
	}
	return nil
}
