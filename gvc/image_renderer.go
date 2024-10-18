package gvc

import (
	"bytes"
	"context"
	"image/jpeg"
	"io"
	"os"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
)

type ImageRenderer struct {
	*DefaultRenderEngine
	ctx      *gg.Context
	fontFace func(float64) (font.Face, error)
}

func (r *ImageRenderer) SetFontFace(fn func(size float64) (font.Face, error)) {
	r.fontFace = fn
}

func (r *ImageRenderer) toX(job *Job, x float64) float64 {
	return job.Scale().X() * x
}

func (r *ImageRenderer) toY(job *Job, y float64) float64 {
	return job.Scale().Y() * y
}

func (r *ImageRenderer) BeginPage(ctx context.Context, job *Job) error {
	gctx := gg.NewContext(int(job.Width()), int(job.Height()))
	translation := job.Translation()
	gctx.Translate(r.toX(job, translation.X()), r.toY(job, -translation.Y()))
	r.ctx = gctx
	return nil
}

func (r *ImageRenderer) isPNG(job *Job) bool {
	return job.OutputLangName() == "png"
}

func (r *ImageRenderer) isJPG(job *Job) bool {
	return job.OutputLangName() == "jpg"
}

func (r *ImageRenderer) encodeJPG(w io.Writer) error {
	return jpeg.Encode(w, r.ctx.Image(), &jpeg.Options{
		Quality: jpeg.DefaultQuality,
	})
}

func (r *ImageRenderer) saveJPG(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return r.encodeJPG(file)
}

func (r *ImageRenderer) setPenStyle(job *Job) {
	o := job.Object()
	switch o.Pen() {
	case PenDashed:
		r.ctx.SetDash(4.0)
	case PenDotted:
		r.ctx.SetDash(2.0, 4.0)
	case PenSolid, PenNone:
	}
	r.ctx.SetLineWidth(o.PenWidth())
}

func (r *ImageRenderer) EndPage(ctx context.Context, job *Job) error {
	var buf bytes.Buffer
	switch {
	case r.isPNG(job):
		if err := r.ctx.EncodePNG(&buf); err != nil {
			return err
		}
	case r.isJPG(job):
		if err := r.encodeJPG(&buf); err != nil {
			return err
		}
	}
	job.SetOutputData(buf.Bytes())
	job.SetOutputDataPosition(uint(len(buf.Bytes())))

	filename := job.OutputFileName()
	if filename != "" {
		switch {
		case r.isPNG(job):
			if err := r.ctx.SavePNG(filename); err != nil {
				return err
			}
		case r.isJPG(job):
			if err := r.saveJPG(filename); err != nil {
				return err
			}
		}
	}
	return nil
}

func (r *ImageRenderer) TextSpan(ctx context.Context, job *Job, p *PointFloat, span *TextSpan) error {
	r.ctx.Push()
	defer r.ctx.Pop()

	rgba := job.Object().PenColor().RGBAUint()
	r.ctx.SetRGB(float64(rgba[0])/255.0, float64(rgba[1])/255.0, float64(rgba[2])/255.0)

	face, err := r.fontFace(r.toX(job, span.Font().Size()))
	if err != nil {
		return err
	}

	p.SetX(r.toX(job, p.X()))
	switch span.Just() {
	case 'r':
		p.SetX(p.X() - r.toX(job, span.Size().X()))
	case 'l':
		// skip
	case 'n':
		p.SetX(p.X() - r.toX(job, span.Size().X()/2.0))
	}
	r.ctx.SetFontFace(face)
	y := r.toY(job, p.Y()+span.YOffsetCenterLine()+span.YOffsetLayout())
	r.ctx.DrawStringAnchored(span.Text(), p.X(), -y, 0, 0)
	return nil
}

