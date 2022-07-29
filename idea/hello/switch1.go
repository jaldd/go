package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("os")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Println(os)
	}

	fmt.Println("saturday is :")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today:
		fmt.Println("today.")
	case today + 1:
		fmt.Println("tomorrow.")
	case today + 2:
		fmt.Println("hou day.")
	default:
		fmt.Println("far")
	}
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("morning")
	case t.Hour() < 17:
		fmt.Println("afternoon")
	default:
		fmt.Println("good night")
	}
}
