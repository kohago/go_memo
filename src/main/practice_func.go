package main

import "fmt"

func PracticeFunc() {
	fmt.Println(childFunc1(1))
	fmt.Println(childFunc2(1))
}

func childFunc1(i int) int {
	t := i
	func() {
		t += 3
	}()
	return t
}

// there is some special scope about defer
// defer will copy the local variable from outside func
// defer will use the named return variable directly
func childFunc2(i int) int {
	t := i
	defer func() {
		t += 3
	}()
	return t
}
