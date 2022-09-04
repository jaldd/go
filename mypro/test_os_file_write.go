package main

import (
	"os"
)

func write() {

	f, _ := os.OpenFile("b.txt", os.O_RDWR, 0755)
	f.Write([]byte("helloworld"))
	f.Close()
}

func main() {

	write()
}
