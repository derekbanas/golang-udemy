package main

import (
	"fmt"
)

var pl = fmt.Println

// ----- INTERFACES -----
type Animal interface {
	AngrySound()
	HappySound()
}

// Define type with interface methods and its
// own method
type Cat string

func (c Cat) Attack() {
	pl("Cat Attacks its Prey")
}

// Return the cats name with a type conversion
func (c Cat) Name() string {
	return string(c)
}

func (c Cat) AngrySound() {
	pl("Cat says Hissssss")
}
func (c Cat) HappySound() {
	pl("Cat says Purrr")
}

func main() {
	// ----- INTERFACES -----
	// Interfaces allow you to create contracts
	// that say if anything inherits it that
	// they will implement defined methods

	// If we had animals and wanted to define that
	// they all perform certain actions, but in their
	// specific way we could use an interface

	// With Go you don't have to say a type uses
	// an interface. When your type implements
	// the required methods it is automatic
	var kitty Animal
	kitty = Cat("Kitty")
	kitty.AngrySound()

	// We can only call methods defined in the
	// interface for Cats because of the contract
	// unless you convert Cat back into a concrete
	// Cat type using a type assertion
	var kitty2 Cat = kitty.(Cat)
	kitty2.Attack()
	pl("Cats Name :", kitty2.Name())
}
