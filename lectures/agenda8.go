package lectures

import (
	"fmt"
	"log"
	"net/http"
)

// Control flow
// - Defer
// - Panic
// - Recover

func Agenda8() {
	panicFlow4()
}

func deferFlow() {
	defer fmt.Println("start")
	fmt.Println("middle") // defer executes the function at the end of the block
	// defers are ordered with last declared first executed
	// Run in LIFO (last-in, first-out) order
	fmt.Println("end")
}

func deferFlow2() {
	a := "start"
	defer fmt.Println(a) // take the argument at the time that the functions is called
	// not at the time that the functions is executed
	a = "end"
}

func panicFlow() {
	a, b := 1, 0
	ans := a / b
	fmt.Println(ans)
}

func panicFlow2() {
	fmt.Println("start")
	panic("something bad happened")
	fmt.Println("end")
}

func panicFlow3() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Go!"))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		// panic only happens after the defers
		panic(err.Error())
	}
}

func panicFlow4() {
	fmt.Println("start")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
		}
	}() // this empty parentheses, is what actually calls the anonymous function
	panic("something bad happened")
	fmt.Println("end")
}
