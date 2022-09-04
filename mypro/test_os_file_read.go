package main

import (
	"fmt"
	"os"
)

func openClose() {
	// f, _ := os.Open("b.txt")

	f, _ := os.OpenFile("b1.txt", os.O_RDWR|os.O_CREATE, 755)
	fmt.Println(f.Name())
	f.Close()
}

func readOps() {
	f, _ := os.Open("b.txt")
	// f.Seek(3,0)
	// for {

	// 	buf := make([]byte, 10)
	// 	n, err := f.Read(buf)
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	fmt.Println(n)
	// 	fmt.Println(string(buf))
	// }
	// f.Close()

	// buf := make([]byte, 3)
	// n, _ := f.ReadAt(buf, 1)
	// fmt.Println("n:")
	// fmt.Println(n)
	// fmt.Println(string(buf))

	de, _ := os.ReadDir("test")
	for _, v := range de {
		fmt.Println(v.IsDir())
		fmt.Println(v.Name())
	}
	f.Close()
}

func main() {

	readOps()
}
