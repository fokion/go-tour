package main

import (
	"fmt"
)

func swap(x int, y string) (string, int) {
	return y, x
}
func main() {
	a, b := swap(3, "test")
	fmt.Println(a, b)
	fmt.Printf("Test %T type and %v value\n", a, a)
	sum := 0
	count := 50
	for i := 0; i < count; i++ {
		sum += i
	}
	fmt.Println(sum)
	if sum > 100 {
		fmt.Println("Sum is more than 100")
	}
	if v := 2 * sum; v < 5000 {
		fmt.Println("v is smaller than 5000")
	}
}
