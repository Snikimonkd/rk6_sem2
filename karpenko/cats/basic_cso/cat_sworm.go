package basiccso

import (
	"fmt"
	"math"
	"math/rand"

	ch "cso/utils"
)

// FitnessFunction - функция приспособленности
type FitnessFunction func(args []float64) float64

type CatSworm struct {
	// MR (mixed ratio) - соотношение кошек в seeking mode и tracing mode
	MR float64
	// Cats - кошки
	Cats []Cat
	// F - функция
	F FitnessFunction
}

func genRandCoordinates(dimensions int, area float64) []float64 {
	ret := make([]float64, 0, dimensions)
	for i := 0; i < dimensions; i++ {
		ret = append(ret, math.Mod(rand.NormFloat64(), area))
	}

	return ret
}

func NewCatSworm(amount int, dimensions int, area float64, f FitnessFunction, mr float64, smp int, srd float64, cdc int, maxVelocity float64) CatSworm {
	cats := make([]Cat, 0, amount)

	choices := []ch.Choice[Mode]{
		{Item: SeekingMode, Weight: mr},
		{Item: TracingMode, Weight: 1 - mr},
	}

	for i := 0; i < amount; i++ {
		mode := ch.Choose(choices)
		cords := genRandCoordinates(dimensions, area)

		cat := NewCat(cords, smp, srd, cdc, maxVelocity, mode)
		cat.FS = f(cat.Coordinates)
		cats = append(cats, cat)
	}

	ret := CatSworm{
		MR:   mr,
		Cats: cats,
		F:    f,
	}

	return ret
}

func (cs CatSworm) getBestCat() int {
	best := cs.Cats[0].FS
	ind := 0
	for i, v := range cs.Cats {
		if v.FS < best {
			best = v.FS
			ind = i
		}
	}

	return ind
}

func (cs CatSworm) Optimize(it int) []float64 {
	choices := []ch.Choice[Mode]{
		{Item: SeekingMode, Weight: cs.MR},
		{Item: TracingMode, Weight: 1 - cs.MR},
	}

	for i := 0; i < it; i++ {
		bestCat := cs.getBestCat()
		for j := range cs.Cats {
			if j == bestCat {
				continue
			}
			switch cs.Cats[j].Mode {
			case SeekingMode:
				cs.Cats[j].Seek(cs.F)
			case TracingMode:
				cs.Cats[j].Trace(cs.Cats[bestCat], cs.F)
			}
			cs.Cats[j].Mode = ch.Choose(choices)
			cs.Cats[j].FS = cs.F(cs.Cats[j].Coordinates)
		}
	}

	return cs.Cats[cs.getBestCat()].Coordinates
}

func (cs CatSworm) Print() {
	for _, v := range cs.Cats {
		fmt.Println("Mode: ", v.Mode)
		fmt.Println("Cords: ", v.Coordinates)
		fmt.Println("P: ", v.P)
		fmt.Println("FS: ", v.FS)
		fmt.Println("SMP: ", v.SMP)
		fmt.Println("SRD: ", v.SRD)
		fmt.Println("CDC: ", v.CDC)
		fmt.Println("Vel: ", v.Velosities)
		fmt.Println("MaxVelocity: ", v.MaxVelocity)
		fmt.Println("----------------------------")
	}
}
