package main

import "fmt"

const (
	x = iota
	y
	z = "zz"
	k
	p = iota
)

/**
0 1 zz zz 4
*/
/**
start from 0
when int plus 1
when string, the same value
*/
func PracticeIota() {
	fmt.Println(x, y, z, k, p)
}
