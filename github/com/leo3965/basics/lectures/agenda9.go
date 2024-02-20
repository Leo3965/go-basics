package lectures

import "fmt"

// Creating pointers
// Dereferencing pointers
// The new function
// Working with nil
// Types with internal pointers

func Agenda9() {
	pointersWithMaps()
}

func pointers() {
	//  (&) address operator - addressof
	//  (*) dereferencing operator
	var a int = 42
	var b *int = &a
	fmt.Println(a, b)
	a = 27
	fmt.Println(a, b)
}

func pointersWithArray() {
	// pointers arithmetics is not allowed by default in go
	// but you cna use the unsafe package
	a := [3]int{1, 2, 3}
	b := &a[0]
	c := &a[1] // - 4

	fmt.Printf("%v %p %p\n", a, b, c)
}

func pointersWithObjects() {
	type myStruct struct {
		foo int
	}

	var ms *myStruct
	fmt.Println(ms)
	ms = &myStruct{2}
	ms2 := new(myStruct)
	fmt.Println(ms)
	fmt.Println(ms2)

	var ms3 *myStruct
	ms3 = new(myStruct)
	ms.foo = 42          // how the compiler sees --> (*ms3).foo = 42
	fmt.Println(ms3.foo) // fmt.Println((*ms3).foo)
}

func pointersWithSlices() {
	//a := [3]int{1, 2, 3} // Arrays when copied makes a new object

	a := []int{1, 2, 3} // slices shares the same pointers underneath,
	// so when we pass a slice, we are actually pointing to the same array

	b := a
	fmt.Println(a, b)
	a[1] = 42
	fmt.Println(a, b)
}

func pointersWithMaps() {
	// Just as slices maps shares tha same underling pointer
	// This behaviour does not occur with primitive types like int. array
	a := map[string]string{"foo": "bar"}
	b := a
	fmt.Println(a, b)
	a["foo"] = "qux"
	fmt.Println(a, b)
}
