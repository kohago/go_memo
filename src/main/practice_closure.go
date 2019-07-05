package main

import "fmt"

func PracticeClosure() {
	funcs := testClosureVarible1()
	for i := range funcs {
		funcs[i]()
	}

	funcs = testClosureVarible2()
	for i := range funcs {
		funcs[i]()
	}

	/**
	1000
	1010
	*/
	func1, func2 := testClosure3(1000)
	func1()
	func2()
}

/**
2
2
closure will be create first and executed later.
when executed , i's value has been changed.
*/
func testClosureVarible1() []func() {
	var funcs []func()
	for i := 0; i < 2; i++ {
		funcs = append(funcs, func() {
			fmt.Printf("&i:%v,i:%v,%T\n", &i, i, i)
		})
	}

	return funcs
}

/**
1
2
closure will be create first and executed later.
when executed , i's value has benn changed.
but before it changed, copy it's value to x.
*/
func testClosureVarible2() []func() {
	var funcs []func()
	for i := 0; i < 2; i++ {
		x := i
		funcs = append(funcs, func() {
			fmt.Printf("&i:%v,i:%v,x's Type:%T\n", &i, i, i)
			fmt.Printf("&x:%v,x:%v,x's Type:%T\n", &x, x, x)
		})
	}

	return funcs
}

func testClosure3(x int) (func(), func()) {
	return func() {
			fmt.Println(x)
			x += 10
		}, func() {
			fmt.Println(x)
		}
}
