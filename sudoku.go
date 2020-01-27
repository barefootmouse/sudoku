// Package sudoku implements a solver and generator for 9x9 Sudoku's
package sudoku

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

// Board holds the Sudoku
type Board struct {
	Puzzle       string // Holds the puzzle of the Sudoku
	Solution     string // Holds the solution of the Sudoku
	Level        Level  // The difficulty level of the Sudoku
	Backtracking uint   // Amount of recursive call's made by Solve()
	Solved       bool   // Will be true if Sudoku has been solved
	Cells        []Cell // Each individual cell of the Sudoku board
}

// Cell is an individual cell of the Sudoku Board
type Cell struct {
	Digit  int  // The digit of the cell, for an empty cell the value 0 is used
	Solved bool // Used by the backtracking implementation
	Row    int  // The row on the Sudoku board
	Column int  // The column on the Sudoku board
}

const size int = 9 // Amount of rows/columns in a Sudoku

// Difficulty level of Sudoku puzzle
type Level int

const (
	Diabolic Level = 17
	Extreme  Level = 18
	Expert   Level = 20
	VeryHard Level = 24
	Hard     Level = 28
	Medium   Level = 30
	Easy     Level = 32
	VeryEasy Level = 36
	Unknown  Level = 0
)

// NewLevel generates a Sudoku, with level amount of digits filled.
// Returns an error when level is smaller than 17 or larger than 80.
// Make sure to properly seed math/rand before calling this method.
func (b *Board) NewLevel(l Level) error {

	// Create an empty board
	for row := 1; row <= size; row++ {
		for column := 1; column <= size; column++ {
			b.Cells = append(b.Cells, Cell{
				Digit:  0,
				Solved: false,
				Row:    row,
				Column: column,
			})
		}
	}

	// A Sudoku requires at least 17 digits and at most 80 digits
	if l < 17 || l > 80 {
		return errors.New("level should be between 17 and 80")
	}

	// Fill the board with random values
	counter := 0
	for counter < int(l) {
		// 81 cells in the board
		cell := rand.Intn(81)

		if !b.Cells[cell].Solved {
			// Digit should range from 1-9
			digit := rand.Intn(10)

			if b.isSafe(b.Cells[cell].Row, b.Cells[cell].Column, digit) {
				b.Cells[cell].Digit = digit
				b.Cells[cell].Solved = true
				counter++
			}
		}
	}

	var puzzle strings.Builder
	for _, v := range b.Cells {
		puzzle.WriteString(strconv.Itoa(v.Digit))
	}
	b.Puzzle = puzzle.String()

	b.determineLevel()

	return nil

}

// NewPuzzle generates a Sudoku from puzzle.
// Returns an error when puzzle doesn't contain 81 digits
func (b *Board) NewPuzzle(puzzle string) error {
	if len(puzzle) != 81 {
		return errors.New("puzzle should contain 81 digits")
	}

	counter := 0
	for row := 1; row <= size; row++ {
		for column := 1; column <= size; column++ {
			digit, err := strconv.ParseInt(puzzle[counter:counter+1], 0, 32)
			if err != nil {
				return errors.New("couldn't interpret the digit")
			}

			b.Cells = append(b.Cells, Cell{
				Digit:  int(digit),
				Solved: digit > 0,
				Row:    row,
				Column: column,
			})
			counter++
		}
	}

	b.determineLevel()

	return nil
}

// Print will output a 9x9 Sudoku to stdout. Empty values will be represented with a dot.
func (b *Board) Print() {
	fmt.Println("-------------------------------")

	for k, v := range b.Cells {
		// First digit of the row
		if v.Column == 1 {
			fmt.Print("|")
		}

		// Filter out the zero values and replace it with a dot
		// Otherwise just print the Digit
		if v.Digit == 0 {
			fmt.Print(" . ")
		} else {
			fmt.Printf(" %v ", v.Digit)
		}

		// After 3 digits
		if (k+1)%3 == 0 {
			fmt.Print("|")
		}
		// After 9 digits (one row)
		if (k+1)%9 == 0 {
			fmt.Println("")
		}
		// After 27 digits (3 rows)
		if (k+1)%27 == 0 {
			fmt.Println("-------------------------------")
		}
	}
	fmt.Println()
}

