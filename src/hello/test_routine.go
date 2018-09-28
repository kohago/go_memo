package main

import (
	"fmt"
	"sync"
	"time"
)

func testRoutine() {
	//testSayHello()
	//testChannels()
	//testSelect()
	testMutex()
}

func testMutex() {
	c := SafeCounter{
		v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}
	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}

type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	c.v[key]++
	c.mux.Unlock()
}

func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

func testSelect() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			//will stop here to communicate with c that has benn blcked!
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		//block the c channel. After x are pushed to c, c can be used
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
func testSayHello() {
	go say("slowly world!")
	say("hello")
}

func say(s string) {
	for i := 0; i < 3; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(s)
	}
}

func testChannels() {
	s := []int{1, 2, 3, 4, 5, 6, 7}
	//with buffer,when over buffer will be error!
	c := make(chan int, 100)

	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}
