package lectures

import "fmt"

// Boolean type
// Numeric types - integers, floating point, complex numbers
// Text types

func Agenda2() {
	text()
}

func boolean() {
	var n bool = true
	fmt.Printf("%v, %T\n", n, n)
}

func numeric() {
	var i uint8
	i = 255
	fmt.Printf("%v %T\n", i, i)

	a := 10
	b := 3

	fmt.Println("Operations:", a+b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println("End Operations", a%b)

	// Bit shifting
	a = 8
	fmt.Println(a << 3) // 2^3 * 2^3 = 2^6 = 64
	fmt.Println(a >> 3) // 2^3 / 2^3 = 2^0 = 1

	f := 3.14
	f = 2.1e14 // 2,1 x 10^14
	fmt.Printf("%v %T\n", f, f)

	// Complex numbers
	//var c complex64 = complex(1, 2)
	var c complex64 = 1 + 2i
	fmt.Printf("%v %T\n", c, c)
	fmt.Println(real(c))
	fmt.Println(imag(c))
}

func text() {
	s := "this is a string" // Strings are just alias for bytes
	fmt.Printf("%v, %T\n", s[2], s[2])
	fmt.Printf("%v, %T\n", string(s[2]), s[2])
	b := []byte(s) // slices of bytes
	fmt.Printf("%v, %T\n", b, b)
}
