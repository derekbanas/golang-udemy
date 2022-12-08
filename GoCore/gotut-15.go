package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {
	// ----- MAPS -----
	// Maps are collections of key/value pairs
	// Keys can be any data type that can be compared
	// using == (They can be a different type than
	// the value)
	// var myMap map [keyType]valueType

	// Declare a map variable
	var heroes map[string]string
	// Create the map
	heroes = make(map[string]string)

	// You can do it in one step
	villians := make(map[string]string)

	// Add keys and values
	heroes["Batman"] = "Bruce Wayne"
	heroes["Superman"] = "Clark Kent"
	heroes["The Flash"] = "Barry Allen"
	villians["Lex Luther"] = "Lex Luther"

	// Define with map literal
	superPets := map[int]string{1: "Krypto",
		2: "Bat Hound"}

	// Get value with key (Use %v with Printf)
	fmt.Printf("Batman is %v\n", heroes["Batman"])

	// If you access a key that doesn't exist
	// you get nil
	pl("Chip :", superPets[3])

	// You can check if there is a value or nil
	_, ok := superPets[3]
	pl("Is there a 3rd pet :", ok)

	// Cycle through map
	for k, v := range heroes {
		fmt.Printf("%s is %s\n", k, v)
	}

	// Delete a key value
	delete(heroes, "The Flash")
}
