package main

import (
	"fmt"
	"regexp"
)

var pl = fmt.Println

func main() {
	// ----- REGULAR EXPRESSIONS -----
	// You can use regular expressions to test
	// if a string matches a pattern

	// Search for ape followed by not a space
	reStr := "The ape was at the apex"
	match, _ := regexp.MatchString("(ape[^ ]?)", reStr)
	pl(match)

	// You can compile them
	// Find multiple words ending with at
	reStr2 := "Cat rat mat fat pat"
	r, _ := regexp.Compile("([crmfp]at)")

	// Did you find any matches?
	pl("MatchString :", r.MatchString(reStr2))

	// Return first match
	pl("FindString :", r.FindString(reStr2))

	// Starting and ending index for 1st match
	pl("Index :", r.FindStringIndex(reStr2))

	// Return all matches
	pl("All String :", r.FindAllString(reStr2, -1))

	// Get 1st 2 matches
	pl("All String :", r.FindAllString(reStr2, 2))

	// Get indexes for all matches
	pl("All Submatch Index :", r.FindAllStringSubmatchIndex(reStr2, -1))

	// Replace all matches with Dog
	pl(r.ReplaceAllString(reStr2, "Dog"))
}
