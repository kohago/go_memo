package main

import (
	"errors"
	"fmt"
	"time"
)

func testChanChan() {
	req := make(chan int)
	stopReq := make(chan chan string)
	ticker := time.NewTicker(5 * time.Second)
	c := ticker.C

	count := make(chan int)
	// chan chan
	// send something to another goRoutine using a chan.
	// also want to get the responese of that roution,so
	// send a chan to that routiune for the response!
	stop := make(chan chan error)

	go func() {
		var c int
		for {
			select {
			//when req <-i then v:= <-req
			case v := <-req:
				fmt.Printf("received %d\n\n", v)
				c++
			//when <-count then count<-c
			case count <- c:
				fmt.Printf("put c:%d to count!\n", c)
			//when stop <- then <-stop
			case sch := <-stop:
				//sch is a chan for stop,give back the message
				sch <- errors.New("stop Field!")
			}
		}
	}()

	go func() {
		for i := 0; ; i++ {
			select {
			// maybe can't use the same channle for more than one routine
			case v := <-stopReq:
				fmt.Printf("recive %v signal ,stop sending int to chan req\n", v)
				v <- "ok, i will stop\n"
				return
			default:
				time.Sleep(time.Second)
				fmt.Printf("put %d to req\n", i)
				req <- i
			}
		}
	}()

	go func() {
		<-time.After(7 * time.Second)
		sch := make(chan error)

		stop <- sch
		// if there any response
		if err := <-sch; err != nil {
			fmt.Println(err)
		}

		stopSignal := make(chan string)
		stopReq <- stopSignal
		// if there any response
		response := <-stopSignal
		fmt.Println(response)
		return
	}()

	go func() {
		<-time.After(time.Second * 15)
		ticker.Stop()
	}()

	//here is the main routine
	fmt.Printf("Type:%T,Value:%v\n", c, c)
	for _ = range c {
		fmt.Printf("count=%d\n", <-count)
	}
}
