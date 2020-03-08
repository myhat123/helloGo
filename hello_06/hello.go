package main

import "fmt"
import "hello_06/tools"

func main() {
	var x = tools.Sum(2, 3)
	fmt.Printf("Hello sum: %d\n", x)
}