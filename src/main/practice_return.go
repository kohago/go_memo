package main

import (
	"errors"
	"fmt"
)

var FailError = errors.New("have some trouble, failed")

func DoTheThing(really bool) (err error) {
	if really {
		/**
		err inside will replace err outside,so the outside err will be nil
		*/
		re, err := tryToDo()
		if err != nil || re != "it works" {
			err = FailError
		}
	}
	return err
}

func tryToDo() (string, error) {
	return "", FailError
}

/**
nil
nil
*/
func PracticeReturn() {
	fmt.Println(DoTheThing(true))
	fmt.Println(DoTheThing(false))
}
