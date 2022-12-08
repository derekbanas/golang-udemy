package main

import (
	"fmt"
	_ stuff "example/project/mypackage"
)

var pl = fmt.Println

// ----- PACKAGES -----
// Packages allow you to keep related code together
// Go looks for package code in a directory

// If you are using VSC and have multiple
// modules you get this error
// gopls requires a module at the root of
// your workspace
// 1. Settings
// 2. In search type gopls
// 3. Paste "gopls": { "experimentalWorkspaceModule": true, }
// 4. Restart VSC

// cd /D D:\Tutorials\GoTutorial

// Create a go directory : mkdir app (Create in VSC)
// cd app (In terminal)
// Choose a module path and create a go.mod file
// go mod init example/project (In terminal)

// Go modules allow you to manage libraries
// They contain one project or library and a
// collection of Go packages
// go.mod : contains the name of the module and versions
// of other modules your module depends on

// Create a main.go file at the same level as go.mod

// You can have many packages and sub packages
// create a directory called mypackage in the project
// directory mkdir mypackage (Create in VSC)
// cd mypackage

// Create file mypackage.go in it

// Package names should be all lowercase

func main() {
	fmt.Println("Hello", stuff.Name)
	intArr := []int{2,3,5,7,11}
	strArr := stuff.IntArrToStrArr(intArr)
	fmt.Println(strArr)
	fmt.Println(reflect.Typeof(strArr))
}

/* mypackage.go
package stuff

import(
	"errors"
	"strconv"
	"time"
)

var Name string = "Derek"
func IntArrToStrArr(intArr []int) []string{
	var strArr []string
	for _, i := range intArr{
		strArr = append(strArr, strconv.Itoa(i))
	}
	return strArr
}

*/
