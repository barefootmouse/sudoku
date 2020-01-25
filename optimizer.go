package sudoku

import (
	"sort"
	"strconv"
	"strings"
)

// Optimizer is an interface that implements the Solve function.
type Optimizer interface {
	Solve() bool
}

// Scanner is an optimizer that starts at cell 1,1 and will continue down the line to the next cell 1.2, 1.3
//
// Scanner with the Reverse flag to true will start at 9.9 and continue to 9.8, 9.7
type Scanner struct {
	Board   *Board
	Reverse bool
}

// Solve solves the Sudoku, will return false if Sudoku is unsolvable.
func (o *Scanner) Solve() bool {
	for k := range o.Board.Cells {
		number := &o.Board.Cells[k]
		if o.Reverse {
			number = &o.Board.Cells[len(o.Board.Cells)-1-k]
		}

		if !number.Solved {
			for d := 1; d <= 9; d++ {
				if o.Board.isSafe(number.Row, number.Column, d) {
					number.Solved = true
					number.Digit = d

					o.Board.Backtracking++
					if o.Solve() {
						return true
					}

					number.Solved = false
					number.Digit = 0
				}
			}
			return false
		}
	}

	o.Board.Solved = true

	var s strings.Builder
	for _, v := range o.Board.Cells {
		s.WriteString(strconv.Itoa(v.Digit))
	}
	o.Board.Solution = s.String()

	return true
}

// HeatMap is an optimizer that will solve the cells with the most neighbors first
type HeatMap struct {
	Board   *Board
	heatmap map[int][]*Cell
	mapkeys []int
}

// Solve solves the Sudoku, will return false if Sudoku is unsolvable.
func (o *HeatMap) Solve() bool {
	if len(o.heatmap) == 0 {
		o.heatmap = make(map[int][]*Cell)

		for k, v := range o.Board.Cells {
			heat := 0

			r := o.Board.row(v.Row)
			c := o.Board.column(v.Column)
			b := o.Board.box(v.Row, v.Column)

			for _, rv := range r {
				if rv.Digit != 0 {
					heat++
				}
			}

			for _, vc := range c {
				if vc.Digit != 0 {
					heat++
				}
			}

			for _, vb := range b {
				if vb.Digit != 0 {
					heat++
				}
			}

			o.heatmap[heat] = append(o.heatmap[heat], &o.Board.Cells[k])
		}

		for key := range o.heatmap {
			o.mapkeys = append(o.mapkeys, key)
		}
		sort.Ints(o.mapkeys)

	}

	//fmt.Println(o.heatmap)
	for key := range o.mapkeys {
		key = o.mapkeys[len(o.mapkeys)-1-key]
		//fmt.Printf("%v: %v\n", key, o.heatmap[key])

		for _, value := range o.heatmap[key] {
			if !value.Solved {
				for d := 1; d <= 9; d++ {
					if o.Board.isSafe(value.Row, value.Column, d) {
						value.Solved = true
						value.Digit = d

						o.Board.Backtracking++
						if o.Solve() {
							return true
						}

						value.Solved = false
						value.Digit = 0
					}
				}
				return false
			}
		}

	}
	o.Board.Solved = true

	var s strings.Builder
	for _, v := range o.Board.Cells {
		s.WriteString(strconv.Itoa(v.Digit))
	}
	o.Board.Solution = s.String()

	return true

}
