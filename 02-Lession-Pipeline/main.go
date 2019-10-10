package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)
	go thinkSomething(ch)
	go thinkSomething2(ch)
	go thinkSomething3(ch)

	<-ch
	<-ch
	<-ch
}

func doSomething() {
	time.Sleep(2 * time.Second)
	fmt.Println("do some thing")
}

func thinkSomething(ch chan int) {
	time.Sleep(2 * time.Second)
	fmt.Println("think some thing")
	ch <- 1
}

func thinkSomething2(ch chan int) {
	time.Sleep(2 * time.Second)
	fmt.Println("think some thing")
	ch <- 1

}

func thinkSomething3(ch chan int) {
	time.Sleep(2 * time.Second)
	fmt.Println("think some thing")
	ch <- 1
}
