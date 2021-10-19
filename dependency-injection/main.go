package main

import (
	"fmt"
	"io"
	"os"
)

// func Greet(writer *bytes.Buffer, name string) {

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hey, %s", name)
}

func main() {
	Greet(os.Stdout, "Yo!")
}
