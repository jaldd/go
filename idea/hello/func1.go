package main

import "fmt"

func compute(fn func(float64, float64) float64) float64 {

	return fn(3, 4)
}
func adder() func(int) int {

	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func fibonacci() func() int {

	var x = 0
	var y = 1
	return func() int {
		z := x + y
		x = y
		y = z
		return z
	}
}

func main() {
	//hypot := func(x, y float64) float64 {
	//	return math.Sqrt(x*x + y*y)
	//}
	//fmt.Println(hypot(5, 12))
	//fmt.Println(compute(hypot))
	//fmt.Println(compute(math.Pow))
	//
	//pos, neg := adder(), adder()
	//for i := 0; i < 10; i++ {
	//	fmt.Println(pos(i), neg(-2*i))
	//}
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
