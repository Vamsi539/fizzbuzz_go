package main

import (
	"fmt"
)

func add(a int, b int) int {
	return a + b
}
func sub(a int, b int) int {
	return a - b
}
func mul(a int, b int) int {
	return a * b
}
func div(a float32, b float32) float32 {
	return a / b
}
func modl(a int, b int) int {
	return a % b
}
func fizzbuzz(i int) string {
	var temp = ""
	if i%3 == 0 {
		temp += "fizz"
	}
	if i%5 == 0 {
		temp += "buzz"
	}
	return temp
}

func main() {
	for i := 1; i <= 100; i++ {
		val := fizzbuzz(i)
		if val == "" {
			fmt.Println(i)
		} else {
			fmt.Println(val)
		}
	}
}
