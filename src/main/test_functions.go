package main

import (
	"fmt"
	"math"
)

func testFunction() {
	//testFunctionValue()

	//testClosure()

	//practiceFibonacci()

	testReceiver()

}

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) PointerScale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func testReceiver() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	v.Scale(10)
	fmt.Println("receiver:", v.Abs())

	v.PointerScale(10)
	fmt.Println("pointerReceiver:", v.Abs())
}

func practiceFibonacci() {
	f := fib()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

//practice fibonacci
func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

//test closure
func testClosure() {
	aClosure, bClosure := someClousure(), someClousure()
	for i := 0; i < 10; i++ {
		fmt.Println(aClosure(i), bClosure(-2*i))
	}
}

func someClousure() func(int) int {
	sum := 0
	return func(x int) int {
		sum = sum + x
		return sum
	}
}

//test Function value
func testFunctionValue() {
	myFunc := func(x, y float64) float64 {
		return x * y
	}

	fmt.Println(myFunc(2, 3))
	fmt.Println(compute(myFunc))
	fmt.Println(compute(math.Pow))
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(4, 5)
}
