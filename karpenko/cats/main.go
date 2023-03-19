package main

import (
	basiccso "cso/basic_cso"
	"fmt"
)

func main() {
	var f basiccso.FitnessFunction = func(args []float64) float64 {
		return (args[0] + 2) * (args[0] + 2)
	}

	sworm := basiccso.NewCatSworm(2, 1, 5, f, 0.2, 20, 2, 1, 1)
	sworm.Print()
	args := sworm.Optimize(1000)
	fmt.Println(args)
}
