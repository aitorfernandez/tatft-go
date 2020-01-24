package main

import "fmt"

func Add(nums ...int) int {
	var result int
	for _, n := range nums {
		result += n
	}
	return result
}

func Sub(i int, nums ...int) int {
	for _, n := range nums {
		i -= n
	}
	return i
}

func main() {
	fmt.Println("tatft")
}
