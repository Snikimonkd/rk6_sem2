package main

import (
	"fmt"
)

type Figure string

const BLACK Figure = "BLACK"
const WHITE Figure = "WHITE"
const EMPTY Figure = "EMPTY"

type Cord struct {
	i, j int
}

var final = [][]Figure{
	{WHITE, WHITE, WHITE},
	{EMPTY, BLACK, WHITE},
	{EMPTY, WHITE, BLACK},
	{BLACK, BLACK, BLACK},
}

var start = [][]Figure{
	{BLACK, BLACK, BLACK},
	{EMPTY, WHITE, BLACK},
	{EMPTY, BLACK, WHITE},
	{WHITE, WHITE, WHITE},
}

var final2 = [][]Figure{
	{EMPTY, WHITE, WHITE, WHITE},
	{EMPTY, EMPTY, BLACK, WHITE},
	{EMPTY, EMPTY, WHITE, BLACK},
	{EMPTY, BLACK, BLACK, BLACK},
}

var start2 = [][]Figure{
	{EMPTY, BLACK, BLACK, BLACK},
	{EMPTY, EMPTY, WHITE, BLACK},
	{EMPTY, EMPTY, BLACK, WHITE},
	{EMPTY, WHITE, WHITE, WHITE},
}

var final3 = [][]Figure{
	{WHITE, WHITE, WHITE},
	{EMPTY, WHITE, WHITE},
	{EMPTY, BLACK, BLACK},
	{BLACK, BLACK, BLACK},
}

var start3 = [][]Figure{
	{BLACK, BLACK, BLACK},
	{EMPTY, BLACK, BLACK},
	{EMPTY, WHITE, WHITE},
	{WHITE, WHITE, WHITE},
}

var final4 = [][]Figure{
	{WHITE, WHITE, WHITE},
	{EMPTY, BLACK, BLACK},
	{EMPTY, WHITE, WHITE},
	{BLACK, BLACK, BLACK},
}

var start4 = [][]Figure{
	{BLACK, BLACK, BLACK},
	{EMPTY, WHITE, WHITE},
	{EMPTY, BLACK, BLACK},
	{WHITE, WHITE, WHITE},
}

const max = 12

func newPos(cords Cord, n, m int) []Cord {
	ret := make([]Cord, 0, 8)
	// вниз вправо
	newCords := Cord{
		i: cords.i + 2,
		j: cords.j + 1,
	}

	if newCords.i < n && newCords.j < m {
		ret = append(ret, newCords)
	}

	// вниз влево
	newCords = Cord{
		i: cords.i + 2,
		j: cords.j - 1,
	}

	if newCords.i < n && newCords.j >= 0 {
		ret = append(ret, newCords)
	}

	// влево вниз
	newCords = Cord{
		i: cords.i + 1,
		j: cords.j - 2,
	}

	if newCords.i < n && newCords.j >= 0 {
		ret = append(ret, newCords)
	}

	// влево вверх
	newCords = Cord{
		i: cords.i - 1,
		j: cords.j - 2,
	}

	if newCords.i >= 0 && newCords.j >= 0 {
		ret = append(ret, newCords)
	}

	// вверх вправо
	newCords = Cord{
		i: cords.i - 2,
		j: cords.j + 1,
	}

	if newCords.i >= 0 && newCords.j < m {
		ret = append(ret, newCords)
	}

	// вверх влево
	newCords = Cord{
		i: cords.i - 2,
		j: cords.j - 1,
	}

	if newCords.i >= 0 && newCords.j >= 0 {
		ret = append(ret, newCords)
	}

	// вправо вниз
	newCords = Cord{
		i: cords.i + 1,
		j: cords.j + 2,
	}

	if newCords.i < n && newCords.j < m {
		ret = append(ret, newCords)
	}

	// вправо вверх
	newCords = Cord{
		i: cords.i - 1,
		j: cords.j + 2,
	}

	if newCords.i >= 0 && newCords.j < m {
		ret = append(ret, newCords)
	}

	return ret
}

func copy(state [][]Figure) [][]Figure {
	newState := make([][]Figure, len(state), len(state))
	for i := 0; i < len(newState); i++ {
		newState[i] = make([]Figure, len(state[i]), len(state[i]))
	}

	for i := 0; i < len(state); i++ {
		for j := 0; j < len(state[1]); j++ {
			newState[i][j] = state[i][j]
		}
	}

	return newState
}

func applyMove(state [][]Figure, oldPos, newPos Cord) [][]Figure {
	cp := copy(state)
	if cp[newPos.i][newPos.j] != EMPTY {
		return nil
	}
	figure := cp[oldPos.i][oldPos.j]
	cp[newPos.i][newPos.j] = figure
	cp[oldPos.i][oldPos.j] = EMPTY

	return cp
}

