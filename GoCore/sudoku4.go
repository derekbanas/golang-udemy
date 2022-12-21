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

	// NEW Check if column is valid
	for index := range puzz {
		if puzz[index][column] == guess && column != index {
			return false
		}
	}

	// NEW Check if the box is valid
	// We need to convert all rows and columns
	// into their 3x3 version so we can cycle
	// through them

	// We need to do math to generate the row and column
	// numbers to check in our slice

	// Let's use Row 0 Column 3 as an example
	// Row - (Row % 3) + Val for Cycling (0-2)
	// 0 - (0 % 3) + 0 = 0 - 0 + 0 = 0 (1st row in box)
	// 0 - (0 % 3) + 1 = 0 - 0 + 1 = 1 (2nd row in box)
	// 0 - (0 % 3) + 2 = 0 - 0 + 2 = 2 (3rd row in box)

	// Col - (Col % 3) + Val for Cycling (0-2)
	// 3 - (3 % 3) + 0 = 3 - 0 + 0 = 3 (1st col in box)
	// 3 - (3 % 3) + 1 = 3 - 0 + 1 = 4 (2nd col in box)
	// 3 - (3 % 3) + 2 = 3 - 0 + 2 = 5 (3rd col in box)

	for k := 0; k < 3; k++ {
		for l := 0; l < 3; l++ {
			// We need to check if any element in the box
			// is equal to what we just added but skip
			// checking the position we just placed our
			// value in
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
6 4 8 1 2 #9 (Change 9 to 0 and try another 2)
9 2 (1st row done)
7 9 3 (2nd row done)
#9
-------------------------
| 6 4 8 | 1 3 5 | 9 7 2 |
| 2 5 7 | 9 4 6 | 8 3 1 |
| 0 1 3 | 7 0 8 | 0 4 9 |
-------------------------
*/

/*
9 8 4 1 3 5 6 7 2
2 5 7 9 4 6 8 3 1
6 1 3 7 2 8 5 4 9
1 9 6 3 5 7 2 8 4
7 4 5 8 1 2 3 9 6
8 3 2 6 9 4 1 5 7
3 7 8 2 6 9 4 1 5
5 6 1 3 7 4 9 2 8
4 2 9 5 8 1 7 6 3
*/

// Recursively solves puzzle by placing values
// if it reaches a point where there is no
// valid answer it changes the previous
// supposedly correct answer back and tries
// another value for what we previously though was
// correct

// By using this backtracking when we get to the 9 9
// position on our board we know we solved the puzzle
func solvePuzzle(puzz [][]int) bool {
	// Check for empty space and if there
	// is none return true meaning the puzzle
	// is solved, otherwise solve the puzzle
	row, column := getEmptySpace(puzz)
	if row == -1 {
		return true
	} else {
		row, column = getEmptySpace(puzz)
	}

	// Check if we can add a 1 through 9 into the
	// empty space
	for i := 1; i <= 9; i++ {
		// Is this a valid number?
		if isNumValid(puzz, i, row, column) {
			// If so put it in the empty space
			puzz[row][column] = i

			// Continue until solvePuzzle with new puzz
			// until it returns true meaning we no longer
			// have empty spaces and we are finished
			if solvePuzzle(puzz) {
				return true
			}
			// If solvePuzzle doesn't return true we
			// reset to an empty space and backtrack
			// and try the next value in the for loop
			puzz[row][column] = 0
		}
	}
	return false
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
	// fmt.Println(isNumValid(puzz, 1, 0, 0))
	// Is 7 valid for row 1
	// fmt.Println(isNumValid(puzz, 7, 0, 0))

	// NEW
	// Testing if row and column is valid
	fmt.Println(isNumValid(puzz, 1, 0, 0))
	fmt.Println(isNumValid(puzz, 6, 0, 0))
	fmt.Println(isNumValid(puzz, 9, 0, 0))
	fmt.Println(isNumValid(puzz, 2, 0, 1))

	fmt.Println(isNumValid(puzz, 7, 4, 0))
	fmt.Println(isNumValid(puzz, 8, 4, 1))

	displayBoard(puzz)
	fmt.Println()
	solvePuzzle(puzz)
	displayBoard(puzz)
}
