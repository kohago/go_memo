package main

import "fmt"

func testChan()  {
	//testChanNoBufferDeadLock()
	testChanWithBuffer()

	//the good one (without buffer)
	//testChanNoBufferOK1()

}

//after chan1
func testChanWithBuffer()  {
	c1 :=make(chan interface{},1)
	c1 <- struct {

	}{}
	fmt.Println("after chan1")
}

//fatal error: all goroutines are asleep - deadlock!
//chan without buffer will wait until someone receive it
func testChanNoBufferDeadLock()  {
	c1 :=make(chan interface{})
	c1 <- struct {

	}{}
	fmt.Println("after chan1")
}

//the main chan will early finish.
func testChanNoBufferOK1() {
	c1 := make(chan interface{})

	go func() {
		for {
			select {
			case <- c1 :
				fmt.Println("receive!")
			}
		}
	}()

		c1 <- struct {
		}{}
		fmt.Println("after chan1")
}