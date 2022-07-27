package main

import (
	"fmt"
	"sync"
	"time"
)

type Money struct {
	lock   sync.Mutex
	amount int64
}

func (m *Money) Add(i int64) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.amount += i
}

func (m *Money) Minute(i int64) {
	m.lock.Lock()
	defer m.lock.Unlock()
	if m.amount >= i {
		m.amount -= i
	}
}

func (m *Money) Get() int64 {
	return m.amount
}

func main() {

	m := new(Money)
	m.Add(2000)
	for i := 0; i < 1000; i++ {
		go func() {
			m.Minute(1)
		}()
	}
	time.Sleep(2 * time.Second)
	fmt.Println(m.Get())
}
