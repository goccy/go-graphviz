package cdt

import (
	"unsafe"

	"github.com/goccy/go-graphviz/internal/ccall"
)

type Dict struct {
	*ccall.Dict
}

type Hold struct {
	*ccall.Dthold
}

type Link struct {
	*ccall.Dtlink
}

type Method struct {
	*ccall.Dtmethod
}

type Data struct {
	*ccall.Dtdata
}

type Disc struct {
	*ccall.Dtdisc
}

type Stat struct {
	*ccall.Dtstat
}

type Search func(*Dict, unsafe.Pointer, int) unsafe.Pointer
type Make func(*Dict, unsafe.Pointer, *Disc) unsafe.Pointer
type Memory func(*Dict, unsafe.Pointer, uint, *Disc) unsafe.Pointer
type Free func(*Dict, unsafe.Pointer, *Disc)
type Compare func(*Dict, unsafe.Pointer, unsafe.Pointer, *Disc) int
type Hash func(*Dict, unsafe.Pointer, *Disc) uint
type Event func(*Dict, int, unsafe.Pointer, *Disc) int

func StrHash(a0 uint, a1 unsafe.Pointer, a2 int) uint {
	return ccall.Dtstrhash(a0, a1, a2)
}

func toDict(v *ccall.Dict) *Dict {
	if v == nil {
		return nil
	}
	return &Dict{Dict: v}
}

func toDisc(v *ccall.Dtdisc) *Disc {
	if v == nil {
		return nil
	}
	return &Disc{Dtdisc: v}
}

func toData(v *ccall.Dtdata) *Data {
	if v == nil {
		return nil
	}
	return &Data{Dtdata: v}
}

func toLink(v *ccall.Dtlink) *Link {
	if v == nil {
		return nil
	}
	return &Link{Dtlink: v}
}

func Open(a0 *Disc, a1 *Method) *Dict {
	return toDict(ccall.Dtopen(a0.Dtdisc, a1.Dtmethod))
}

func (d *Dict) Close() int {
	return ccall.Dtclose(d.Dict)
}

func (d *Dict) View(a0 *Dict) *Dict {
	return toDict(ccall.Dtview(d.Dict, a0.Dict))
}

func (d *Dict) Disc(a0 *Disc, a1 int) *Disc {
	return toDisc(ccall.Dtdiscf(d.Dict, a0.Dtdisc, a1))
}

func (d *Dict) GetDisc() *Disc {
	return toDisc(d.Dict.Disc())
}

func (d *Dict) SetDisc(v *Disc) {
	if v == nil {
		return
	}
	d.Dict.SetDisc(v.Dtdisc)
}

func (d *Dict) Data() *Data {
	return toData(d.Dict.Data())
}

func (d *Dict) SetData(v *Data) {
	if v == nil {
		return
	}
	d.Dict.SetData(v.Dtdata)
}

func (d *Dict) Method(a0 *Method) *Method {
	return &Method{Dtmethod: ccall.Dtmethodf(d.Dict, a0.Dtmethod)}
}

func (d *Dict) Flatten() *Link {
	return &Link{Dtlink: ccall.Dtflatten(d.Dict)}
}

func (d *Dict) Extract() *Link {
	return &Link{Dtlink: ccall.Dtextract(d.Dict)}
}

func (d *Dict) Restore(a0 *Link) int {
	return ccall.Dtrestore(d.Dict, a0.Dtlink)
}

func (d *Dict) TreeSet(a0 int, a1 int) int {
	return ccall.Dttreeset(d.Dict, a0, a1)
}

func (d *Dict) Walk(walk func(a0 *Dict, a1 unsafe.Pointer, a2 unsafe.Pointer) int, data unsafe.Pointer) int {
	return ccall.Dtwalk(d.Dict, func(a0 *ccall.Dict, a1 unsafe.Pointer, a2 unsafe.Pointer) int {
		return walk(&Dict{Dict: a0}, a1, a2)
	}, data)
}

func (d *Dict) Renew(a0 unsafe.Pointer) unsafe.Pointer {
	return ccall.Dtrenew(d.Dict, a0)
}

func (d *Dict) Size() int {
	return ccall.Dtsize(d.Dict)
}

func (d *Dict) Stat(a0 *Stat, a1 int) int {
	return ccall.Dtstatf(d.Dict, a0.Dtstat, a1)
}

func (l *Link) Right() *Link {
	return toLink(l.Dtlink.Right())
}

func (l *Link) SetRight(v *Link) {
	if v == nil {
		return
	}
	l.Dtlink.SetRight(v.Dtlink)
}

func (l *Link) Left() *Link {
	return toLink(l.Dtlink.Left())
}

func (l *Link) SetLeft(v *Link) {
	if v == nil {
		return
	}
	l.Dtlink.SetLeft(v.Dtlink)
}

func (l *Link) Hash() uint {
	return l.Dtlink.Hash()
}

func (l *Link) SetHash(v uint) {
	l.Dtlink.SetHash(v)
}

func (h *Hold) Header() *Link {
	return toLink(h.Dthold.Hdr())
}

func (h *Hold) SetHeader(v *Link) {
	if v == nil {
		return
	}
	h.Dthold.SetHdr(v.Dtlink)
}

func (h *Hold) Object() unsafe.Pointer {
	return h.Dthold.Obj()
}

func (h *Hold) SetObject(v unsafe.Pointer) {
	h.Dthold.SetObj(v)
}
