package main

import (
	"fmt"
	"math"
	"time"
)

//factored vars
var (
	MyBool    bool   = false
	YouString string = "You String"
	HisInt    int    = 10
)

const (
	Hello = "Hello"
	World = "World"
	Big   = 1 << 100
	Small = Big >> 99
)

func basic() {
	//const
	fmt.Println(Hello, World)
	fmt.Println(Small)

	// vars
	fmt.Printf("Type:%T value:%v\n", MyBool, MyBool)
	fmt.Printf("Type:%T value:%v\n", YouString, YouString)
	fmt.Printf("Type:%T value:%v\n", HisInt, HisInt)

	var i, you, he, her string = "a", "b", "c", "d"
	var message string = "my name is jama!"
	longMsg := "jama him"

	const nerverChange string = "never change!"
	//nerverChange = "change it!"

	fmt.Println("Hello,World!")
	fmt.Println(message)

	if 2 > 1 {
		fmt.Println("2 is bigger than 1")
	}

	if 1 > 2 {
		fmt.Println("1 is bigger than 2")
	} else if 2 > 3 {
		fmt.Println("2 is bigger than 3")
	} else {
		fmt.Println("3 is the biggest")
	}

	fmt.Printf("%s,%s,%s,%s:%s\n", i, you, he, her, longMsg)

	for i := 0; i < 10; i++ {
		fmt.Printf("i is %d\n ", i)
		if i == 5 {
			break
		}
	}

	sum := 1
	for sum < 100 {
		sum += sum
	}
	fmt.Println("the sum is", sum)

	n := 10
	switch n {
	case 10:
		fmt.Println("n is 10")
	default:
		fmt.Println("n is the default value")
	}

	fmt.Printf("the time is", time.Now(), "\n")
	fmt.Println(math.Pi)

	fmt.Println(add(10, 29))

	fmt.Println(minus(30, 10))

	x, y := swap("hello", "world")
	fmt.Println(x, y)
	fmt.Println("x:", 10, " y:", 20, "after swap")

	fmt.Println(split(17))

	//statement before if
	fmt.Println(pow(3, 2, 10))
	fmt.Println(pow(3, 3, 26))

	//defer
	defer fmt.Println("defer")
	fmt.Println("hello")
}

//return value
func add(x int, y int) int {
	return x + y
}

//paramater type use one type
func minus(x, y int) int {
	return x - y
}

//double return values
func swap(x, y string) (string, string) {
	return y, x
}

//naked return values
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - y
	return
}

// if expression
func pow(x, y, lim float64) float64 {
	if v := math.Pow(x, y); v < lim {
		return v
	}
	return lim
}
