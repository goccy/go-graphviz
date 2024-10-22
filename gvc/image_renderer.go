package gvc

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"os"
	"strings"
	"sync"

	"github.com/flopp/go-findfont"
	"github.com/fogleman/gg"
	"github.com/goccy/go-graphviz/internal/wasm"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/font/sfnt"
)

var (
	fontMu    sync.RWMutex
	fontCache = make(map[string]font.Face)
)

type ImageRenderer struct {
	*DefaultRenderEngine
	ctx *gg.Context
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

	font := span.Font()
	face, err := r.getFontFace(ctx, job, font)
	if face == nil || err != nil {
		defaultFont, err := r.defaultFontFace(ctx, job, font)
		if err != nil {
			return err
		}
		face = defaultFont
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

func (r *ImageRenderer) getFontFace(ctx context.Context, job *Job, font *TextFont) (font.Face, error) {
	return r.lookupFontWithCache(ctx, job, font)
}

func (r *ImageRenderer) lookupFontWithCache(ctx context.Context, job *Job, font *TextFont) (font.Face, error) {
	fontSize := font.Size() * job.Zoom()
	fontName := font.Name()
	cacheKey := fmt.Sprintf("%s:%f", fontName, fontSize)
	fontMu.RLock()
	if font, exists := fontCache[cacheKey]; exists {
		fontMu.RUnlock()
		return font, nil
	}
	fontMu.RUnlock()

	fontLoaderMu.RLock()
	defer fontLoaderMu.RUnlock()

	if fontLoader != nil {
		face, err := fontLoader(ctx, job, font)
		if err != nil {
			return nil, err
		}
		if face != nil {
			return face, nil
		}
	}

	ft, err := r.lookupFont(fontName, fontSize, job.DPI())
	if err != nil {
		return nil, err
	}
	fontMu.Lock()
	fontCache[cacheKey] = ft
	fontMu.Unlock()
	return ft, nil
}

func (r *ImageRenderer) lookupFont(fontName string, fontSize float64, dpi *PointFloat) (font.Face, error) {
	fontPath, err := findfont.Find(fontName)
	if err == nil {
		return r.lookupFontFromTTFFile(fontName, fontSize, dpi, fontPath)
	}
	parts := strings.Split(fontName, "-")
	for i := len(parts) - 1; i > 0; i-- {
		baseName := strings.Join(parts[:len(parts)-1], "-")
		ttfFace, err := r.lookupFontFromTTFFile(fontName, fontSize, dpi, baseName+".ttf")
		if err != nil {
			return nil, err
		}
		if ttfFace != nil {
			return ttfFace, nil
		}
		ttcFace, err := r.lookupFontFromTTCFile(fontName, fontSize, dpi, baseName+".ttc")
		if err != nil {
			return nil, err
		}
		if ttcFace != nil {
			return ttcFace, nil
		}
	}
	return nil, fmt.Errorf("failed to find font by %s", fontName)
}

func (r *ImageRenderer) lookupFontFromTTFFile(fontName string, fontSize float64, dpi *PointFloat, fontPath string) (font.Face, error) {
	fontData, err := os.ReadFile(fontPath)
	if err != nil {
		return nil, nil
	}
	ft, err := truetype.Parse(fontData)
	if err != nil {
		return nil, err
	}
	return truetype.NewFace(ft, &truetype.Options{
		Size: fontSize,
	}), nil
}

func (r *ImageRenderer) lookupFontFromTTCFile(fontName string, fontSize float64, dpi *PointFloat, fontPath string) (font.Face, error) {
	parts := strings.Split(fontName, "-")
	fontPath, err := findfont.Find(fontPath)
	if err != nil {
		return nil, nil
	}
	fontData, err := os.ReadFile(fontPath)
	if err != nil {
		return nil, err
	}
	c, err := opentype.ParseCollection(fontData)
	if err != nil {
		return nil, err
	}
	for j := 0; j < c.NumFonts(); j++ {
		ft, err := c.Font(j)
		if err != nil {
			return nil, err
		}
		var buf sfnt.Buffer
		name, err := ft.Name(&buf, sfnt.NameIDFull)
		if err != nil {
			return nil, err
		}
		if strings.Join(parts, " ") == name {
			return opentype.NewFace(ft, &opentype.FaceOptions{
				Size: fontSize,
				DPI:  dpi.X(),
			})
		}
	}
	return nil, fmt.Errorf("failed to find %s font from %s file", fontName, fontPath)
}

func (r *ImageRenderer) defaultFontFace(ctx context.Context, job *Job, font *TextFont) (font.Face, error) {
	ft, err := truetype.Parse(goregular.TTF)
	if err != nil {
		return nil, err
	}
	return truetype.NewFace(ft, &truetype.Options{
		Size: font.Size() * job.Zoom(),
	}), nil
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

func (r *ImageRenderer) LoadImage(ctx context.Context, job *Job, shape *UserShape, bf *BoxFloat, filled bool) error {
	r.ctx.Push()
	defer r.ctx.Pop()

	fs := wasm.FileSystem()
	f, err := fs.Open(shape.Name())
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	io.Copy(&buf, f)
	img, _, err := image.Decode(&buf)
	if err != nil {
		return err
	}
	r.ctx.DrawImageAnchored(img, int(job.Scale().X()*bf.LL().X()), -int(job.Scale().Y()*bf.LL().Y()), 0, 1)
	return nil
}

type FontLoader func(ctx context.Context, job *Job, font *TextFont) (font.Face, error)

var (
	fontLoaderMu sync.RWMutex
	fontLoader   FontLoader
)

func SetFontLoader(loader FontLoader) {
	fontLoaderMu.Lock()
	defer fontLoaderMu.Unlock()
	fontLoader = loader
}
