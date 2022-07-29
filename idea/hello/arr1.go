package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"golang.org/x/tour/pic"
)

//var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func Pic(dx, dy int) [][]uint8 {
	var board [][]uint8
	var buffer bytes.Buffer
	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
			binary.Write(&buffer, binary.BigEndian, 1)
		}
		board[i] = buffer.Bytes()
	}
	//{
	//	{
	//		1, 1, 1
	//	},
	//	{
	//		1, 1, 1
	//	},
	//	{
	//		1, 1, 1
	//	},
	//}
	//board[0][0] = 2
	//board[2][2] = 2
	//board[1][2] = 2
	//board[1][0] = 2
	//board[0][2] = 2
	return board
}

func main() {

	//var a [2]string
	//a[0] = "Hello"
	//a[1] = "World"
	//fmt.Println(a[0], a[1])
	//fmt.Println(a)
	//primes := [6]int{2, 3, 5, 7, 11, 13}
	//fmt.Println(primes)
	//var s []int = primes[1:4]
	//fmt.Println(s)

	//names := [4]string{"JOHN", "PAUL", "GEORGE", "RINGO"}
	//fmt.Println(names)
	//
	//a := names[0:2]
	//b := names[1:3]
	//fmt.Println(a, b)
	//b[0] = "XXX"
	//fmt.Println(a, b)
	//fmt.Println(names)

	//q := []int{2, 3, 4, 7, 1, 13}
	//fmt.Println(q)
	//
	//r := []bool{true, false, true, true, false, true}
	//fmt.Println(r)
	//
	//s := []struct {
	//	i int
	//	b bool
	//}{
	//	{2, true},
	//	{3, false},
	//	{5, false},
	//}
	//fmt.Println(s)

	//s := []int{1, 2, 3, 4, 5, 6, 7}
	//s = s[1:4]
	//fmt.Println(s)
	//s = s[:2]
	//fmt.Println(s)
	//s = s[1:]
	//fmt.Println(len(s))
	//fmt.Println(cap(s))
	//fmt.Println(s)
	//s = s[1:]
	//fmt.Println(s == nil)
	//fmt.Println(s)
	//fmt.Println(s == nil)
	//fmt.Println(len(s))
	//s = append(s, 0)
	//fmt.Println(s)
	//fmt.Println(len(s))

	//s1 := []int{}
	//fmt.Println(s1 == nil)

	//var s2 []int
	//fmt.Println(len(s2))
	//fmt.Println(s2 == nil)
	//
	//a := make([]int, 5)
	//fmt.Println(a)
	//
	//b := make([]int, 0, 5)
	//fmt.Println(b)
	//fmt.Println(len(b))
	//fmt.Println(b == nil)

	//board := [][]string{
	//	[]string{"_", "_", "_"},
	//	[]string{"_", "_", "_"},
	//	[]string{"_", "_", "_"},
	//}
	//board[0][0] = "X"
	//board[2][2] = "O"
	//board[1][2] = "X"
	//board[1][0] = "O"
	//board[0][2] = "X"
	//fmt.Println(board)
	//for i := 0; i < len(board); i++ {
	//	fmt.Println(strings.Join(board[i], " "))
	//}

	//for i, v := range pow {
	//	fmt.Printf("2**%d=%d\n", i, v)
	//}

	//pow := make([]int, 10)
	//for i := range pow {
	//	pow[i] = 1 << uint(i)
	//}
	//for _, value := range pow {
	//	fmt.Println(value)
	//}
	fmt.Println("--------")
	pic.Show(Pic)
	fmt.Println("=========")
}
