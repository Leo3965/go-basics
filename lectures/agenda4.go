package lectures

import "fmt"

// Arrays - creation, built-in functions, working with
// Slices - º//º

func Agenda4() {
	slicesWithMake()
}

func arrays() {
	//        [3]
	grades := [...]int{97, 85, 93}
	var students [3]string
	students[0] = "Leonardo"
	students[2] = "Lissa"
	students[1] = "Arnold"

	fmt.Printf("%v, %T\n", grades, grades)
	fmt.Printf("%v, %T\n", students, students)
	fmt.Printf("Student #1 %v\n", students[0])
	fmt.Printf("Number of students %v\n", len(students))

	var identityMatrix [3][3]int
	identityMatrix[0] = [3]int{1, 0, 0}
	identityMatrix[1] = [3]int{0, 1, 0}
	identityMatrix[2] = [3]int{0, 0, 1}
	fmt.Println(identityMatrix)

	a := [...]int{1, 2, 3}
	b := a  // Copy - deep copy
	c := &a // Pointer - shallow copy
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

func slices() {
	a1 := []int{1, 2, 3}
	b1 := a1 // Shallow copy
	b1[1] = 5
	fmt.Println(a1)
	fmt.Println(b1)
	fmt.Printf("Length: %v\n", len(a1))
	fmt.Printf("Capacity: %v\n", cap(a1))

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[:]   // slice of all elements
	c := a[3:]  // slices from 4th element to end
	d := a[:6]  // slice first until the 6th element
	e := a[3:6] // slice of the 4th to 6th elements
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	fmt.Printf("A type is %T", a)
}

func slicesWithMake() {
	a := make([]int, 3, 100)
	a = append(a, 1)
	a = append(a, []int{2, 3, 4}...)
	a = append(a, 5, 6, 7)

	b := a[1:] // removing first element
	fmt.Println("Removed first element: ", b)
	b = a[:len(a)-1] // removing last element
	fmt.Println("Removed last element: ", b)
	b = append(a[:2], a[3:]...) // Careful with this operation
	fmt.Println("Getting elements at the middle of the slices: ", b)

	fmt.Printf("Type: %T\n", a)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))
	fmt.Println("Slice: ", a)
}
