package improvedcso

import (
	"math"
	"math/rand"

	ch "cso/utils"
)

// Mode - поведение кошки
type Mode string

const (
	// TracingMode - режим выслеживания
	TracingMode Mode = "seeking"
	// SeekingMode - режим поиска лучшей позиции
	SeekingMode Mode = "tracing"
)

// Cat - кошка
type Cat struct {
	// Mode - режим кошки
	Mode Mode

	// Coordinates - позиция кошки
	Coordinates []float64
	// FS (fitness value) - значение функции в данной позиции
	FS float64

	// Переменные для seeking mode
	// ClonesAmount - количество копий кошки в seeking mode
	ClonesAmount int
	// MaxSeekingSpeed - максимальная скорость в seeking mode
	MaxSeekingSpeed float64

	// Переменные для chasing mode
	// Velosities - скорость кошки в каждом направлении
	Velosities []float64
	// MaxVelocity - максимальная скорость кошки
	MaxVelocity float64

	//
	zatup int
}

func genRandMinus() float64 {
	r := rand.Float32()
	if r < 0.5 {
		return -1.
	}

	return 1.
}

func genRandVelocities(dimensions int, maxSeekingVelosicty float64) []float64 {
	velocities := make([]float64, 0, dimensions)
	for i := 0; i < dimensions; i++ {
		randVelocity := rand.Float64() * maxSeekingVelosicty * genRandMinus()
		velocities = append(velocities, randVelocity)
	}

	return velocities
}

func NewCat(coordinates []float64, clonesAmount int, maxSeekingSpeed float64, maxVelocity float64, mode Mode) Cat {
	cords := make([]float64, len(coordinates))
	copy(cords, coordinates)
	return Cat{
		Coordinates:     cords,
		ClonesAmount:    clonesAmount,
		MaxSeekingSpeed: maxSeekingSpeed,
		Velosities:      genRandVelocities(len(coordinates), maxSeekingSpeed),
		MaxVelocity:     maxVelocity,
		Mode:            mode,
	}
}

func newSeekingSpeed(args []float64) float64 {
	res := 0.
	for i := 0; i < len(args); i++ {
		res += args[i] * args[i]
	}

	return math.Sqrt(res)
}

func (c Cat) Seek(f ch.FitnessFunction, isBest bool, maxSeekingSpeed float64) Cat {
	clones := make([]Cat, 0, c.ClonesAmount)
	clones = append(clones, c)

	fsMin := c.FS
	ind := 0

	for i := 1; i < c.ClonesAmount; i++ {
		newCat := NewCat(c.Coordinates, c.ClonesAmount, c.MaxSeekingSpeed, c.MaxVelocity, SeekingMode)
		newCat = newCat.move()
		newCat.FS = f(newCat.Coordinates)
		if newCat.FS < fsMin {
			fsMin = newCat.FS
			ind = i
		}

		clones = append(clones, newCat)
	}

	if isBest && ind == 0 {
		clones[ind].MaxSeekingSpeed = clones[ind].MaxSeekingSpeed * 0.5
	}

	// if isBest && ind != 0 {
	// 	clones[ind].MaxSeekingSpeed = maxSeekingSpeed
	// }

	return clones[ind]
}

func isNegative(in float64) bool {
	if in < 0 {
		return true
	}

	return false
}

func (c Cat) Trace(bestCat Cat, f ch.FitnessFunction) Cat {
	// какаято рандомная константа
	C := 0.5
	for i := range c.Velosities {
		R := rand.Float64()
		c.Velosities[i] = c.Velosities[i] + C*R*(bestCat.Coordinates[i]-c.Coordinates[i])
		// если скорость больше максимальной - ставим значение равное максимальному
		if math.Abs(c.Velosities[i]) > c.MaxVelocity {
			if isNegative(c.Velosities[i]) {
				c.Velosities[i] = -c.MaxVelocity
			} else {
				c.Velosities[i] = c.MaxVelocity
			}
		}
	}

	c = c.move()
	c.FS = f(c.Coordinates)

	return c
}

func (c Cat) move() Cat {
	for i := range c.Coordinates {
		c.Coordinates[i] += c.Velosities[i]
	}

	return c
}
