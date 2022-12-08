package main

import (
	"fmt"
)

var pl = fmt.Println

// ----- FUNCTIONS -----
func sayHello() {
	pl("Hello")
}

// Returns sum of values
func getSum(x int, y int) int {
	return x + y
}

// Return multiple values
func getTwo(x int) (int, int) {
	return x + 1, x + 2
}

// Return potential error
func getQuotient(x float64, y float64) (ans float64, err error) {
	if y == 0 {
		// Define error message returned with dummy value
		// for ans
		return 0, fmt.Errorf("You can't divide by zero")
	} else {
		// If no error return nil
		return x / y, nil
	}
}

// Variadic function
func getSum2(nums ...int) int {
	sum := 0
	// nums gets converted into a slice which is
	// iterated by range (More on slices later)
	for _, num := range nums {
		sum += num
	}
	return sum
}

func getArraySum(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}

func changeVal(f3 int) int {
	f3 += 1
	return f3
}

func main() {
	// ----- FUNCTIONS -----
	// func funcName(parameters) returnType {BODY}
	// If you only need a function in the current package
	// start with lowercase letter
	// Letters and numbers in camelcase
	sayHello()
	pl(getSum(5, 4))
	f1, f2 := getTwo(5)
	fmt.Printf("%d %d\n", f1, f2)

	// Function that can return an error
	ans, err := getQuotient(5, 0)
	if err == nil {
		pl("5/4 =", ans)
	} else {
		pl(err)
		// End program
		// log.Fatal(err)
	}

	// Function receives unknown number of parameters
	// Variadic Function
	pl("Unknown Sum :", getSum2(1, 2, 3, 4))

	// Pass an array to a function by value
	vArr := []int{1, 2, 3, 4}
	pl("Array Sum :", getArraySum(vArr))

	// Go passes the value to functions so it isn't changed
	// even if the same variable name is used
	f3 := 5
	pl("f3 before func :", f3)
	changeVal(f3)
	pl("f3 after func :", f3)
}
