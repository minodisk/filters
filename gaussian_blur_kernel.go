package filters

import (
	"math"

	"code.google.com/p/graphics-go/graphics/convolve"
)

type GaussianBlurKernel struct {
	convolve.SeparableKernel
}

func NewGaussianBlurKernel(sigma float64, size int) GaussianBlurKernel {
	if sigma == 0 {
		sigma = 0.5
	}
	if size < 1 {
		size = int(math.Ceil(sigma))
	}

	kernel := make([]float64, size*2+1)
	for i := 0; i <= size; i++ {
		x := float64(i) / sigma
		x = math.Pow(1/math.SqrtE, x*x)
		kernel[size-i] = x
		kernel[size+i] = x
	}

	Normalize(&kernel)

	k := GaussianBlurKernel{}
	k.SeparableKernel.X = kernel
	k.SeparableKernel.Y = kernel
	return k
}
