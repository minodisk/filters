package filters

import (
	"image"
	"image/draw"
)

func clamp(x, x0, x1 float64) float64 {
	if x < x0 {
		return x0
	}
	if x > x1 {
		return x1
	}
	return x
}

type ColorMatrix []float64

// Brightness -1.0 ~ 1.0
func Brightness(val float64) ColorMatrix {
	v := 255 * val
	return ColorMatrix{
		1, 0, 0, 0, v,
		0, 1, 0, 0, v,
		0, 0, 1, 0, v,
		0, 0, 0, 1, 0}
}

// Contrast -1.0 ~ 1.0
func Contrast(val float64) ColorMatrix {
	s := val + 1
	o := 128 * (1 - s)
	return ColorMatrix{
		s, 0, 0, 0, o,
		0, s, 0, 0, o,
		0, 0, s, 0, o,
		0, 0, 0, 1, 0}
}

// Saturation -1.0 ~ 1.0
func Saturation(val float64) ColorMatrix {
	lumaR := 0.212671
	lumaG := 0.71516
	lumaB := 0.072169
	i := -val
	v := val + 1
	r := i * lumaR
	g := i * lumaG
	b := i * lumaB
	return ColorMatrix{
		r + v, g, b, 0, 0,
		r, g + v, b, 0, 0,
		r, g, b + v, 0, 0,
		0, 0, 0, 1, 0}
}

func Transform(dst draw.Image, src image.Image, mat ColorMatrix) error {
	if dst == nil || src == nil || mat == nil {
		return nil
	}

	b := dst.Bounds()
	dstRgba, ok := dst.(*image.RGBA)
	if !ok {
		dstRgba = image.NewRGBA(b)
	}

	if err := transform(dstRgba, src, mat); err != nil {
		return err
	}

	if !ok {
		draw.Draw(dst, b, dstRgba, b.Min, draw.Src)
	}
	return nil
}

func transform(dst *image.RGBA, src image.Image, mat ColorMatrix) error {
	bounds := dst.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			sr, sg, sb, sa := src.At(x, y).RGBA()
			r := float64(sr >> 8)
			g := float64(sg >> 8)
			b := float64(sb >> 8)
			a := float64(sa >> 8)
			off := (y-dst.Rect.Min.Y)*dst.Stride + (x-dst.Rect.Min.X)*4
			dst.Pix[off+0] = uint8(clamp(mat[0]*r+mat[1]*g+mat[2]*b+mat[3]*a+mat[4], 0, 255))
			dst.Pix[off+1] = uint8(clamp(mat[5]*r+mat[6]*g+mat[7]*b+mat[8]*a+mat[9], 0, 255))
			dst.Pix[off+2] = uint8(clamp(mat[10]*r+mat[11]*g+mat[12]*b+mat[13]*a+mat[14], 0, 255))
			dst.Pix[off+3] = uint8(clamp(mat[15]*r+mat[16]*g+mat[17]*b+mat[18]*a+mat[19], 0, 255))
		}
	}
	return nil
}
