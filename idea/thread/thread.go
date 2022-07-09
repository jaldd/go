package main

import (
	"fmt"
	"time"
)

func Hu(ch chan int) {
	time.Sleep(2 * time.Second)
	fmt.Println("run 2 seconds")

	ch <- 1000
}

func Receive(ch chan int) {
	time.Sleep(2 * time.Second)
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				fmt.Println("close receive", v)
				return
			}
			fmt.Println("receive:", v)
		}
	}
}

func Send(ch chan int) {

	for i := 0; i < 13; i++ {
		ch <- i
		fmt.Println("send:", i)
	}
	close(ch)
}

func main() {

	ch := make(chan int)
	//go Hu(ch)
	//fmt.Println("start hu")
	////for {
	////	time.Sleep(1 * time.Second)
	////}
	//v := <-ch
	//fmt.Println("receive:", v)

	go Receive(ch)
	go Send(ch)
	for {
		time.Sleep(1 * time.Second)
	}
}
