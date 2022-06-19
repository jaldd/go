package main

import (
	"fmt"
	"reflect"
)

func main() {

	fmt.Println("hello")

	//p := true
	//a := 3
	//b := 6.0
	//c := "hi"
	//d := [3]string{"1", "2", "3"}
	//e := []int64{1, 2, 3}
	//f := map[string]int64{"a": 3, "b": 4}
	//fmt.Printf("type:%T:%v\n", p, p)
	//fmt.Printf("type:%T:%v\n", a, a)
	//fmt.Printf("type:%T:%v\n", b, b)
	//fmt.Printf("type:%T:%v\n", c, c)
	//fmt.Printf("type:%T:%v\n", d, d)
	//fmt.Printf("type:%T:%v\n", e, e)
	//fmt.Printf("type:%T:%v\n", f, f)
	//
	//e[0] = 9
	//e = append(e, 3)
	//f["f"] = 5
	//v, ok := f["f"]
	//fmt.Println(v, ok)
	//v, ok = f["ff"]
	//fmt.Println(v, ok)
	//
	//if a > 0 {
	//	fmt.Println("a>0")
	//} else {
	//	fmt.Println("a<=0")
	//}
	//a = 0
	//for {
	//	if a >= 10 {
	//		fmt.Println("out")
	//		break
	//	}
	//
	//	a = a + 1
	//	if a > 5 {
	//		continue
	//	} else {
	//		fmt.Println(a)
	//	}
	//}
	//
	//for i := 9; i <= 10; i++ {
	//	fmt.Printf("i=%d\n", i)
	//}
	//
	//for k, v := range e {
	//	fmt.Println(k, v)
	//}
	//
	//for k, v := range f {
	//	fmt.Println(k, v)
	//}
	//var h, i int64
	//sum := sum(h, i)
	//fmt.Printf("sum(h+i),h=%v,i=%v,%v\n", h, i, sum)
	//
	//g := diy.Diy{A: 2}
	//fmt.Printf("type:%T:%v\n", g, g)
	//g.Set(1, 1)
	//fmt.Printf("type:%T:%v\n", g, g)
	//
	//g.Set2(3, 3)
	//fmt.Printf("type:%T:%v\n", g, g)
	//
	//m := new(diy.Diy)
	//m.A = 2
	//fmt.Printf("type:%T:%v\n", m, m)
	//m.Set(1, 1)
	//fmt.Printf("type:%T:%v\n", m, m)
	//m.Set2(3, 3)
	//fmt.Printf("type:%T:%v\n", m, m)
	//
	//k := &diy.Diy{A: 2}
	//fmt.Printf("type:%T:%v\n", k, k)
	//k.Set(1, 1)
	//fmt.Printf("type:%T:%v\n", k, k)
	//k.Set2(3, 3)
	//fmt.Printf("type:%T:%v\n", k, k)
	//
	//s := make([]int64, 5)
	//s1 := make([]int64, 0, 5)
	//m1 := make(map[string]int64, 5)
	//m2 := make(map[string]int64)
	//fmt.Printf("%#v,cap:%#v,len:%#v\n", s, cap(s), len(s))
	//fmt.Printf("%#v,cap:%#v,len:%#v\n", s1, cap(s1), len(s1))
	//fmt.Printf("%#v,len:%#v\n", m1, len(m1))
	//fmt.Printf("%#v,len:%#v\n", m2, len(m2))
	//
	//var ll []int64
	//fmt.Printf("%#v\n", ll)
	//ll = append(ll, 1)
	//fmt.Printf("%#v\n", ll)
	//ll = append(ll, 2, 3, 4, 5, 6)
	//fmt.Printf("%#v\n", ll)
	//ll = append(ll, []int64{7, 8, 9}...)
	//fmt.Printf("%#v\n", ll)
	//
	//fmt.Println(ll[0:2])
	//fmt.Println(ll[:2])
	//fmt.Println(ll[0:])
	//fmt.Println(ll[:])

	//a := 5
	//b := 6
	//fmt.Println("a:", a)
	//c := &a
	//fmt.Println("a 地址", c)
	//
	//*c = 4
	//fmt.Println("a:", a)
	//
	//c = &b
	//fmt.Println("b 地址", c)
	//
	//*c = 8
	//fmt.Println("a:", a)
	//fmt.Println("b:", b)
	//
	//c1 := c
	//fmt.Println("c 地址", c)
	//fmt.Println("c1地址", c1)
	//
	//*c = 9
	//fmt.Println("c", *c)
	//fmt.Println("c1", *c1)
	//
	//c = &a
	//fmt.Println("c 地址", c)
	//fmt.Println("c1地址", c1)

	var a interface{}
	a = 2
	fmt.Printf("%T,%v\n", a, a)

	print(a)
	print(3)
	print("ldd")

	v, ok := a.(int)
	if ok {
		fmt.Printf("a:%d\n", v)
	}

	switch a.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("...")
	}

	t := reflect.TypeOf(a)
	fmt.Printf("a type:%s", t.Name())
}

func print(i interface{}) {

	fmt.Println(i)
}
func sum(h int64, i int64) int64 {
	return h + i
}
