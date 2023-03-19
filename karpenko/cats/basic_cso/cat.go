package basiccso

import (
	ch "cso/utils"
	"math"
	"math/rand"
)

// Mode - поведение кошки
type Mode int

const (
	// TracingMode - режим выслеживания
	TracingMode Mode = iota
	// SeekingMode - режим поиска лучшей позиции
	SeekingMode
)

// Cat - кошка
type Cat struct {
	Mode Mode

	// Coordinates - позиция кошки
	Coordinates []float64
	// P - вероятность выбора кошки
	P float64
	// FS (fitness value) - значение функции в данной позиции
	FS float64

	// Переменные для seeking mode
	// SMP (seeking memory pool) - количество копий кошки в seeking mode
	SMP int
	// SRD (seeking range of the selected dimension) - границы поиска
	SRD float64
	// CDC (counts of dimensions to change) - количество направлений, которые надо мутировать
	CDC int

	// Переменные для chasing mode
	// Velosities - скорость кошки в каждом направлении
	Velosities []float64
	// MaxVelocity - максимальная скорость кошки
	MaxVelocity float64
}

func genRandMinus() float64 {
	r := rand.Float32()
	if r < 0.5 {
		return -1.
	}

	return 1.
}

func genRandVelocities(dimensions int, srd float64) []float64 {
	velocities := make([]float64, 0, dimensions)
	for i := 0; i < dimensions; i++ {
		randVelocity := rand.Float64() * srd * genRandMinus()
		velocities = append(velocities, randVelocity)
	}

	return velocities
}

func NewCat(coordinates []float64, smp int, srd float64, cdc int, maxVelocity float64, mode Mode) Cat {
	return Cat{
		Coordinates: coordinates,
		SMP:         smp,
		SRD:         srd,
		CDC:         cdc,
		Velosities:  genRandVelocities(len(coordinates), srd),
		MaxVelocity: maxVelocity,
		Mode:        mode,
	}
}

func (c Cat) Seek(f FitnessFunction) {
	clones := make([]Cat, 0, c.SMP+1)
	clones = append(clones, c)

	c.FS = f(c.Coordinates)

	fsMax := c.FS
	fsMin := c.FS
	for i := 0; i < c.SMP; i++ {
		newCat := NewCat(c.Coordinates, c.SMP, c.SRD, c.CDC, c.MaxVelocity, SeekingMode)
		newCat.move()
		newCat.FS = f(newCat.Coordinates)
		if newCat.FS > fsMax {
			fsMax = newCat.FS
		}
		if newCat.FS < fsMin {
			fsMin = newCat.FS
		}

		clones = append(clones, newCat)
	}

	choices := make([]ch.Choice[Cat], 0, c.SMP+1)
	for i := range clones {
		clones[i].P = math.Abs(clones[i].FS-fsMax) / (fsMax - fsMin)
		choices = append(choices, ch.Choice[Cat]{Item: clones[i], Weight: c.P})
	}

	c = ch.Choose(choices)
}

func isNegative(in float64) bool {
	if in < 0 {
		return true
	}

	return false
}

func (c Cat) Trace(bestCat Cat, f FitnessFunction) {
	// какаято рандомная константа
	C := 1.
	// рандомное число [0,1]
	R := rand.Float64()
	for i := range c.Velosities {
		c.Velosities[i] = c.Velosities[i] + C*R*(bestCat.Coordinates[i]-c.Coordinates[i])
		// если скорость больше максимальной - ставим значение равное максимальному
		if math.Abs(c.Velosities[i]) > c.MaxVelocity {
			if isNegative(c.Velosities[i]) {
				c.Velosities[i] = -c.MaxVelocity
			} else {
				c.Velosities[i] = c.MaxVelocity
			}
		}

		// меняем позицию кошки
	}

	c.move()
	c.FS = f(c.Coordinates)
}

func (c Cat) move() {
	for i := range c.Coordinates {
		c.Coordinates[i] += c.Velosities[i]
	}
}
