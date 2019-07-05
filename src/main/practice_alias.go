package main

import "fmt"

func PracticeAlias() {
	//MyInt1 is a new type. 定義なのでdefition
	type MyInt1 int
	//MyInt2 is a alias of int
	type Myintt2 = int

	i := 1

	/**
	MyInt1 is a new type so can't assign int to it
	*/
	//var i1 MyInt1 = i

	/**
	should cast int to Myint1
	*/
	var i1 MyInt1 = MyInt1(i)

	/**
	can put int value to Myint2,because it is a alias
	*/
	var i2 Myintt2 = i

	fmt.Println(i, i1, i2)

	var u1 MyUser1
	var u2 MyUser2

	u1.echo1()
	/**
	because MyUser2 is MyUser's alias,so can use user's method directly
	*/
	u2.echo2()

	/**
	Myuser1 has no echo2 method
	*/
	//u1.echo2()

}

type MyUser struct {
}

type MyUser1 MyUser
type MyUser2 = MyUser

func (i MyUser1) echo1() {
	fmt.Println("MyUser1's echo1")
}
func (i MyUser) echo2() {
	fmt.Println("MyUser's echo2")
}
