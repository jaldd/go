package main

import (
	"fmt"
	"os"
)

func createFile() {
	f, err := os.Create("a.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(f.Name())
	}

}
func createDir() {
	err := os.MkdirAll("test/test1", os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
}
func remove() {
	// err := os.Remove("a.txt")
	err := os.RemoveAll("test")
	if err != nil {
		fmt.Println(err)
	}
}
func getWd() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dir)
	os.Chdir("/usr/local")
	dir, err = os.Getwd()
	fmt.Println(dir)
	s := os.TempDir()
	fmt.Println(s)
}

func rename() {
	err := os.Rename("a.txt", "b.txt")
	if err != nil {
		fmt.Println(err)
	}
}
func writeFile() {
	os.WriteFile("b.txt", []byte("hello"), os.ModeAppend)
}

func readFile() {
	b2, _ := os.ReadFile("b.txt")
	fmt.Println(b2)
	fmt.Println(string(b2[:]))
}
func main() {
	writeFile()
	readFile()

}