// determineLevel determines the difficulty of the Sudoku puzzle.
func (b *Board) determineLevel() {
	l := 0
	for _, value := range b.Cells {
		if value.Digit != 0 {
			l++
		}
	}

	switch Level(l) {
	case Diabolic:
		b.Level = Diabolic
	case Extreme:
		b.Level = Extreme
	case Expert:
		b.Level = Expert
	case VeryHard:
		b.Level = VeryHard
	case Hard:
		b.Level = Hard
	case Medium:
		b.Level = Medium
	case Easy:
		b.Level = Easy
	case VeryEasy:
		b.Level = VeryEasy
	default:
		b.Level = Unknown
	}
}

// isSafe returns true if number isn't found in the row, column or box.
func (b *Board) isSafe(row, column, number int) bool {
	return !b.inBox(row, column, number) && !b.inRow(row, number) && !b.inColumn(column, number)
}

// inRow returns true when number is in the row.
func (b *Board) inRow(row, number int) bool {
	for _, v := range b.Cells {
		if v.Row == row && v.Digit == number {
			return true
		}
	}
	return false
}

// row returns a []Cell for the given row.
func (b *Board) row(row int) []Cell {
	var c []Cell

	for _, v := range b.Cells {
		if v.Row == row {
			c = append(c, v)
		}
	}

	return c
}

// inColumn returns true when number is in the column.
func (b *Board) inColumn(column, number int) bool {
	for _, v := range b.Cells {
		if v.Column == column && v.Digit == number {
			return true
		}
	}
	return false
}

// column returns a []Cell for the given column.
func (b *Board) column(column int) []Cell {
	var c []Cell

	for _, v := range b.Cells {
		if v.Column == column {
			c = append(c, v)
		}
	}

	return c
}

// inBox returns true if number is in the box (3x3).
func (b *Board) inBox(row, column, number int) bool {
	row--
	column--

	rowMin := (row - (row % 3)) + 1
	columnMin := (column - (column % 3)) + 1

	rowMax := rowMin + 3
	columnMax := columnMin + 3

	for _, v := range b.Cells {
		if v.Row >= rowMin && v.Row < rowMax && v.Column >= columnMin && v.Column < columnMax && v.Digit == number {
			return true
		}
	}

	return false
}

// box returns a []Cell for the given box.
func (b *Board) box(row, column int) []Cell {
	var c []Cell

	row--
	column--

	rowMin := (row - (row % 3)) + 1
	columnMin := (column - (column % 3)) + 1

	rowMax := rowMin + 3
	columnMax := columnMin + 3

	for _, v := range b.Cells {
		if v.Row >= rowMin && v.Row < rowMax && v.Column >= columnMin && v.Column < columnMax {
			c = append(c, v)
		}
	}

	return c
}

// Solve solves the Sudoku, will return false if Sudoku is unsolvable.
func (b *Board) Solve() bool {
	for k := range b.Cells {
		number := &b.Cells[k]

		if !number.Solved {
			for d := 1; d <= 9; d++ {
				if b.isSafe(number.Row, number.Column, d) {
					number.Solved = true
					number.Digit = d

					b.Backtracking++
					if b.Solve() {
						return true
					}

					number.Solved = false
					number.Digit = 0
				}
			}
			return false
		}
	}

	b.Solved = true

	var s strings.Builder
	for _, v := range b.Cells {
		s.WriteString(strconv.Itoa(v.Digit))
	}
	b.Solution = s.String()

	return true
}

// SolveWithOptimizer solves the Sudoku using the specified Optimizer, will return false if Sudoku is unsolvable.
func (b *Board) SolveWithOptimizer(optimizer Optimizer) bool {
	return optimizer.Solve()
}
