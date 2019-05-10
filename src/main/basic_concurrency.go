package main

import (
	"fmt"
	"math/rand"
	"time"
)

func testConcurrency() {
	//simpleBoring()

	////boring thing is something to ignore!
	//testRoutineBoring()

	//testChanBoring()

	//testFuncReturnsChan()

	//testConCurFunChan()

	//testConCurFunChanWitchSelect()

	//testTimeOut()

	testDaisyChain()

	fmt.Println("You are boring,I leaving.")
}

func testDaisyChain() {
	const n = 3
	leftmost := make(chan int)
	left := leftmost
	right := leftmost
	fmt.Printf("leftmost: %T,%v\n", leftmost, leftmost)
	fmt.Printf("left: %T,%v\n", left, left)
	fmt.Printf("right:%T,%v\n", right, right)
	fmt.Println("")

	//one go routine is one person,has two chan
	//firstly create the queue from left to right
	//then,kick it from right to left
	//just like domino

	//1:   chan1(left,right,leftmost)
	//2:   chan2(right)
	//3:   chan1<-chan2
	//4:   chan2(right,left)
	//5:   chan3(right)
	//    ....
	//    chanx<-1
	for i := 0; i < n; i++ {
		fmt.Println("----", i, "begin---")
		right = make(chan int)
		fmt.Printf("right after make: %T,%v\n", right, right)
		go messageSomeone(left, right)
		left = right
		fmt.Printf("left after change: %T,%v\n", left, left)
		fmt.Printf("right after change:%T,%v\n", right, right)
		fmt.Println("----", i, "end---")
		fmt.Println("")
	}

	time.Sleep(n * time.Second)
	right <- 1

	fmt.Println(<-leftmost)
}
func messageSomeone(left, right chan int) {
	fmt.Println("")
	fmt.Println("**************--------************")
	fmt.Printf("left in routine: %T,%v\n", left, left)
	fmt.Printf("right in routine:%T,%v\n", right, right)
	fmt.Println("**************--------************")
	left <- 1 + <-right
}

func testTimeOut() {
	c := funcChanBoring("jone")
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-time.After(4 * time.Second):
			fmt.Println("You're too slow")
			return
		}
	}
}

func testConCurFunChanWitchSelect() {
	testConCurc := conCurFunChanWithSelect(funcChanBoring("jone boring "), funcChanBoring("ann boring"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-testConCurc)
	}
}

func conCurFunChanWithSelect(input1, input2 <-chan string) <-chan string {
	conCurC := make(chan string)
	//<-input1: things getted from input1,Name it as I1
	//c <-I1 :put things from I1 to c
	go func() {
		// let it live until the main is over using for
		for {
			select {
			//the second : is the needed by case
			case s := <-input1:
				conCurC <- s
			case s := <-input2:
				conCurC <- s
			}

		}
	}()
	return conCurC
}

func testConCurFunChan() {
	testConCurc := conCurFunChan(funcChanBoring("jone boring "), funcChanBoring("ann boring"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-testConCurc)
	}
}

func conCurFunChan(input1, input2 <-chan string) <-chan string {
	conCurC := make(chan string)
	go func() {
		// let it live until the main is over using for
		for {
			//<-input1: things from input1,Name it as I1
			//c <-I1 :put things from I1 to c
			conCurC <- <-input1
		}
	}()
	go func() {
		for {
			conCurC <- <-input2
		}
	}()
	return conCurC
}

//just get a chan, get some message from it for 5 times
//juse like a service
func testFuncReturnsChan() {
	jone := funcChanBoring("Jone boring!")
	ann := funcChanBoring("Ann boring")
	for i := 0; i < 5; i++ {
		fmt.Printf("jone say %q\n", <-jone)
		fmt.Printf("Ann say %q\n", <-ann)
	}
}

//create a chan which receive some msg,then return it
func funcChanBoring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e4)) * time.Millisecond)
		}
	}()

	return c
}

func testChanBoring() {
	c := make(chan string)
	go chanBoring("boring", c)

	for i := 0; i < 5; i++ {
		fmt.Printf("You say %q\n", <-c)
	}

}
func chanBoring(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func testRoutineBoring() {
	go simpleBoring()
	fmt.Println("I am listening!")
	//just listening for 2 seconds
	time.Sleep(2 * time.Second)
}

func simpleBoring() {
	boring("boring!")
}

func boring(msg string) {
	for i := 0; i < 10; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
