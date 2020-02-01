package gvc

import (
	"bytes"

	"github.com/fogleman/gg"
	"github.com/goccy/go-graphviz/internal/ccall"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/gobold"
	"golang.org/x/xerrors"
)

type pngRenderer struct {
	*DefaultRenderer
	ctx *gg.Context
}

func (r *pngRenderer) BeginPage(job *Job) error {
	scale := job.Scale()
	translation := job.Translation()
	ctx := gg.NewContext(int(job.Width()), int(job.Height()))
	ctx.Scale(scale.X, scale.Y)
	ctx.Translate(translation.X, -translation.Y)
	r.ctx = ctx
	return nil
}

func (r *pngRenderer) EndPage(job *Job) error {
	if job.OutputData() != nil {
		var buf bytes.Buffer
		if err := r.ctx.EncodePNG(&buf); err != nil {
			return xerrors.Errorf("failed to encode png: %w", err)
		}
		job.SetOutputData(buf.Bytes())
	}
	filename := job.OutputFilename()
	if filename != "" {
		if err := r.ctx.SavePNG(job.OutputFilename()); err != nil {
			return xerrors.Errorf("failed to save png: %w", err)
		}
	}
	return nil
}

func (r *pngRenderer) TextSpan(job *Job, p Pointf, span *TextSpan) error {
	r.ctx.Push()
	defer r.ctx.Pop()

	r.ctx.SetRGB(0, 0, 0)

	ft, err := truetype.Parse(gobold.TTF)
	if err != nil {
		return xerrors.Errorf("failed to parse truetype gobold.TTF: %w", err)
	}
	opt := &truetype.Options{
		Size:              span.Font().Size(),
		DPI:               0,
		Hinting:           0,
		GlyphCacheEntries: 0,
		SubPixelsX:        0,
		SubPixelsY:        0,
	}
	face := truetype.NewFace(ft, opt)
	r.ctx.SetFontFace(face)
	y := p.Y + span.YOffsetCenterLine() + span.YOffsetLayout()
	r.ctx.DrawStringAnchored(span.Str(), p.X, -y, 0.5, 0)
	return nil
}

func (r *pngRenderer) Ellipse(job *Job, a0, a1 Pointf, filled int) error {
	r.ctx.Push()
	defer r.ctx.Pop()
	rx := a1.X - a0.X
	ry := a1.Y - a0.Y
	var c ccall.GVColor
	if filled > 0 {
		c = job.Obj().FillColor()
		r.ctx.FillPreserve()
	} else {
		c = job.Obj().PenColor()
	}
	r.ctx.SetRGB(float64(c.R)/255.0, float64(c.G)/255.0, float64(c.B)/255.0)
	r.ctx.DrawEllipse(a0.X, -a0.Y, rx, ry)
	if filled > 0 {
		r.ctx.Fill()
	} else {
		r.ctx.Stroke()
	}
	return nil
}

func (r *pngRenderer) Polygon(job *Job, a []Pointf, filled int) error {
	r.ctx.Push()
	defer r.ctx.Pop()
	var c ccall.GVColor
	if filled > 0 {
		c = job.Obj().FillColor()
	} else {
		c = job.Obj().PenColor()
	}
	r.ctx.SetRGB(float64(c.R)/255.0, float64(c.G)/255.0, float64(c.B)/255.0)
	r.ctx.MoveTo(a[0].X, -a[0].Y)
	for i := 1; i < len(a); i++ {
		r.ctx.LineTo(a[i].X, -a[i].Y)
	}
	r.ctx.ClosePath()
	if filled > 0 {
		r.ctx.Fill()
	} else {
		r.ctx.Stroke()
	}
	return nil
}

func (r *pngRenderer) BezierCurve(job *Job, a []Pointf, arrowAtStart, arrowAtEnd int) error {
	r.ctx.Push()
	defer r.ctx.Pop()
	c := job.Obj().PenColor()
	r.ctx.SetRGB(float64(c.R)/255.0, float64(c.G)/255.0, float64(c.B)/255.0)
	r.ctx.MoveTo(a[0].X, -a[0].Y)
	for i := 1; i < len(a); i += 3 {
		r.ctx.CubicTo(a[i].X, -a[i].Y, a[i+1].X, -a[i+1].Y, a[i+2].X, -a[i+2].Y)
	}
	r.ctx.Stroke()
	return nil
}

func init() {
	RegisterRenderer("png", &pngRenderer{})
}
