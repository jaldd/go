package main

import (
	"fmt"
	"math"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("%v %s", e.When, e.What)
}

func run() error {

	return &MyError{time.Now(), "not work"}
}

func main() {
	//if err := run(); err != nil {
	//	fmt.Println(err)
	//}

	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("%f", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		e := ErrNegativeSqrt(x)
		return 0, e
	} else {
		return math.Sqrt(x), nil
	}
}
