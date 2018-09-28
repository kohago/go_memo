package main

import (
	"fmt"
	"math"
)

func testFunction() {
	//testFunctionValue()

	//testClosure()

	//practiceFibonacci()

	//testReceiver()

	testInterface()
	//testTypeSwitches()
}

func testTypeSwitches() {
	do(21)
	do("hello")
	do(true)
}
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I can't know about type %T\n", v)
	}
}

func testInterface() {

	var theInterface TestInterface

	//Type assertions
	var i interface{} = "hello"
	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	testF, ok := i.(float64)
	fmt.Println(testF, ok)

	testF = i.(float64)
	fmt.Println(testF)

	//testEmptyIF
	//you can put any value into empty interface
	var emptyIf interface{}
	emptyIf = 42
	fmt.Printf("(%v,%T\n)", emptyIf, emptyIf)
	emptyIf = "test"
	fmt.Printf("(%v,%T\n)", emptyIf, emptyIf)

	//test null value
	var theNilIf TestInterface
	fmt.Printf("(%v,%T\n)", theNilIf, theNilIf)

	f := Myfloat(-math.Sqrt2)
	v := Vertex{3, 4}

	theInterface = f
	fmt.Printf("(%v,%T)\n", theInterface, theInterface)

	theInterface = &v
	fmt.Printf("(%v,%T)\n", theInterface, theInterface)

	theInterface = v
	fmt.Printf("(%v,%T)\n", theInterface, theInterface)

	fmt.Println(theInterface.Abs())
}

type TestInterface interface {
	Abs() float64
}

type Myfloat float64

func (f Myfloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	} else {
		return float64(f)
	}
}

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//the merit of pointer function
//.can change the value of the parameter
//.no necessary to copy the parameter
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
