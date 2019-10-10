package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	ch := make(chan int, 1)
	ch <- 100

	aInt := <-ch
	fmt.Println(aInt)

	cn := asChannel(1, 2, 3, 4, 5, 6)
	c2 := doubleVALUE(cn)
	fmt.Println("channel value : ")
	for v := range c2 {
		fmt.Println(v, " ")
	}

}

func asChannel(vs ...int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, v := range vs {
			out <- v
			time.Sleep(10 * time.Millisecond)
		}
	}()
	return out
}

func doubleVALUE(c chan int) chan int {
	va := make(chan int)
	go func() {
		defer close(va)
		for v := range c {
			va <- v * 2
			time.Sleep(10 * time.Millisecond)
		}
	}()
	return va
}

func doSomething() {
	time.Sleep(2 * time.Second)
	fmt.Println("do some thing")
}

func thinkSomething(wg *sync.WaitGroup) {
	time.Sleep(2 * time.Second)
	fmt.Println("think some thing")
	wg.Done()
}

func thinkSomething2(wg *sync.WaitGroup) {
	time.Sleep(2 * time.Second)
	fmt.Println("think some thing")
	wg.Done()
}

func thinkSomething3(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(2 * time.Second)
	fmt.Println("think some thing")
}
