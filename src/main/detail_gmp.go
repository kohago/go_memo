package main

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

