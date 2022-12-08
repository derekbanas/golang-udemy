package main

import "fmt"

var pl = fmt.Println

func main() {
	// ----- FOR LOOPS -----
	// for initialization; condition; postStatement {BODY}
	// Print numbers 1 through 5
	for x := 1; x <= 5; x++ {
		pl(x)
	}
	// Do the opposite
	for x := 5; x >= 1; x-- {
		pl(x)
	}

	// x is out of the scope of the for loop so it doesn't exist
	// pl("x :", x)

	// For is used to create while loops as well
	fX := 0
	for fX < 5 {
		pl(fX)
		fX++
	}

	// Cycle through an array with range
	// More on arrays later
	// We don't need the index so we ignore it
	// with the blank identifier _
	aNums := []int{1, 2, 3}
	for _, num := range aNums {
		pl(num)
	}

	// We can allow a condition in the for loop
	// to decide when to exit
	xVal := 1
	for true {
		if xVal == 5 {
			// Break exits the loop
			break
		}
		pl(xVal)
		xVal++
	}
}
