package lectures

import (
	"fmt"
	"sync"
)

// CHANNELS
// channel basics
// restricting data flow
// buffered channels
// closing channels
// for...range loops with channels
// select statements

func Agenda13() {
	loop2()
}

var wgc = sync.WaitGroup{}

func basicsChannels() {
	ch := make(chan int)
	wg.Add(2)

	go func() {
		i := <-ch // receiving data from channel
		fmt.Println(i)
		wg.Done()
	}()

	go func() {
		ch <- 42 // adding data to the channel
		wg.Done()
	}()
	wg.Wait()
}

func loop() {
	ch := make(chan int, 10)
	wg.Add(2)
	go func(ch <-chan int) {
		for i := range ch {
			fmt.Println(i)
		}
		wg.Done()
	}(ch)
	go func(ch chan<- int) {
		ch <- 42 // adding data to the channel
		ch <- 27
		close(ch) // closing the channel
		wg.Done()
	}(ch)
	wg.Wait()
}

func loop2() {
	ch := make(chan int, 10)
	wg.Add(2)
	go func(ch <-chan int) {
		for {
			if i, ok := <-ch; ok {
				fmt.Println(i)
			} else {
				break
			}
		}
		wg.Done()
	}(ch)
	go func(ch chan<- int) {
		ch <- 42 // adding data to the channel
		ch <- 27
		close(ch) // closing the channel
		wg.Done()
	}(ch)
	wg.Wait()
}
