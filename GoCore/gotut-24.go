package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {
	// ----- AUTOMATED TESTING -----
	// Automated tests make sure your program still
	// works while you change the code
	// Create app2 directory with testemail.go
	// cd app2
	// Create testemail_test.go
	// go mod init app2
	// Run Tests : go test -v
}

/* testemail.go
package app2

import (
	"fmt"
	"regexp"
)

func IsEmail(s string) (string, error) {
	// Used a raw string here so I didn't have
	// to double backslashes
	r, _ := regexp.Compile(`[\w._%+-]{1,20}@[\w.-]{2,20}.[A-Za-z]{2,3}`)

	if r.MatchString(s) {
		return "Valid Email", nil
	} else {
		return "", fmt.Errorf("not a valid email")
	}
}
*/

/* testemail_test.go
package app2

import (
	"testing"
)

func TestIsEmail(t *testing.T) {
	_, err := IsEmail("hello")
	if err == nil {
		t.Error("hello is not an email")
	}

	_, err = IsEmail("derek@aol.com")
	if err != nil {
		t.Error("derek@aol.com is an email")
	}

	_, err = IsEmail("derek@aol")
	if err != nil {
		t.Error("derek@aol is not email")
	}
}
*/
