package main

import (
	"errors"
	"fmt"
	"time"
)

func testChanChan() {
	req := make(chan int)
	count := make(chan int)
	//chan chan
	// send something to another goRoutine using a chan.
	// also want to get the responese of that roution,so
	// send a chan to that routiune for the response!
	stop := make(chan chan error)

	go func() {
		var c int
		for {
			select {
			case v := <-req:
				fmt.Printf("received %d\n", v)
				c++
			case count <- c:
			case sch := <-stop:
				//sch is a chan for stop,give back the message
				sch <- errors.New("stop Field!")
				//time.Sleep(2 * time.Second)
				return
			}
		}
	}()

	go func() {
		for i := 0; ; i++ {
			time.Sleep(time.Second)
			req <- i
		}
	}()

	go func() {
		<-time.After(15 * time.Second)
		sch := make(chan error)
		stop <- sch
		// if there any response
		if err := <-sch; err != nil {
			fmt.Println(err)
		}
	}()

	c := time.Tick(5 * time.Second)
	for _ = range c {
		fmt.Printf("count=%d\n", <-count)
	}
}
