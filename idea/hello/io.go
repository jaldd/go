package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(out byte) byte {
	switch {
	case out >= 'A' && out <= 'M' || out >= 'a' && out <= 'm':
		out += 13
	case out >= 'N' && out <= 'Z' || out >= 'n' && out <= 'z':
		out -= 13
	}
	return out
}

func (fz rot13Reader) Read(b []byte) (int, error) {

	n, e := fz.r.Read(b)
	for i := 0; i < n; i++ {
		b[i] = rot13(b[i])
	}
	return n, e
}

func main() {
	//r := strings.NewReader("Hello world")
	//b := make([]byte, 8)
	//for {
	//	n, err := r.Read(b)
	//	fmt.Printf("n=%v err=%v b=%v\n", n, err, b)
	//	fmt.Printf("b[:n]=%q\n", b[:n])
	//	if err == io.EOF {
	//		break
	//	}
	//}

	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
