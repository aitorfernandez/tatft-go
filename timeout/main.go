package main

import "fmt"

func Add(nums ...int) int {
	var result int
	for _, n := range nums {
		result += n
	}
	return result
}

func main() {
	fmt.Println("tatft")
}
