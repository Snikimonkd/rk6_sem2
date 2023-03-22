package main

import (
	basiccso "cso/basic_cso"
	"fmt"
)

func main() {
	var f basiccso.FitnessFunction = func(args []float64) float64 {
		return 0.26*(args[0]*args[0]+args[1]*args[1]) - 0.48*args[0]*args[1]
	}

	sworm := basiccso.NewCatSworm(100, 2, 5, f, 0.2, 20, 2, 1, 1)
	args := sworm.Optimize(1000)
	fmt.Println(args)
}
