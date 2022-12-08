package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
)

/*
+---+
    |
    |
    |
   ===

Secret Word : ______
Incorrect Guesses :

Guess a Letter : a

Sorry Your Dead! The word is ZOMBIE
Yes the Secret Word is ZOMBIE

Please Enter Only One Letter
Please Enter a Letter
Please Enter a Letter you Haven't Guessed
*/

var pl = fmt.Println

var hmArr = [7]string{
	" +---+\n" +
		"     |\n" +
		"     |\n" +
		"     |\n" +
		"    ===\n",
	" +---+\n" +
		" 0   |\n" +
		"     |\n" +
		"     |\n" +
		"    ===\n",
	" +---+\n" +
		" 0   |\n" +
		" |   |\n" +
		"     |\n" +
		"    ===\n",
	" +---+\n" +
		" 0   |\n" +
		"/|   |\n" +
		"     |\n" +
		"    ===\n",
	" +---+\n" +
		" 0   |\n" +
		"/|\\  |\n" +
		"     |\n" +
		"    ===\n",
	" +---+\n" +
		" 0   |\n" +
		"/|\\  |\n" +
		"/    |\n" +
		"    ===\n",
	" +---+\n" +
		" 0   |\n" +
		"/|\\  |\n" +
		"/ \\  |\n" +
		"    ===\n",
}

var wordArr = [7]string{
	"JAZZ", "ZIGZAG", "ZILCH", "ZIPPER",
	"ZODIAC", "ZOMBIE", "FLUFF",
}

// Stores the random word to be guessed
var randWord string

// Stores all letters guessed
var guessedLetters string

// Stores correct guesses
var correctLetters []string

// Letters guessed that aren't in the randWord
var wrongGuesses []string

func getRandWord() string {
	// Returns seconds as int
	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)
	// Get random word from array
	randWord = wordArr[rand.Intn(7)]

	// Generate correctLetters array with correct size
	correctLetters = make([]string, len(randWord))
	return randWord
}

func showBoard() {
	// Prints current hangman state
	pl(hmArr[len(wrongGuesses)])

	fmt.Print("Secret Word : ")
	// Prints spaces for missing letter or the
	// letter
	// Cycle and if is a character print it
	// If not print an underscore
	for _, v := range correctLetters {
		if v == "" {
			fmt.Print("_")
		} else {
			fmt.Print(v)
		}
	}

	fmt.Print("\nIncorrect Guesses : ")
	if len(wrongGuesses) > 0 {
		for _, v := range wrongGuesses {
			fmt.Print(v + " ")
		}
	}
	pl()
}

func getUserLetter() string {
	// Setup buffered reader that gets text from the keyboard
	reader := bufio.NewReader(os.Stdin)
	// The guess the user makes
	var guess string

	for true {
		fmt.Print("\nGuess a Letter : ")
		guess, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		// Convert to upper so it is easier to search
		guess = strings.ToUpper(guess)

		// Remove spaces
		guess = strings.TrimSpace(guess)

		// Create regular expression that verifies
		// if a single character is a letter
		var IsLetter = regexp.MustCompile(`^[a-zA-Z]$`).MatchString

		if len(guess) != 1 {
			fmt.Println("Please Enter Only One Letter")
		} else if !IsLetter(guess) {
			fmt.Println("Please Enter a Letter")
		} else if strings.Contains(guessedLetters, guess) {
			fmt.Println("Please Enter a Letter you Haven't Guessed")
		} else {
			return guess
		}
	}
	return guess
}

// Returns all indices matching the substring
func getAllIndexes(theStr, subStr string) (indices []int) {
	if (len(subStr) == 0) || (len(theStr) == 0) {
		return indices
	}

	// Offset will find 1st matching index and then
	// move forward to find the rest
	offset := 0
	for {
		i := strings.Index(theStr[offset:], subStr)
		// If no matches or no more matches
		// since we cycled to the end
		// return indices
		if i == -1 {
			// fmt.Println(indices)
			return indices
		}
		offset += i
		// Add index
		indices = append(indices, offset)
		// Jump ahead the length of the substring
		offset += len(subStr)
	}
}

// Updates correctLetters string to display guessed
// letters an _s
func updateCorrectLetters(letter string) {
	indexMatches := getAllIndexes(randWord, letter)

	// Assign all matching indexes the
	// supplied letter
	for _, v := range indexMatches {
		correctLetters[v] = letter
	}
}

// Checks if array contains an empty value
func sliceHasEmptys(theSlice []string) bool {
	for _, v := range theSlice {
		if len(v) == 0 {
			return true
		}
	}
	return false
}

func main() {
	pl(getRandWord())

	for true {
		// Display opening board
		showBoard()

		// Get user guess
		guess := getUserLetter()

		if strings.Contains(randWord, guess) {
			updateCorrectLetters(guess)

			// Check if user guessed all
			// letters
			if sliceHasEmptys(correctLetters) {
				fmt.Println("More Letters to Guess")
			} else {
				fmt.Println("Yes the Secret Word is", randWord)
				break
			}

		} else {
			// They guessed a letter not in the secret
			// word
			// Append to string
			guessedLetters += guess
			// Append to slice
			wrongGuesses = append(wrongGuesses, guess)

			// Check if users ran out of guesses
			if len(wrongGuesses) >= 6 {
				fmt.Println("Sorry Your Dead! The word is", randWord)
				break
			}
		}
	}
}
