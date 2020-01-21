package sudoku_test

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/barefootmouse/sudoku"
)

func ExampleBoard_NewLevel() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	s := sudoku.Board{}

	if err := s.NewLevel(17); err != nil {
		log.Fatalln(err)
	}

	s.Print()

	if s.Solve() == true {
		s.Print()
	} else {
		fmt.Println("Sudoku is unsolvable.")
	}
}

func ExampleBoard_NewPuzzle() {
	s := sudoku.Board{}

	err := s.NewPuzzle("800000000003600000070090200050007000000045700000100030001000068008500010090000400")
	if err != nil {
		log.Fatalln(err)
	}

	s.Print()

	if s.Solve() == true {
		s.Print()
	} else {
		fmt.Println("Sudoku is unsolvable.")
	}
}
