package lectures

import "fmt"

// For statements
//	- simple loops
//	- exiting early
//	- looping through collections

func Agenda7() {
	slicesLoop()
}

func slicesLoop() {
	s := []int{1, 2, 3}
	for k, v := range s {
		fmt.Printf("Key: %d, Value: %v\n", k, v)
	}

	states := map[string]string{
		"SP": "São Paulo",
		"RJ": "Rio de Janeiro",
	}

	for k, v := range states {
		fmt.Printf("Key: %s, Value: %v\n", k, v)
	}

	text := "hello Go!"
	for k, v := range text {
		fmt.Printf("Key: %d, Value: %v\n", k, string(v))
	}
}

func loops() {
	for k := 0; k < 5; k++ {
		fmt.Println(i)
	}

	for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}

	o := 0
	for ; o < 5; o++ {
		fmt.Println(o)
	}

	// do while
	o = 0
	for o < 5 {
		fmt.Println(o)
		o++
	}

	// infinite lop
	l := 0
	for {
		l++
		if l > 100_000 {
			break
		}

		if l%1_000 == 0 {
			fmt.Println(l)
			continue
		}
	}
}

func nestedLoop() {
Loop: // Label
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(i * j)
			if i*j >= 3 {
				break Loop
			}
		}
	}
}
