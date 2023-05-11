package improvedcso

import (
	"fmt"
	"math/rand"

	utils "cso/utils"
)

type CatSworm struct {
	// MR (mixed ratio) - соотношение кошек в seeking mode и tracing mode
	MR float64
	// Cats - кошки
	Cats []Cat
	// F - функция
	F utils.FitnessFunction
}

func genRandCoordinates(dimensions int, area float64) []float64 {
	ret := make([]float64, 0, dimensions)
	for i := 0; i < dimensions; i++ {
		ret = append(ret, genRandMinus()*rand.Float64()*area)
	}

	return ret
}

func NewCatSworm(amount int, dimensions int, area float64, f utils.FitnessFunction, mr float64, clonesAmount int, maxSeekingSpeed float64, maxVelocity float64) CatSworm {
	cats := make([]Cat, 0, amount)

	choices := []utils.Choice[Mode]{
		{Item: SeekingMode, Weight: mr},
		{Item: TracingMode, Weight: 1 - mr},
	}

	for i := 0; i < amount; i++ {
		mode := utils.Choose(choices)
		cords := genRandCoordinates(dimensions, area)

		cat := NewCat(cords, clonesAmount, maxSeekingSpeed, maxVelocity, mode, 0)
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

func (cs CatSworm) Optimize(it int) ([]float64, []utils.Graph) {
	res := make([]utils.Graph, 0, it)
	choices := []utils.Choice[Mode]{
		{Item: SeekingMode, Weight: cs.MR},
		{Item: TracingMode, Weight: 1 - cs.MR},
	}

	for i := 0; i < it; i++ {
		bestCat := cs.getBestCat()
		res = append(res, utils.Graph{F: cs.F(cs.Cats[bestCat].Coordinates), It: i})
		cs.Cats[bestCat].Mode = SeekingMode
		for j := range cs.Cats {
			switch cs.Cats[j].Mode {
			case SeekingMode:
				cs.Cats[j] = cs.Cats[j].Seek(cs.F, j == bestCat, i)
			case TracingMode:
				cs.Cats[j] = cs.Cats[j].Trace(cs.Cats[bestCat], cs.F)
			}
			cs.Cats[j].Mode = utils.Choose(choices)
			cs.Cats[j].FS = cs.F(cs.Cats[j].Coordinates)
		}
	}

	return cs.Cats[cs.getBestCat()].Coordinates, res
}

func (cs CatSworm) Print() {
	for _, v := range cs.Cats {
		fmt.Println("Mode: ", v.Mode)
		fmt.Println("Cords: ", v.Coordinates)
		fmt.Println("FS: ", v.FS)
		fmt.Println("ClonesAmount: ", v.ClonesAmount)
		fmt.Println("MaxSeekingSpeed: ", v.MaxSeekingSpeed)
		fmt.Println("Velosities: ", v.Velosities)
		fmt.Println("MaxVelocity: ", v.MaxVelocity)
		fmt.Println("----------------------------")
	}
}
