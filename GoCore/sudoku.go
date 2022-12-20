package main

import (
	"fmt"
	"strconv"
)

/*
SUDOKU RULES
1. 9 by 9 square
2. Each row and column must contain numbers 1-9
3. Each 3x3 square must contain numbers 1-9
4. No repeats allowed in rows, columns or squares
*/

// Outputs the current board to screen
func displayBoard(puzz [][]int) {
	puzz_cols := len(puzz[0])
	// Cycles through rows in the puzz
	for i := 0; i < len(puzz); i++ {
		// Cycle through columns in puzz
		for j := 0; j < puzz_cols; j++ {
			fmt.Print(strconv.Itoa(puzz[i][j]) + " ")
		}
		fmt.Print("\n")
	}
}

func getEmptySpace(puzz [][]int) (int, int) {
	puzz_cols := len(puzz[0])
	// Cycles through rows in the puzz
	for i := 0; i < len(puzz); i++ {
		// Cycle through columns in puzz
		for j := 0; j < puzz_cols; j++ {
			if puzz[i][j] == 0 {
				return i, j
			}
		}
	}
	return -1, -1
}

// Receives slice, number to check, row &
// column and decides if that number fits
// the rules of the game
func isNumValid(puzz [][]int, guess int, row int, column int) bool {

	// Cycle through all values in row to see
	// if the row is valid (We only want index
	// which is the 1st value)
	for index := range puzz {
		// Use index for cycling
		// Check if any of the values are equal
		// to our guess and also that we didn't
		// already place it there in this
		// function
		if puzz[row][index] == guess && column != index {
			return false
		}
	}
	return true
}

func main() {

	// Holds starting puzzle
	puzz := [][]int{
		{0, 0, 0, 0, 3, 5, 0, 7, 0},
		{2, 5, 0, 0, 4, 6, 8, 0, 1},
		{0, 1, 3, 7, 0, 8, 0, 4, 9},
		{1, 9, 0, 0, 0, 7, 0, 0, 4},
		{0, 0, 5, 0, 0, 2, 0, 9, 6},
		{8, 0, 2, 0, 9, 4, 0, 0, 7},
		{3, 7, 0, 0, 0, 9, 0, 0, 0},
		{0, 6, 1, 0, 7, 0, 0, 0, 0},
		{4, 0, 0, 5, 8, 1, 0, 0, 0},
	}

	// ----- TESTING FUNCTIONS -----
	// Displays current board
	displayBoard(puzz)
	// Returns first open space
	row, _ := getEmptySpace(puzz)
	if row != -1 {
		fmt.Println(getEmptySpace(puzz))
	} else {
		fmt.Println("Puzzle is Solved")
	}
	// Testing if a guess is valid
	// Is 1 valid for row 1
	fmt.Println(isNumValid(puzz, 1, 0, 0))
	// Is 7 valid for row 1
	fmt.Println(isNumValid(puzz, 7, 0, 0))

}
