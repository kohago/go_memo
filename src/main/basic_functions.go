package main

import (
	"fmt"
	"math"
)

func testFunction() {
	testFunctionValue()

	//testClosure()

	//practiceFibonacci()

	//testReceiver()

	//testInterface()

	//testTypeSwitches()
}

// When return value named, they are initialized to the zero values for their types when the function begins;
//-->good!
// func some() (t int,err error)

//testTypeSwitches
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

//testInterface
func testInterface() {

	var theInterface TestInterface
	//testEmptyIF
	//you can put any value into empty interface  interface{}
	//Empty interfaces are used by code that handles values of unknown type
	var emptyIf interface{}
	emptyIf = 42
	fmt.Printf("(%v,%T)\n", emptyIf, emptyIf)
	emptyIf = "test"
	fmt.Printf("(%v,%T)\n", emptyIf, emptyIf)

	//test null value
	var theNilIf TestInterface
	//(<nil>,<nil>)--> nil is the type nil.nil is a type
	fmt.Printf("(%v,%T)\n", theNilIf, theNilIf)

	f := Myfloat(-math.Sqrt2)
	v := Vertex{3, 4}

	theInterface = f
	fmt.Printf("(%v,%T)\n", theInterface, theInterface)

	theInterface = &v
	fmt.Printf("(%v,%T)\n", theInterface, theInterface)

	theInterface = v
	fmt.Printf("(%v,%T)\n", theInterface, theInterface)

	fmt.Println(theInterface.Abs())

	//Type assertions==>assert an empty interface to some type
	var i interface{} = "hello"
	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	testF, ok := i.(float64)
	//0 false-->when failed will be the default value 0
	//there is no panic
	fmt.Println(testF, ok)

	//will cause a panic,code behind will not executed!
	testF = i.(float64)
	fmt.Println(testF)
}

// like java's interface ,have some signatures
// a method is just a function with a receiver argument.
//If the concrete value inside the interface itself is nil, the method will be called with a nil receiver
type TestInterface interface {
	Abs() float64
}

type Myfloat float64

//A type implements an interface by implementing its methods. There is no explicit declaration of intent, no "implements" keyword.
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

//(v Vertex) is receiver.
// like java's class method
// Go does not have classes. However, you can define methods on types.
//A method is a function with a special receiver argument.
//A method is just a function with a receiver argument.
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//the merit of pointer function
// because of pointer's byRef
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
	fmt.Println("----test receiver--")
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	v.Scale(10)
	fmt.Println("receiver default is byValue,not changed", v.Abs())

	v.PointerScale(10)
	fmt.Println("pointerReceiver changed!:", v.Abs())
}

func practiceFibonacci() {
	fmt.Println("--test fibonacci---")
	f := fib()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

//practice fibonacci
//this is a closure with a,b and one function
//the function returns a function---> func fib()--> func() int
func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b //a=b b=a+b
		return a
	}
}

//test Function value
//function can be a var
func testFunctionValue() {
	fmt.Println("--test function var--")
	myFunc := func(x, y float64) float64 {
		return x * y
	}

	fmt.Println(myFunc(2, 3))
	fmt.Println(compute(myFunc))
	//math.Pow-->x,y={x*...x} y'th x
	fmt.Println(compute(math.Pow))
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(4, 5)
}

//test closure
//A closure is a function value that references variables from outside its body
//A closure has a value that can interact with it's function
func testClosure() {
	fmt.Println("--test closure--")
	aClosure, bClosure := someClousure(), someClousure()
	for i := 0; i < 10; i++ {
		fmt.Println(aClosure(i), bClosure(-2*i))
	}
}

// function can be reture value type
func someClousure() func(int) int {
	sum := 0
	return func(x int) int {
		sum = sum + x
		return sum
	}
}
