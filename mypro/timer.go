package main

import (
	"fmt"
	"time"
)

func main() {
	// time1 := time.NewTimer(time.Second * 2)
	// t1 := time.Now()
	// fmt.Println("t1:", t1)
	// t2 := time1.C
	// fmt.Println("t2:", t2)

	// time2 := time.NewTimer(time.Second)
	// go func() {
	// 	<-time2.C
	// 	fmt.Println("3")
	// }()
	// time.Sleep(time.Second * 2)
	// stop := time2.Stop()
	// if stop {
	// 	fmt.Println("3 stop")
	// }

	fmt.Println("befort")
	time4 := time.NewTimer(time.Second * 5)
	time4.Reset(time.Second * 10)
	<-time4.C
	fmt.Println("after")
}
