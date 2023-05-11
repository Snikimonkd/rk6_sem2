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

func distLevi(arr []float64) float64 {
	x := arr[0]
	y := arr[1]

	return math.Sqrt(math.Pow(x-1., 2) + math.Pow(y-1., 2))

}

func distBut(arr []float64) float64 {
	x := arr[0]
	y := arr[1]

	return math.Sqrt(math.Pow(x-1., 2) + math.Pow(y-3., 2))

}

func distBil(arr []float64) float64 {
	x := arr[0]
	y := arr[1]

	return math.Sqrt(math.Pow(x-3., 2) + math.Pow(y-0.5, 2))
}

func distRozenbrok(x []float64) float64 {
	res := 0.
	for i := 0; i < len(x); i++ {
		res += (x[i] - 1) * (x[i] - 1)
	}
	return math.Sqrt(res)
}

func dist(x []float64) float64 {
	res := 0.
	for i := 0; i < len(x); i++ {
		res += x[i] * x[i]
	}
	return math.Sqrt(res)
}

var f = utils.Rozenbrok

const amount = 50
const area = 4.5
const mr = 0.2
const clonesAmount = 10
const maxSeekingSpeed = 0.5
const maxVelocity = 0.5
const lyambda = 0.01

var rasm = []int{2, 4, 8, 16, 32, 64}

const actual = 0.

func main() {
	for r := 0; r < len(rasm); r++ {
		eps := float64(area*2) * math.Sqrt(math.Sqrt(float64(rasm[r]))) * lyambda
		fmt.Println("eps=", eps)
		fmt.Printf("Размерность: %d\n", rasm[r])

		basicG := make([]utils.Graph, 1000, 1000)
		impG := make([]utils.Graph, 1000, 1000)

		rand.Seed(time.Now().Unix())
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
			sworm := basiccso.NewCatSworm(amount, rasm[r], area, f, mr, clonesAmount, maxSeekingSpeed, maxVelocity)
			args, graph := sworm.Optimize(1000)
			basicG = utils.Add(basicG, graph)

			res := f(args)
			fAvgBasic += res
			if res < fBestBasic {
				fBestBasic = res
			}

			dis := distRozenbrok(args)
			xAvgBasic += dis
			if dis < xBestBasic {
				xBestBasic = dis
			}

			if math.Abs(res-actual) < eps {
				cBasic++
			}

			improvedSworm := improvedcso.NewCatSworm(amount, rasm[r], area, f, mr, clonesAmount, maxSeekingSpeed, maxVelocity)
			args, graph = improvedSworm.Optimize(1000)
			impG = utils.Add(impG, graph)

			res = f(args)
			fAvgImp += res
			if res < fBestImp {
				fBestImp = res
			}

			dis = distRozenbrok(args)
			xAvgImp += dis
			if dis < xBestImp {
				xBestImp = dis
			}

			if math.Abs(res-actual) < eps {
				cImp++
			}
		}

		fmt.Printf("F лучшее: %e\n", fBestBasic)
		fmt.Printf("X лучшее: %e\n", xBestBasic)
		fmt.Printf("F среднее: %e\n", fAvgBasic/100.)
		fmt.Printf("X среднее: %e\n", xAvgBasic/100.)
		fmt.Printf("Вероятность: %f\n", float64(cBasic)/100.)
		fmt.Println("----------------------------------")

		fmt.Printf("F лучшее: %e\n", fBestImp)
		fmt.Printf("X лучшее: %e\n", xBestImp)
		fmt.Printf("F среднее: %e\n", fAvgImp/100.)
		fmt.Printf("X среднее: %e\n", xAvgImp/100.)
		fmt.Printf("Вероятность: %f\n", float64(cImp)/100.)
		fmt.Println("----------------------------------")

		fb, err := os.Create(fmt.Sprintf("basic%d", rasm[r]))
		if err != nil {
			fmt.Printf("can't create file: %v\n", err)
		}
		utils.Plot(fb, basicG)

		fi, err := os.Create(fmt.Sprintf("improved%d", rasm[r]))
		if err != nil {
			fmt.Printf("can't create file: %v\n", err)
		}
		utils.Plot(fi, impG)

		fb.Close()
		fi.Close()

	}

}
