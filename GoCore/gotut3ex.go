package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var pl = fmt.Println

func main() {
	// Setup buffered reader that gets text from the keyboard
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("What is your age : ")
	age, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	// Trims whitespace from age
	age = strings.TrimSpace(age)
	iAge, err := strconv.Atoi(age)
	if err != nil {
		log.Fatal(err)
	}

	if iAge < 5 {
		pl("Too young for school")
	} else if iAge == 5 {
		pl("Go to Kindergarten")
	} else if (iAge > 5) && (iAge <= 17) {
		grade := iAge - 5
		fmt.Printf("Go to grade %d\n", grade)
	} else {
		pl("Go to college")
	}
}
