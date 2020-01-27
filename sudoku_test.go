package sudoku

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestBoard_NewPuzzle(t *testing.T) {

	// Test with correct Sudoku.
	s := Board{}
	err := s.NewPuzzle("800000000003600000070090200050007000000045700000100030001000068008500010090000400")
	if err != nil {
		t.Errorf(err.Error())
	}

	// Test with incorrect Sudoku.
	s = Board{}
	err = s.NewPuzzle("this is not a sudoku")
	if err == nil {
		t.Errorf("Incorrect Sudoku should have returned a non-nil Error, but did not!")
	}

	// Test with incorrect Sudoku.
	s = Board{}
	err = s.NewPuzzle("BAD_00000003600000070090200050007000000045700000100030001000068008500010090000400")
	if err == nil {
		t.Errorf("Incorrect Sudoku should have returned a non-nil Error, but did not!")
	}
}

func TestBoard_Solve(t *testing.T) {
	s := Board{}
	err := s.NewPuzzle("800000000003600000070090200050007000000045700000100030001000068008500010090000400")
	if err != nil {
		t.Errorf(err.Error())
	}

	solved := s.Solve()
	if !solved {
		t.Error("Failed to solve correct Sudoku")
	}
}

func TestBoard_NewLevel(t *testing.T) {
	// Should provide us with a correct Sudoku
	rand.Seed(42)

	s := Board{}
	err := s.NewLevel(Diabolic)
	if err != nil {
		t.Error(err.Error())
	}

	solved := s.Solve()
	if !solved {
		t.Error("Failed to solve correct Sudoku")
	}

	// Should return an error
	s = Board{}
	err = s.NewLevel(16)
	if err == nil {
		t.Error("Using 16 as level should fail")
	}

	// Should return an error
	s = Board{}
	err = s.NewLevel(82)
	if err == nil {
		t.Error("Using 82 as level should fail")
	}

	s = Board{}
	err = s.NewLevel(Extreme)
	if err != nil {
		t.Error(err.Error())
	}

	s = Board{}
	err = s.NewLevel(Expert)
	if err != nil {
		t.Error(err.Error())
	}

	s = Board{}
	err = s.NewLevel(VeryHard)
	if err != nil {
		t.Error(err.Error())
	}

	s = Board{}
	err = s.NewLevel(Hard)
	if err != nil {
		t.Error(err.Error())
	}

	s = Board{}
	err = s.NewLevel(Medium)
	if err != nil {
		t.Error(err.Error())
	}

	s = Board{}
	err = s.NewLevel(Easy)
	if err != nil {
		t.Error(err.Error())
	}

	s = Board{}
	err = s.NewLevel(VeryEasy)
	if err != nil {
		t.Error(err.Error())
	}

	s = Board{}
	err = s.NewLevel(40)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestBoard_Print(t *testing.T) {
	s := Board{}
	err := s.NewPuzzle("800000000003600000070090200050007000000045700000100030001000068008500010090000400")
	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println("Printing puzzle of Sudoku to stdout")
	s.Print()

	s.Solve()

	fmt.Println("Printing solution of Sudoku to stdout")
	s.Print()

	fmt.Println("Solution:", s.Solution)
	fmt.Println("Backtracking:", s.Backtracking)
}

func TestBoard_SolveWithOptimizer(t *testing.T) {
	s := Board{}
	err := s.NewPuzzle("000070065000050000001000000800000000007000050002000000000100002070800049000200070")
	if err != nil {
		t.Errorf(err.Error())
	}

	optimizer := Scanner{
		Board:   &s,
		Reverse: true,
	}

	fmt.Println("Printing puzzle of Sudoku to stdout")
	s.Print()

	s.SolveWithOptimizer(&optimizer)

	fmt.Println("Printing solution of Sudoku to stdout")
	s.Print()

	fmt.Println("Solution:", s.Solution)
	fmt.Println("Backtracking:", s.Backtracking)

	s = Board{}
	err = s.NewPuzzle("000070065000050000001000000800000000007000050002000000000100002070800049000200070")
	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println("Printing puzzle of Sudoku to stdout")
	s.Print()

	optimizerHM := HeatMap{Board: &s}
	s.SolveWithOptimizer(&optimizerHM)

	fmt.Println("Printing solution of Sudoku to stdout")
	s.Print()

	fmt.Println("Solution:", s.Solution)
	fmt.Println("Backtracking:", s.Backtracking)

}
