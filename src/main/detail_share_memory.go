package main

import (
	"context"
	"fmt"
	"sync"
)

/**
critical section 每次只允许一个goroutine进入某个代码块，此代码块区域称之为”临界区(critical section)

*****锁Mutex
var s sync.Mutex
s.Lock()
...
用锁创造了一个临界区，
...
s.Unlock()


*****条件变量Cond
用来自行调度groutine
Cond的使用，离不开互斥锁(初始化时需要一个Mutex
Cond提供Wait、Signal、Broadcast 三种方法。
Wait表示线程(groutine)阻塞等待；
Signal表示唤醒等待的groutine；
Broadcast表示唤醒等待的所有groutine;)


*****原子操作(atomicity)
原子操作是硬件芯片级别的支持，所以可以保证绝对的线程安全。而且执行效率比其他方式要高出好几个数量级。
Go语言的原子操作当然也是基于CPU和操作系统的，Go语言提供的原子操作的包是sync/atomic


*****为了能更好的调度goroutine，Go提供了sync.WaitGroup、sync.Once还有context
var wait sync.WaitGroup

    wait.Add(2) //必须是运行的goroutine的数量

    go func() {
        //TODO 一顿小操作
        defer wait.Done() // done函数用在goroutine中，表示goroutine操作结束
    }()

    go func() {
        //TODO 一顿小操作
        defer wait.Done() // done函数用在goroutine中，表示goroutine操作结束
    }()

    wait.Wait() // block住了，直到所有goroutine都结束

 */

func testShareMemory()  {
	//testWaitAndSyncOnce()
	testCancelContext()
}

 //use waitGroup to keep the main goRoutine alive
 //trans i to goRoutines,because if not i may be changed when the goRoutine begin to run
 //once,really just once!
func testWaitAndSyncOnce()  {
	var once sync.Once
	onceBody:= func() {
		fmt.Println("run once")
	}

	var wait sync.WaitGroup
	wait.Add(10)

	for i:=0;i<10;i++{
		go func(i int) {
			fmt.Println(i)
			once.Do(onceBody)
			defer  wait.Done()
		}(i)
	}

	wait.Wait()
}

/**
context可以用来实现一对多的goroutine协作。
这个包的应用场景主要是在API中。字面意思也很直接，上下文。
当一个请求来时，会产生一个goroutine，但是这个goroutine往往要衍生出许多额外的goroutine去处理操作，例如链接database、请求rpc请求。。等等，
这些衍生的goroutine和主goroutine有很多公用数据的，例如同一个请求生命周期、用户认证信息、token等，
当这个请求超时或者被取消的时候，这里所有的goroutine都应该结束。
context就可以帮助我们达到这个效果。

有两个根context:background和todo；这两个根都是contenxt空的，没有值的。两者也没啥太本质的区别，
Background是最常用的，作为Context这个树结构的最顶层的Context，它不能被取消。
当不知道用啥context的时候就可以用TODO。

根生成子节点有以下方法：

//生成可撤销的Context (手动撤销)
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)

//生成可定时撤销的Context (定时撤销)
func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)

//也是生成可定时撤销的Context (定时撤销)
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)

//不可撤销的Context,可以存一个kv的值
func WithValue(parent Context, key, val interface{}) Context
 */
func testCancelContext()  {
	//gen is func that will return a chan ,when the chan receive a int value
	gen := func(ctx context.Context) <-chan int{
		dst:=make(chan int)
		n:=1

		go func() {
			for{
				select {
					case <-ctx.Done():
						return
					case dst <- n:
						n++
				}
			}
		}()

		return dst
	}

	ctx,cancel:=context.WithCancel(context.Background())
	//cancel is a func that finish ctx.call it will finish the ctx
	defer cancel()

	//range a chan
	//means always  use the chan's received value unit the chan die...!
	//https://gobyexample.com/range-over-channels
	for n:=range gen(ctx){
		//print the value received from gen
		fmt.Println(n)
		if n==5 {
			break;
		}
	}
}
