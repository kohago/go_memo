package main

import "fmt"

//will fill up the array with zero value
/**
[0 0 0 0 0]
[0 0 0 0 0 1 2 3]

[]
[1 2 3]
*/

func PracticeMakeAndNew() {
	practiceMake()
	practiceNew()
}

func practiceMake() {
	m := make([]int, 5)
	fmt.Println(m)
	m = append(m, 1, 2, 3)
	fmt.Println(m)

	m1 := make([]int, 0)
	fmt.Println(m1)
	m1 = append(m1, 1, 2, 3)
	fmt.Println(m1)
}

/**
// The new built-in function allocates memory. The first argument is a type,
// not a value, and the value returned is a pointer to a newly
// allocated zero value of that type.
func new(Type) *Type
*/
func practiceNew() {
	alist := new([]int)
	//alist = append(alist, 1)
	fmt.Println(alist)
}
