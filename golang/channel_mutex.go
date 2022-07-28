package main

import (
	"fmt"
	"time"
)

/*
 使用channel 实现互斥锁
 1. 实现tryLock 功能
 2. 实现tryWithTimeOut功能
*/

func a(c ChanMutex, index int) {
	c.Lock()
	defer c.UnLock()
	fmt.Println(index)
}

func b(c ChanMutex, index int) {
	ret := c.TryLock()
	if ret == false {
		fmt.Println("not lock")
		return
	}
	fmt.Println(index)
}

func c(c ChanMutex, index int) {
	ret := c.TryWithTimeout(time.Second)
	if ret == false {
		fmt.Println("not lock")
		return
	}
	defer c.UnLock()
	fmt.Println(index)

}

type ChanMutex struct {
	c chan struct{}
}

func (m ChanMutex) Lock() {
	m.c <- struct{}{}
}

func (m ChanMutex) UnLock() {
	<-m.c
}

func (m ChanMutex) TryLock() bool {
	select {
	case m.c <- struct{}{}:
		return true
	default:
		return false
	}
}

func (m ChanMutex) TryWithTimeout(t time.Duration) bool {
	ticker := time.NewTimer(t)
	select {
	case m.c <- struct{}{}:
		return true
	case <-ticker.C:
		return false
	}
}

func NewChanMutex() ChanMutex {
	c := make(chan struct{}, 1)
	cM := ChanMutex{
		c: c,
	}
	return cM
}

func main() {

	cM := NewChanMutex()

	for i := 0; i < 10; i++ {
		go c(cM, i)
	}
	time.Sleep(time.Second)
}
