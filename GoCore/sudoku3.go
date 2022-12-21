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

	for index := range puzz {
		if puzz[index][column] == guess && column != index {
			return false
		}
	}

	// Is number valid for box?
	// Row 0 & Column 3
	// Row - (Row % 3) + Value for cycling (0-2)
	// 0 - (0 % 3) + 0 = 0 - 0 + 0 = 0 (1st row in box)
	// 0 - (0 % 3) + 1 = 0 - 0 + 1 = 1 (1st row in box)
	// 0 - (0 % 3) + 2 = 0 - 0 + 2 = 2 (1st row in box)

	// Col - (Col % 3) + Value for cycling (0-2)
	// 3 - (3 % 3) + 0 = 3 - 0 + 0 = 3 (1st col in box)
	// 3 - (3 % 3) + 1 = 3 - 0 + 1 = 4 (1st col in box)
	// 3 - (3 % 3) + 2 = 3 - 0 + 2 = 5 (1st col in box)

	for k := 0; k < 3; k++ {
		for l := 0; l < 3; l++ {
			if puzz[row-row%3+k][column-column%3+l] == guess && (row-row%3+k != row || column-column%3+l != column) {
				return false
			}
		}
	}
	return true
}

/*
-------------------------
| 0 0 0 | 0 3 5 | 0 7 0 |
| 2 5 0 | 0 4 6 | 8 0 1 |
| 0 1 3 | 7 0 8 | 0 4 9 |
-------------------------
| 1 9 0 | 0 0 7 | 0 0 4 |
| 0 0 5 | 0 0 2 | 0 9 6 |
| 8 0 2 | 0 9 4 | 0 0 7 |
-------------------------
| 3 7 0 | 0 0 9 | 0 0 0 |
| 0 6 1 | 0 7 0 | 0 0 0 |
| 4 0 0 | 5 8 1 | 0 0 0 |
-------------------------

// Recursion Problem (Solution)
// 1. Cycle across the rows column by column (1-9)
// 2. Check if valid number
// a. If true update array
// b. If false change back to 0
// c. If false find a new value for previous we
// thought was correct
// 3. Check next column
*/

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
	fmt.Println(isNumValid(puzz, 1, 0, 6))

}
