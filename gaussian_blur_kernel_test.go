package filters

import (
	"image"
	"image/png"
	"os"
	"testing"

	"code.google.com/p/graphics-go/graphics/convolve"
)

func TestGaussianBlurKernel(t *testing.T) {
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

	filter := NewGaussianBlurKernel(1.0, 2)
	// log.Println(filter.Weights())

	if err := convolve.Convolve(dst, src, &filter); err != nil {
		t.Errorf("%s", err)
	}

	dstFile, err := os.OpenFile("lenna_gaussian_blur.png", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		t.Fatalf("File can't be created")
	}
	defer dstFile.Close()

	if err := png.Encode(dstFile, dst); err != nil {
		t.Fatalf("Can't encode")
	}
}
