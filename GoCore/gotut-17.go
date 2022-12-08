package main

import (
	"fmt"
)

var pl = fmt.Println

// ----- STRUCTS -----
type customer struct {
	name    string
	address string
	bal     float64
}

// This struct has a function associated
type rectangle struct {
	length, height float64
}

func (r rectangle) Area() float64 {
	return r.length * r.height
}

// Customer passed as values
func getCustInfo(c customer) {
	fmt.Printf("%s owes us %.2f\n", c.name, c.bal)
}

func newCustAdd(c *customer, address string) {
	c.address = address
}

// Struct composition : Putting a struct in another
type contact struct {
	fName string
	lName string
	phone string
}

type business struct {
	name    string
	address string
	contact
}

func (b business) info() {
	fmt.Printf("Contact at %s is %s %s\n", b.name, b.contact.fName, b.contact.lName)
}

func main() {
	// ----- STRUCTS -----
	// Structs allow you to store values with many
	// data types

	// Add values
	var tS customer
	tS.name = "Tom Smith"
	tS.address = "5 Main St"
	tS.bal = 234.56

	// Pass to function as values
	getCustInfo(tS)
	// or as reference
	newCustAdd(&tS, "123 South st")
	pl("Address :", tS.address)

	// Create a struct literal
	sS := customer{"Sally Smith", "123 Main", 0.0}
	pl("Name :", sS.name)

	// Structs with functions
	rect1 := rectangle{10.0, 15.0}
	pl("Rect Area :", rect1.Area())

	// Go doesn't support inheritance, but it does
	// support composition by embedding a struct
	// in another
	con1 := contact{
		"James",
		"Wang",
		"555-1212",
	}

	bus1 := business{
		"ABC Plumbing",
		"234 North St",
		con1,
	}

	bus1.info()
}
