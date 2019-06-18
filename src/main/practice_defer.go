package main

import "fmt"

func PracticeDefer() {
	//PracticeDeferPanic()
	//PraticeDeferFuncParma()
	PracticeDeferReturnValue()
}

//panic will be executed after defer;
//defer like stack, is LIFO
//so the answer is
/**
three
two
one
there is a panic
*/
func PracticeDeferPanic() {
	defer func() { fmt.Println("one") }()
	defer func() { fmt.Println("two") }()
	defer func() { fmt.Println("three") }()

	panic("there is a panic")
}

func calc(name string, a, b int) int {
	ret := a + b
	fmt.Println(name, a, b, ret)
	return ret
}

//defer will be executed later.but func in defer will be executed immediately
/**
10 1 2 3
20 0 2 2
2 0 2 2
1 1 3 4
*/
func PraticeDeferFuncParma() {
	a, b := 1, 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}

//defer will be executed before func return
//return value name has the whole func scope(include the child func)!!!
/**
4
5
8
*/
func PracticeDeferReturnValue() {
	fmt.Println(func1(1))
	fmt.Println(func2(5))
	fmt.Println(func3(6))
}

func func1(i int) (t int) {
	t = i
	defer func() {
		t += 3
	}()
	return t
}

func func2(i int) int {
	t := i
	defer func() {
		t += 3
	}()
	return t
}

func func3(i int) (t int) {
	defer func() {
		t += i
	}()
	return 2
}
