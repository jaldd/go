package diy

type Diy struct {
	A int64
	b float64
}

func (diy *Diy) Set(i int64, j float64) {
	diy.A = i
	diy.b = j
	return
}

func (diy Diy) Set2(i int64, j float64) {
	diy.A = i
	diy.b = j
	return
}

func (diy Diy) set(i int64, j float64) {
	diy.A = i
	diy.b = j
	return
}

func sum(a, b int64) int64 {
	return a + b
}
