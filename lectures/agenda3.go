package lectures

import "fmt"

// Naming convention
// Typed constants
// Untyped constants
// Enumerated constants
// Enumeration expressions

// Questions what are bytes? And what is beat shifting - Restudy this topic

func Agenda3() {
	untyped()
}

func typed() {
	const (
		a = iota
		b
		c
	)

	const (
		a2 = iota
	)

	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)

	fmt.Printf("%v\n", a2)

	const myConst int = 42
	fmt.Printf("%v, %T\n", myConst, myConst)
}

func untyped() {
	const (
		_ = iota
		catSpecialist
		dogSpecialist
		snakeSpecialist
	)

	var specialist int
	fmt.Printf("%v\n", specialist == dogSpecialist)
}
