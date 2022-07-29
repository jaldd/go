package main

import (
	"fmt"
	"math"
)

type V1 struct {
	X, Y float64
}

func (v V1) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Abs1(v V1) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat float64

func (f MyFloat) Abc2() float64 {

	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *V1) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Scale1(v *V1, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {

	v := V1{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(Abs1(v))

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abc2())
	v.Scale(10)
	fmt.Println(v.Abs())

	Scale1(&v, 10)
	fmt.Println(Abs1(v))
	v1 := &v
	v1.Scale(5)
	fmt.Println(Abs1(*v1))
}
