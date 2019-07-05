package main

import "fmt"

var (

	/**
	compile error
	: is not necessary with var
	var1 := 1
	var2 := var1 * 2
	*/

	var1 = 1
	var2 = var1 * 2
)

const const1 = 1

var var3 = 3

func PracticeVar() {
	fmt.Println("var1:%v,var2:%v", var1, var2)

	fmt.Println(&var3, var3)
	/**
	can't take address of const,because const is not a var(?? of course),const will get its memory before run
	*/
	//fmt.Println(&const1, const1)
}
