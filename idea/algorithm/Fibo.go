package main

import "fmt"

func Fibo(n, a1, a2 int) int {
	if n == 0 {
		return a1
	}
	return Fibo(n-1, a2, a1+a2)
}

func main() {

	fmt.Println(Fibo(5, 1, 1))
}
