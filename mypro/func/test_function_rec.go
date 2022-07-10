package main

import "fmt"

func f2(a int) int {
	if a == 1 {

		return 1
	} else {
		return a * f2(a-1)
	}
}

func main() {

	fmt.Printf("f2(5): %v\n", f2(5))
}
