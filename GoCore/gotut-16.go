package main

import (
	"fmt"
)

var pl = fmt.Println

// ----- FUNCTION THAT EXCEPTS GENERICS -----
// This generic type parameter is capital, between
// square brackets and has a rule for what data
// it will except called a constraint
// any : anything
// comparable : Anything that supports ==
// More Constraints : pkg.go.dev/golang.org/x/exp/constraints

// You can also define what is excepted like this
// Define that my generic must be an int or float64
type MyConstraint interface {
	int | float64
}

func getSumGen[T MyConstraint](x T, y T) T {
	return x + y
}

func main() {
	// ----- GENERICS -----
	// We can specify the data type to be used at a
	// later time with generics
	// It is mainly used when we want to create
	// functions that can work with
	// multiple data types
	pl("5 + 4 =", getSumGen(5, 4))
	pl("5.6 + 4.7 =", getSumGen(5.6, 4.7))

	// This causes an error
	// pl("5.6 + 4.7 =", getSumGen("5.6", "4.7"))
}
