package main

import "fmt"

func main() {
	fmt.Println("Hello, world.")
}


// Add is our function that sums two integers
func Add(x, y int) (res int) {
	fmt.Println("Hello, world.")
	if x==4 {
		return x
	}
    return x + y
}

// Subtract subtracts two integers
func Subtract(x, y int) (res int) {
    return x - y
}