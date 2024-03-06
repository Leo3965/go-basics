package lectures

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// GOUROUTINES

// Creating goroutines
// Synchronization
// - WaitGroups
// - Mutexes
// Parallelism

// Don't create goroutines in libraries
// - Let consumer control concurrency
// When creating a goroutine, know how it will end
// - Avoids subtle memory leaks
// Check fo race conditions at compile time

func Agenda12() {
	parallelism()
}

func creatingGoroutines() {
	hello := func() {
		fmt.Println("Hello")
	}

	go hello() // green thread
	time.Sleep(100 * time.Millisecond)
}

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func sayHello() {
	fmt.Printf("Hello #%v\n", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	counter++
	m.Unlock()
	wg.Done()
}

func waitGroups() {
	var msg = "Hello"
	wg.Add(1)
	go func(msg string) {
		fmt.Println(msg)
		wg.Done()
	}(msg)
	wg.Wait()
}

func synchronization() {
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go sayHello()
		m.Lock()
		go increment()
	}
	wg.Wait()
}

func parallelism() {
	fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(1)) // -1 returns the same value as before
	runtime.GOMAXPROCS(1)
	fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1))

}
