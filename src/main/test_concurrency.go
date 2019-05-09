package main

import (
	"fmt"
	"math/rand"
	"time"
)
/**
Go实现了两种并发形式。第一种是大家普遍认知的：多线程通过共享内存来实现协作。其实就是Java或者C++等语言中的多线程开发。
衍生出了，threadSafe的数据结构
另外一种是Go语言特有的，也是Go语言推荐的：CSP（communicating sequential processes）并发模型。以通信来实现协作

Do not communicate by sharing memory; instead, share memory by communicating.

Go的CSP并发模型，是通过goroutine和channel来实现的。


无论语言层面何种并发模型，到了操作系统层面，一定是以线程的形态存在的。
而操作系统根据资源访问权限的不同，体系架构可分为用户空间和内核空间；
内核空间主要操作访问CPU资源、I/O资源、内存资源等硬件资源，为上层应用程序提供最基本的基础资源，
用户空间呢就是上层应用程序的固定活动空间，用户空间不可以直接访问资源，
必须通过“系统调用”、“库函数”或“Shell脚本”来调用内核空间提供的资源。

我们现在的计算机语言，可以狭义的认为是一种“软件”，
它们中所谓的“线程”，往往是用户态的线程，和操作系统本身内核态的线程（简称KSE），还是有区别的。
USE:KSE=M:1 (多个用户态的线程对应着一个内核线程，程序线程的创建、终止、切换或者同步等线程工作必须自身来完成)
USE:KSE=1:1 (直接调用操作系统的内核线程，所有线程的创建、终止、切换、同步等操作，都由内核来完成。C++就是这种)
USE:KSE=M:N (先创建多个内核级线程，然后用自身的用户级线程去对应创建的多个内核级线程，自身的用户级线程需要本身程序去调度，内核级的线程交给操作系统内核去调度。goLang是这一种)
KSE-Machine-Processor-GoRoutine
http://morsmachine.dk/go-scheduler

一个M会对应一个内核线程，一个M也会连接一个上下文P，一个上下文P相当于一个“处理器”，一个上下文连接一个或者多个Goroutine。
P(Processor)的数量是在启动时被设置为环境变量GOMAXPROCS的值，或者通过运行时调用函数runtime.GOMAXPROCS()进行设置。
Processor数量固定意味着任意时刻只有固定数量的线程在运行go代码。
Goroutine中就是我们要执行并发的代码。待执行状态的Goroutine形成了队列runqueues

要上下文的目的---当遇到内核线程(KSE)阻塞的时候,M(M0)可以通过直接放开P及依附在P上的其他线程，交给其它M(M1)去执行。
  无事可做的P会定期检查全局runqueue,去执行可被执行的Goroutine。当核线程(KSE)阻塞结束时，M0会获得一个，得以被执行。

每个P中的Goroutine不同导致他们运行的效率和时间也不同，在一个有很多P和M的环境中，不能让一个P跑完自身的Goroutine就没事可做了，
因为或许其他的P有很长的goroutine队列要跑，得需要均衡。
该如何解决呢？Go的做法倒也直接，从其他P中偷一半！


*/

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
