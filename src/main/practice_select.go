package main

import (
	"fmt"
	"runtime"
)

func PracticeSelect() {
	runtime.GOMAXPROCS(1)

	ic := make(chan int, 1)
	sc := make(chan string, 1)

	ic <- 1
	sc <- "test"

	/**
	when there are more than one case,select will select a case randomly.
	when there is no case with chan,select will wait.
	when there is only one case can be executed,select will run immediately.
	otherwise select will do the default case.
	*/
	select {
	case someInt := <-ic:
		fmt.Printf("receive from ic:%v", someInt)
	case someStr := <-sc:
		panic(someStr)
	}
}
