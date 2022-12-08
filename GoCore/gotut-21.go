package main

import (
	"fmt"
)

var pl = fmt.Println

// ----- CLOSURES -----
// Pass a function to a function
func useFunc(f func(int, int) int, x, y int) {
	pl("Answer :", (f(x, y)))
}

func sumVals(x, y int) int {
	return x + y
}

func main() {
	// ----- CLOSURES -----
	// Closures are functions that don't have to be
	// associated with an identifier (Anonymous)

	// Create a closure that sums values
	intSum := func(x, y int) int { return x + y }
	pl("5 + 4 =", intSum(5, 4))

	// Closures can change values outside the function
	samp1 := 1
	changeVar := func() { samp1 += 1 }
	changeVar()
	pl("samp1 =", samp1)

	// Pass a function to a function
	useFunc(sumVals, 5, 8)
}