func genVariants(state [][]Figure) [][][]Figure {
	res := make([][][]Figure, 0, 8*len(state)*len(state[0]))
	for i := 0; i < len(state); i++ {
		for j := 0; j < len(state[0]); j++ {
			if state[i][j] != EMPTY {
				cords := newPos(Cord{i, j}, len(state), len(state[0]))
				for k := 0; k < len(cords); k++ {
					newState := applyMove(state, Cord{i, j}, cords[k])
					if newState != nil {
						res = append(res, newState)
					}
				}
			}
		}
	}

	return res
}

func Print(state [][]Figure) {
	for i := range state {
		for j := range state[i] {
			switch state[i][j] {
			case BLACK:
				fmt.Print("🎠|")
			case WHITE:
				fmt.Print("🐎|")
			case EMPTY:
				fmt.Print("  |")
			}
		}
		fmt.Print("\n")
	}
}

func statesEqual(left, right step) bool {
	if left.short == right.short {
		return true
	}

	return false
}

func RemoveIndex(s [][][]Figure, index int) [][][]Figure {
	return append(s[:index], s[index+1:]...)
}

func heuristic(state [][]Figure) int {
	ret := 0
	for i := 0; i < len(final); i++ {
		for j := 0; j < len(final[0]); j++ {
			if state[i][j] != final[i][j] {
				ret++
			}

		}
	}

	return ret
}

type step struct {
	state [][]Figure
	// эвристика
	h int
	// стоимость
	g int
	// приоритет
	f int

	short string

	parent *step
}

func short(state [][]Figure) string {
	res := ""
	for i := 0; i < len(state); i++ {
		for j := 0; j < len(state[0]); j++ {
			switch state[i][j] {
			case BLACK:
				res += "b"
			case WHITE:
				res += "w"
			case EMPTY:
				res += "e"
			}
		}
	}
	return res
}

func newStep(state [][]Figure, h, g int, parent *step) step {
	return step{
		state:  state,
		h:      h,
		g:      g,
		f:      h + g,
		parent: parent,
		short:  short(state),
	}
}

func Pop(queue []step) (step, []step) {
	element := queue[0]
	if len(queue) == 1 {
		queue = make([]step, 0)
		return element, queue
	}

	ind := 0
	for i := 1; i < len(queue); i++ {
		if queue[i].f < element.f {
			element = queue[i]
			ind = i
		}
	}

	if ind == len(queue)-1 {
		queue = queue[:len(queue)-1]
	} else {
		queue = append(queue[:ind], queue[ind+1:]...)
	}

	return element, queue
}

func printSteps(in []step) {
	for i := 0; i < len(in); i++ {
		Print(in[i].state)
	}
}

func cleanAmongAllVariants(closed []step, opened []step, newVars []step) []step {
	var ind []int
	for j := 0; j < len(closed); j++ {
		for i := 0; i < len(newVars); i++ {
			if newVars[i].short == closed[j].short {
				ind = append(ind, i)
			}
		}
	}

	for j := 0; j < len(opened); j++ {
		for i := 0; i < len(newVars); i++ {
			if newVars[i].short == opened[j].short {
				ind = append(ind, i)
			}
		}
	}

	res := make([]step, 0, len(newVars)-len(ind))
	for i := 0; i < len(newVars); i++ {
		flag := false
		for j := 0; j < len(ind); j++ {
			if i == ind[j] {
				flag = true
			}
		}
		if !flag {
			res = append(res, newVars[i])
		}
	}

	return res
}

func way(end *step) {
	arr := []step{}
	for end != nil {
		arr = append(arr, *end)
		end = end.parent
	}

	iter := 0
	for i := len(arr) - 1; i >= 0; i-- {
		fmt.Println(iter)
		iter++
		printSteps([]step{arr[i]})
		fmt.Println("---------------------")
	}
}

func aStar(start [][]Figure) step {
	// reader := bufio.NewReader(os.Stdin)
	opened := []step{newStep(start, heuristic(start), 0, nil)}
	closed := []step{}

	finale := newStep(final, 0, 0, nil)

	n := len(opened)
	iter := 0
	for n > 0 {
		iter++
		if iter%100 == 0 {
			fmt.Println(iter)
		}
		var current step
		current, opened = Pop(opened)
		if statesEqual(current, finale) {
			return current
		}
		closed = append(closed, current)
		n = len(opened)

		newVars := genVariants(current.state)
		var newSteps []step
		for i := 0; i < len(newVars); i++ {
			newSteps = append(newSteps, newStep(newVars[i], heuristic(newVars[i]), current.g+1, &current))
		}
		newSteps = cleanAmongAllVariants(closed, opened, newSteps)
		opened = append(opened, newSteps...)
		n = len(opened)
	}

	fmt.Println("fail")
	return step{}
}

func main() {
	res := aStar(start)
	way(&res)
}
