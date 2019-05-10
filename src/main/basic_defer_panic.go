package main

import "fmt"

//use [panic defer recover ] insidely.
//use error interface outsidely
//recover can get the panic's error content.
func testDefer()  {
	fmt.Print("start\n")
	defer func() {
		fmt.Print("End\n")
		if err:=recover();err!=nil{
			fmt.Printf("Recover!:",err)
		}
	}()

	panic("Panic\n")
}
