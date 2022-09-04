package main

import (
	"os"
)

func write() {

	f, _ := os.OpenFile("b.txt", os.O_RDWR|os.O_APPEND, 0755)
	f.Write([]byte("helloworld"))
	f.WriteString("hello world")
	f.Close()
	f, _ = os.OpenFile("b.txt", os.O_RDWR, 0755)
	f.WriteAt([]byte("ldd 6"), 5)
	f.Close()
}

func main() {

	write()
}
