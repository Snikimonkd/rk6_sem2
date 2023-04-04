package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"

	basiccso "cso/basic_cso"
	improvedcso "cso/improved_cso"

	"cso/utils"
)

func dist(x []float64) float64 {
	res := 0.
	for i := 0; i < len(x); i++ {
		res += x[i] * x[i]
	}
	return math.Sqrt(res)
}

func main() {
	basicG := make([]utils.Graph, 100, 100)
	impG := make([]utils.Graph, 100, 100)

	rand.Seed(time.Now().Unix())
	rasm := []int{2}
	f := utils.Bil
	for j := 0; j < len(rasm); j++ {
		fBestBasic := 1000.
		fAvgBasic := 0.
		xBestBasic := 1000.
		xAvgBasic := 0.
		cBasic := 0

		fBestImp := 1000.
		fAvgImp := 0.
		xBestImp := 1000.
		xAvgImp := 0.
		cImp := 0

		for i := 0; i < 100; i++ {
			sworm := basiccso.NewCatSworm(50, rasm[j], 4.5, f, 0.2, 10, 0.05, 0.5)
			args, graph := sworm.Optimize(100)
			basicG = utils.Add(basicG, graph)

			res := f(args)
			fAvgBasic += res
			if res < fBestBasic {
				fBestBasic = res
			}

			dis := dist(args)
			xAvgBasic += dis
			if dis < xBestBasic {
				xBestBasic = dis
			}

			if res < 0.01 {
				cBasic++
			}

			improvedSworm := improvedcso.NewCatSworm(50, rasm[j], 4.5, f, 0.2, 10, 0.1, 0.5)
			args, graph = improvedSworm.Optimize(100, 0.2)
			impG = utils.Add(impG, graph)

			res = f(args)
			fAvgImp += res
			if res < fBestImp {
				fBestImp = res
			}

			dis = dist(args)
			xAvgImp += dis
			if dis < xBestImp {
				xBestImp = dis
			}

			if res < 0.01 {
				cImp++
			}
		}

		fmt.Printf("Размерность: %d\n", rasm[j])
		fmt.Printf("F лучшее: %e\n", fBestBasic)
		fmt.Printf("F среднее: %e\n", fAvgBasic/100)
		fmt.Printf("X лучшее: %e\n", xBestBasic)
		fmt.Printf("X среднее: %e\n", xAvgBasic/100)
		fmt.Printf("Вероятность: %f\n", float64(cBasic)/100.)
		fmt.Println("----------------------------------")

		fmt.Printf("Размерность: %d\n", rasm[j])
		fmt.Printf("F лучшее: %e\n", fBestImp)
		fmt.Printf("F среднее: %e\n", fAvgImp/100)
		fmt.Printf("X лучшее: %e\n", xBestImp)
		fmt.Printf("X среднее: %e\n", xAvgImp/100)
		fmt.Printf("Вероятность: %f\n", float64(cImp)/100.)
		fmt.Println("----------------------------------")
	}

	fb, err := os.Create("basic")
	if err != nil {
		fmt.Printf("can't create file: %v\n", err)
	}

	for i := range basicG {
		basicG[i].F = basicG[i].F / 100.
	}
	utils.Plot(fb, basicG)

	fi, err := os.Create("improved")
	if err != nil {
		fmt.Printf("can't create file: %v\n", err)
	}
	for i := range impG {
		impG[i].F = impG[i].F / 100.
	}
	utils.Plot(fi, impG)

	fb.Close()
	fi.Close()
}
