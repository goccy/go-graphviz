package gvc

import (
	"image"
	"image/color"
	"testing"
)

func TestResizeImage(t *testing.T) {
	src := image.NewNRGBA(image.Rect(0, 0, 40, 20))
	for y := 0; y < 20; y++ {
		for x := 0; x < 40; x++ {
			src.Set(x, y, color.NRGBA{R: 200, G: 30, B: 30, A: 255})
		}
	}

	tests := []struct {
		name         string
		width        int
		height       int
		wantW, wantH int
	}{
		{name: "both dimensions", width: 10, height: 10, wantW: 10, wantH: 10},
		{name: "width only keeps the aspect ratio", width: 20, height: 0, wantW: 20, wantH: 10},
		{name: "height only keeps the aspect ratio", width: 0, height: 5, wantW: 10, wantH: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := resizeImage(src, tt.width, tt.height)
			b := got.Bounds()
			if b.Dx() != tt.wantW || b.Dy() != tt.wantH {
				t.Fatalf("size = %dx%d, want %dx%d", b.Dx(), b.Dy(), tt.wantW, tt.wantH)
			}
			c := color.NRGBAModel.Convert(got.At(b.Min.X+b.Dx()/2, b.Min.Y+b.Dy()/2)).(color.NRGBA)
			if c.R < 190 || c.A != 255 {
				t.Fatalf("center pixel = %+v, want the source colour", c)
			}
		})
	}

	if got := resizeImage(src, 0, 0); got != image.Image(src) {
		t.Fatal("a zero size must return the image unchanged")
	}
}
