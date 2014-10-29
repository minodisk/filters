package filters

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"testing"
)

func testColorMatrixFilter(mat ColorMatrix, name string, t *testing.T) {
	file, err := os.Open("lenna.png")
	if err != nil {
		t.Fatalf("File doesn't exist")
	}
	defer file.Close()

	src, err := png.Decode(file)
	if err != nil {
		t.Fatalf("Can't decode")
	}

	dst := image.NewRGBA(src.Bounds())
	Transform(dst, src, mat)

	dstFile, err := os.OpenFile(fmt.Sprintf("lenna_cmf_%s.png", name), os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		t.Fatalf("File can't be created")
	}
	defer dstFile.Close()

	if err := png.Encode(dstFile, dst); err != nil {
		t.Fatalf("Can't encode")
	}
}

func TestBrightness(t *testing.T) {
	testColorMatrixFilter(Brightness(0.5), "brightness_0.5", t)
	testColorMatrixFilter(Brightness(-0.5), "brightness_-0.5", t)
}

func TestContrast(t *testing.T) {
	testColorMatrixFilter(Contrast(0.5), "contrast_0.5", t)
	testColorMatrixFilter(Contrast(-0.5), "contrast_-0.5", t)
}

func TestSaturation(t *testing.T) {
	testColorMatrixFilter(Saturation(0.5), "saturation_0.5", t)
	testColorMatrixFilter(Saturation(-0.5), "saturation_-0.5", t)
}

func TestSepia(t *testing.T) {
	testColorMatrixFilter(ColorMatrix{
		0.5, 0.5, 0.5, 0, 0,
		0.33, 0.33, 0.33, 0, 0,
		0.25, 0.25, 0.25, 0, 0,
		0, 0, 0, 1, 0}, "sepia", t)
}

func TestGrayScale(t *testing.T) {
	testColorMatrixFilter(ColorMatrix{
		0.3, 0.59, 0.11, 0, 0,
		0.3, 0.59, 0.11, 0, 0,
		0.3, 0.59, 0.11, 0, 0,
		0, 0, 0, 1, 0}, "gray_scale", t)
}

func TestReverseRB(t *testing.T) {
	testColorMatrixFilter(ColorMatrix{
		0, 0, 1, 0, 0,
		0, 1, 0, 0, 0,
		1, 0, 0, 0, 0,
		0, 0, 0, 1, 0}, "reverse_rb", t)
}

func TestNegative(t *testing.T) {
	testColorMatrixFilter(ColorMatrix{
		-1, 0, 0, 0, 255,
		0, -1, 0, 0, 255,
		0, 0, -1, 0, 255,
		0, 0, 0, 1, 0}, "negative", t)
}
