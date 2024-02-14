package lectures

import (
	"fmt"
	"math"
)

// If
//	- operations
//  - if else
// Switch statements
//	- simple cases
//	- cases with multiple tests
//	- falling through
//	- type switches

func Agenda6() {
	switchFunc()
}

func ifFunc() {
	if true {
		fmt.Println("test is true")
	}

	states := map[string]string{
		"SP": "São Paulo",
		"RJ": "Rio de Janeiro",
	}

	// if with initializer
	if pop, ok := states["RJ"]; ok {
		fmt.Println(pop)
	}

	// if with float numbers
	myNum := 0.123456789
	if math.Abs(myNum/math.Pow(math.Sqrt(myNum), 2)-1) < 0.001 {
		fmt.Println("These are the same")
	} else {
		fmt.Println("There are different")
	}
}

func switchFunc() {
	switch 5 {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3, 4, 5:
		fmt.Println("three, four or five")
	default:
		fmt.Println("not one, two three, four nor five")
	}

	switch i := 2 + 3; i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 5:
		fmt.Println("five")
	default:
		fmt.Println("not one, two three nor five")
	}

	j := 4
	switch {
	case j < 5:
		fmt.Println("less than five")
		fallthrough
	case j < 10:
		fmt.Println("less than ten")
	case j >= 10:
		fmt.Println("greater or equal to ten")
	default:
		fmt.Println("Unhandled value")
	}

	var inter interface{} = [2]int{}
	switch inter.(type) {
	case int:
		fmt.Println("is int")
	case string:
		fmt.Println("is it a string")
	case float64:
		fmt.Println("It is an float")
	case [2]int:
		fmt.Println("is it an array of two elements")
	default:
		fmt.Println("Do not know what king that is")
	}
}
