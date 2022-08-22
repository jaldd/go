package main

import (
	"fmt"
	"sync"
	"time"
)

//func say(s string) {
//	for i := 0; i < 5; i++ {
//		time.Sleep(100 * time.Millisecond)
//		fmt.Println(s)
//	}
//}

//func sum1(s []int, c chan int) {
//
//	sum := 0
//	for _, v := range s {
//		sum += v
//	}
//	fmt.Println(sum)
//	c <- sum
//}
//func fibonacci1(n int, c chan int) {
//	x, y := 0, 1
//	for i := 0; i < n; i++ {
//		c <- x
//		x, y = y, x+y
//	}
//	close(c)
//}
//func fibonacci1(c, quit chan int) {
//
//	x, y := 0, 1
//	for {
//		select {
//		case c <- x:
//			x, y = y, x+y
//		case <-quit:
//			fmt.Println("quit")
//			return
//		}
//	}
//}

type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	c.v[key]++
	c.mux.Unlock()
}
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {

	//go say("world")
	//say("hello")
	//s := []int{7, 2, 8, -9, 4, 0}
	//c := make(chan int)
	//go sum1(s[:len(s)/2], c)
	//go sum1(s[len(s)/2:], c)
	//x, y := <-c, <-c
	//fmt.Println(x, y, x+y)

	//ch := make(chan int, 2)
	//ch <- 1
	//ch <- 2
	//fmt.Println(<-ch)
	//fmt.Println(<-ch)
	//x, y := 0, 1
	//x, y = y, x
	//fmt.Println(x, y)
	//c := make(chan int, 10)
	//go fibonacci1(cap(c), c)
	//for i := range c {
	//	fmt.Println(i)
	//}
	//c := make(chan int)
	//quit := make(chan int)
	//go func() {
	//	for i := 0; i < 10; i++ {
	//		fmt.Println(<-c)
	//	}
	//	quit <- 0
	//}()
	//fibonacci1(c, quit)
	//tick := time.Tick(100 * time.Millisecond)
	//boom := time.After(500 * time.Millisecond)
	//for {
	//	select {
	//	case <-tick:
	//		fmt.Println("tick")
	//	case <-boom:
	//		fmt.Println("boom")
	//		return
	//	default:
	//		fmt.Println(" .")
	//		time.Sleep(50 * time.Millisecond)
	//	}
	//}
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("some key")
	}
	time.Sleep(time.Second)
	fmt.Println(c.Value("some key"))
}
