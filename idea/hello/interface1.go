package main

import (
	"fmt"
	"strconv"
)

//type MyFloat1 float64
//
//func (f MyFloat1) Abs() float64 {
//
//	if f < 0 {
//		return float64(-f)
//	}
//	return float64(f)
//}
//
//type Abser interface {
//	Abs() float64
//}
//
//type V2 struct {
//	X, Y float64
//}
//
//func (v *V2) Abs() float64 {
//	return math.Sqrt(v.X*v.X + v.Y*v.Y)
//}
//
//type I interface {
//	M()
//}
//
//type T struct {
//	S string
//}
//
//func (t T) M() {
//	fmt.Println(t.S)
//}
//
//type T1 struct {
//	S string
//}
//
//func (t *T1) M() {
//	if t == nil {
//		fmt.Println("nil")
//		return
//	}
//	fmt.Println(t.S)
//}
//
//type F float64
//
//func (f F) M() {
//	fmt.Println(f)
//}

func main() {

	//var a Abser
	//f := MyFloat1(-math.Sqrt2)
	//v := V2{3, 4}
	//a = f
	//a = &v
	//
	////a = v
	//fmt.Println(a.Abs())
	//
	//var i I = T{"hello"}
	//i.M()
	//var i I
	//var t *T1
	//i = t
	//i.M()
	//i = &T1{"hello"}
	//fmt.Printf("%v %t\n", i, i)
	//i.M()
	//var i interface{}
	//
	//fmt.Printf("%v %t\n", i, i)
	//i = 42
	//
	//fmt.Printf("%v %t\n", i, i)
	//i = "hello"
	//
	//fmt.Printf("%v %t\n", i, i)
	//s := i.(string)
	//fmt.Println(s)
	//s, ok := i.(string)
	//fmt.Println(s, ok)
	//f, ok := i.(float64)
	//fmt.Println(f, ok)
	//f := i.(float64)
	//fmt.Println(f)
	//do(21)
	//do("hello")
	//do(true)
	//a := Person{"zb", 22}
	//b := Person{"ldd", 23}
	//fmt.Println(a, b)
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Println(name, ip)
	}
}

//func do(i interface{}) {
//	switch v := i.(type) {
//	case int:
//		fmt.Printf("%v", v)
//	case string:
//		fmt.Printf("%q %v", v, v)
//	default:
//		fmt.Printf("%T %v", v, v)
//	}
//}
//type Person struct {
//	Name string
//	Age  int
//}
//
//func (p Person) String() string {
//	return fmt.Sprintf("%v (%v years)\n", p.Name, p.Age)
//}
type IPAddr [4]byte

func (v IPAddr) String() string {

	var s string
	for index, val := range v {
		if index < 3 {
			s += strconv.Itoa(int(val)) + "."
		} else {
			s += strconv.Itoa(int(val))
		}
	}
	return s
}