func (r *ImageRenderer) Ellipse(ctx context.Context, job *Job, p []*PointFloat, filled bool) error {
	r.ctx.Push()
	defer r.ctx.Pop()
	r.setPenStyle(job)
	rx := r.toX(job, p[1].X()-p[0].X())
	ry := r.toY(job, p[1].Y()-p[0].Y())
	var c *Color
	if filled {
		c = job.Object().FillColor()
		r.ctx.FillPreserve()
	} else {
		c = job.Object().PenColor()
	}
	rgba := c.RGBAUint()
	r.ctx.SetRGB(float64(rgba[0])/255.0, float64(rgba[1])/255.0, float64(rgba[2])/255.0)
	r.ctx.DrawEllipse(r.toX(job, p[0].X()), r.toY(job, -p[0].Y()), rx, ry)
	if filled {
		r.ctx.Fill()
	} else {
		r.ctx.Stroke()
	}
	return nil
}

func (r *ImageRenderer) Polygon(ctx context.Context, job *Job, a []*PointFloat, filled bool) error {
	r.ctx.Push()
	defer r.ctx.Pop()
	r.setPenStyle(job)
	var c *Color
	if filled {
		c = job.Object().FillColor()
	} else {
		c = job.Object().PenColor()
	}
	rgba := c.RGBAUint()
	r.ctx.SetRGB(float64(rgba[0])/255.0, float64(rgba[1])/255.0, float64(rgba[2])/255.0)
	r.ctx.MoveTo(r.toX(job, a[0].X()), r.toY(job, -a[0].Y()))
	for i := 1; i < len(a); i++ {
		r.ctx.LineTo(r.toX(job, a[i].X()), r.toY(job, -a[i].Y()))
	}
	r.ctx.ClosePath()
	if filled {
		r.ctx.Fill()
	} else {
		r.ctx.Stroke()
	}
	return nil
}

func (r *ImageRenderer) Polyline(ctx context.Context, job *Job, a []*PointFloat) error {
	r.ctx.Push()
	defer r.ctx.Pop()
	r.setPenStyle(job)
	rgba := job.Object().PenColor().RGBAUint()
	r.ctx.SetRGB(float64(rgba[0])/255.0, float64(rgba[1])/255.0, float64(rgba[2])/255.0)
	r.ctx.MoveTo(r.toX(job, a[0].X()), r.toY(job, -a[0].Y()))
	for i := 1; i < len(a); i++ {
		r.ctx.LineTo(r.toX(job, a[i].X()), r.toY(job, -a[i].Y()))
	}
	r.ctx.Stroke()
	return nil
}

func (r *ImageRenderer) BezierCurve(ctx context.Context, job *Job, a []*PointFloat, filled bool) error {
	r.ctx.Push()
	defer r.ctx.Pop()
	r.setPenStyle(job)
	var c *Color
	if filled {
		c = job.Object().FillColor()
		r.ctx.FillPreserve()
	} else {
		c = job.Object().PenColor()
	}
	rgba := c.RGBAUint()
	r.ctx.SetRGB(float64(rgba[0])/255.0, float64(rgba[1])/255.0, float64(rgba[2])/255.0)
	r.ctx.MoveTo(r.toX(job, a[0].X()), r.toY(job, -a[0].Y()))
	for i := 1; i < len(a); i += 3 {
		r.ctx.CubicTo(
			r.toX(job, a[i].X()),
			r.toY(job, -a[i].Y()),
			r.toX(job, a[i+1].X()),
			r.toY(job, -a[i+1].Y()),
			r.toX(job, a[i+2].X()),
			r.toY(job, -a[i+2].Y()),
		)
	}
	if filled {
		r.ctx.Fill()
	} else {
		r.ctx.Stroke()
	}
	return nil
}

var (
	fontFaceFn = func(size float64) (font.Face, error) {
		ft, err := truetype.Parse(goregular.TTF)
		if err != nil {
			return nil, err
		}
		opt := &truetype.Options{
			Size:              size,
			DPI:               0,
			Hinting:           0,
			GlyphCacheEntries: 0,
			SubPixelsX:        0,
			SubPixelsY:        0,
		}
		return truetype.NewFace(ft, opt), nil
	}
)

func SetFontFace(fn func(size float64) (font.Face, error)) {
	fontFaceFn = fn
}
