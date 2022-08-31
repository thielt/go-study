package main

import (
	"fmt"
)

func main() {
	var wc WriterCloser = &myWriterCloser{}
	fmt.Println(wc)
}

type Writer interface {
	Write([]byte) (int, error) //defining behaviors it holds, holds the write method
}

type Closer interface {
	Close() error
}

type WriterCloser interface { //embeds the writer and closer together, if you just want one then use the other ones
	Writer
	Closer
}

type myWriterCloser struct{} //implementing myWriterCloser interface, which will be initialized in main as having behaviors of WriterCloser

func (mwc *myWriterCloser) Write(data []byte) (int, error) { //method
	return 0, nil
}

func (mwc *myWriterCloser) Close() error { //method
	return nil
}

// use many small interfaces
// single method interfaces are powerful
// don't export interfaces for types that will be consumed, except that will be used by package
