package utils

import "math/rand"

// Choice is a generic wrapper that can be used to add weights for any item.
type Choice[T any] struct {
	Item   T
	Weight float64
	Sum    float64
}

func Choose[T any](in []Choice[T]) T {
	sum := 0.

	for i := 0; i < len(in); i++ {
		sum += in[i].Weight
	}

	in[0].Sum = in[0].Weight / sum
	for i := 1; i < len(in); i++ {
		in[i].Sum = in[i].Weight/sum + in[i-1].Sum
	}

	r := randFloat64N(sum)
	bucket := 0
	for r > in[bucket].Sum {
		bucket++
	}

	return in[bucket].Item
}

func randFloat64N(n float64) float64 {
	return rand.Float64() * n
}
