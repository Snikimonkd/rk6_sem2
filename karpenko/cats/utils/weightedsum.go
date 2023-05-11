package utils

import (
	"fmt"
	"math"
	"math/rand"
	"os"
)

type Graph struct {
	F  float64
	It int
}

func Add(l []Graph, r []Graph) []Graph {
	for i := range l {
		l[i].F += r[i].F
		l[i].It = i
	}

	return l
}

func Plot(f *os.File, graph []Graph) {
	prev := graph[0].F
	it := 0
	for _, v := range graph {
		if math.Abs(prev-v.F) < 0.000001 {
			it++
		} else {
			prev = v.F
			it = 0
		}
		if it == 20 {
			return
		}
		_, err := f.WriteString(fmt.Sprintf("%d\t%E\n", v.It, v.F))
		if err != nil {
			fmt.Printf("can't write to file: %v\n", err)
		}
	}
}

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

	//r := randFloat64N(sum)
	r := rand.Float64()
	bucket := 0
	for r > in[bucket].Sum {
		bucket++
	}

	return in[bucket].Item
}

func randFloat64N(n float64) float64 {
	return rand.Float64() * n
}
