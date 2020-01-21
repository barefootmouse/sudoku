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
	err := s.NewLevel(17)
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
}
