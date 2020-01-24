package main

import "fmt"

func Hello() string {
	return "Hello, world"
}

func Add(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}

	var result int
	for _, n := range nums {
		result += n
	}
	return result
}

func main() {
	fmt.Println(Hello())
}
