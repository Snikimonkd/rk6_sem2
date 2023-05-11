package utils

import "math"

// FitnessFunction - функция приспособленности
type FitnessFunction func(args []float64) float64

// ok {1, 1}
var Levi13 FitnessFunction = func(args []float64) float64 {
	x := args[0]
	y := args[1]
	return math.Pow(math.Sin(3.*math.Pi*x), 2) + (x-1)*(x-1)*(1+math.Pow(math.Sin(3.*math.Pi*y), 2)) + (y-1)*(y-1)*(1+math.Pow(math.Sin(2*math.Pi*y), 2))
}

// ok {0, 0, 0 ...}
var Rastrigin FitnessFunction = func(x []float64) float64 {
	A := 10.
	n := float64(len(x))
	res := A * n
	for i := 0; i < len(x); i++ {
		res += x[i]*x[i] - A*math.Cos(2*math.Pi*x[i])
	}

	return res
}

// ok {0, 0, 0 ...}
var Sphere FitnessFunction = func(args []float64) float64 {
	res := 0.
	for i := 0; i < len(args); i++ {
		res += (args[i]) * (args[i])
	}
	return res
}

// ok {1, 1, 1 ...}
var Rozenbrok FitnessFunction = func(args []float64) float64 {
	res := 0.
	for i := 0; i < len(args)-1; i++ {
		res += (100*math.Pow((args[i+1]-args[i]*args[i]), 2) + (args[i]-1)*(args[i]-1))
	}
	return res
}

// ok {3, 0.5}
var Bil FitnessFunction = func(args []float64) float64 {
	x := args[0]
	y := args[1]
	return math.Pow((1.5-x+x*y), 2) + math.Pow((2.25-x+x*y*y), 2) + math.Pow((2.625-x+x*y*y*y), 2)
}

// ok {1, 3}
var But FitnessFunction = func(args []float64) float64 {
	x := args[0]
	y := args[1]
	return math.Pow((x+2*y-7), 2) + math.Pow((2*x+y-5), 2)
}
