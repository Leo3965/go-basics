package lectures

import (
	"fmt"
	"strconv"
)

// Agenda
// Variable declaration
// Redeclaration and shadowing
// Visibility
// Naming conventions
// Type conversions

var i = 24

var (
	actorName    = "Elisabet"
	companion    = "Sara"
	doctorNumber = 3
	season       = 11
)

const (
	counter int = 0
)

// lowercase variables or methods are private and uppercase are public
// name of variables should match their uses, a count should have a single letter, 'cause it will be used fast and only
// at that block
// however if a variable is going to be used a lot it should have a meaningful name

func Agenda1() {
	var i int // Is the way that you speak
	// Shadowing: the package level variable are still available but is hidden by this scoped variable
	i = 42
	fmt.Println("First way to declare a variable", i)

	var j int = 27
	fmt.Println("Second way to declare a variable", j)

	k := 10
	fmt.Println("Third way to declare a variable", k)

	var o float32
	o = float32(i)
	fmt.Printf("Converting a int to float %v, %T\n", o, o)

	var s string
	s = strconv.Itoa(i)
	fmt.Printf("Converting a int to string %v, %T\n", s, s)

}
