package lectures

import "fmt"

// Basic syntax
// Parameters
// Return values
// Anonymous functions
// Methods
// Functions as types

func Agenda10() {
	//sum(1, 2, 3, 4, 5)

	g := greeter{
		greeting: "Hello",
		name:     "Go",
	}

	g.greet()

}

// Variadic parameters
func sum(values ...int) {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}

	fmt.Println(result)
}

// Named return value
func sum2(values ...int) (result int) {
	fmt.Println(values)
	result = 0
	for _, v := range values {
		result += v
	}

	return // variable result is implicit returned here
}

func funcByValue() {
	g := "Hello"
	n := "Ted"
	sayGreeting(&g, &n)
	fmt.Println(n)
}

func sayGreeting(greeting, name *string) {
	fmt.Println(*greeting, *name) // dereferencing
	*name = "Leo"
}

func divide(a, b float32) (float32, error) {
	if b == 0.0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}

	return a / b, nil
}

func anonymousFunctions() {
	for i := 0; i < 5; i++ {
		func(i int) {
			fmt.Println(i)
		}(i)
	}
}

func functionAsVariable() {
	var f func() = func() {
		fmt.Println("Hello Go!")
	}
	f()

	var divide func(float64, float64) (float64, error)
	divide = func(a, b float64) (float64, error) {
		if b == 0.0 {
			return 0.0, fmt.Errorf("Cannot divide by zero")
		} else {
			return a / b, nil
		}
	}
	d, err := divide(5, 3)
	if err != nil {
		return
	}
	fmt.Println(d)

}

type greeter struct {
	greeting string
	name     string
}

// Methods
func (g *greeter) greet() {
	fmt.Println(g.greeting, g.name)
}
