package main

import "fmt"

func sum(a, b int) int {
	return a+b 
}
func main() {
	var x = sum(2, 3)
	fmt.Printf("Hello sum: %d\n", x)
}