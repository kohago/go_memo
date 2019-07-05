package main

import "fmt"

/**

before finish defer will be executed
defer is FILO
only the last panic will be recover()


 defer panic
*/
func PracticePanic() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("fatal")
		}
	}()

	defer func() {
		panic("defer panic")
	}()

	panic("panic")
}
