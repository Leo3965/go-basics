package lectures

import (
	"bytes"
	"fmt"
	"io"
)

// Basics
// Composing interfaces
// Type conversion
// The empty interface
// Type switches
// Implementing with values vs pointers

// Best practices
// - Use many, small interfaces
// - don't export interfaces for types that will be consumed
// - Do export interfaces for types that will be used by package
//  - Design functions and methods to receive interfaces whenever possible

func Agenda11() {
	emptyInterface()
}

// interfaces do not describe data, interfaces describes behaviours
// Method definition

type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

type Incrementer interface {
	Increment() int
}

type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}

func basics() {
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello Go!"))

	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	for i = 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}
}

type Closer interface {
	Close() error
}

type WriterCloser interface {
	Writer
	Closer
}

type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}

	v := make([]byte, 8)
	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}

		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}
	return n, nil
}

func (bwc *BufferedWriterCloser) Close() error {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

func NewBufferedWriterCloser() *BufferedWriterCloser {
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}

func composingInterfaces() {
	var wc WriterCloser = NewBufferedWriterCloser()
	wc.Write([]byte("Hello Youtube listeners, this is a test"))
	wc.Close()

	bwc := wc.(*BufferedWriterCloser)
	fmt.Println(bwc)

	r, ok := wc.(io.Reader)
	if ok {
		fmt.Println(r)
	} else {
		fmt.Println("Conversion failed")
	}
}

func emptyInterface() {
	var myObj interface{} = NewBufferedWriterCloser()
	if wc, ok := myObj.(WriterCloser); ok {
		wc.Write([]byte("Hi"))
		wc.Close()
	}

	var i interface{} = 0
	switch i.(type) {
	case int:
		fmt.Println("i is an integer")
	case string:
		fmt.Println("i is a string")
	default:
		fmt.Println("I don't know what i is")

	}
}
