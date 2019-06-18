package main

import "fmt"

type SomePeople interface {
	Show()
}

type SomeStudent struct {
}

func (s *SomeStudent) Show() {

}

// golang will verify weather the struct implement the interface's method
// when the verification is ok,a not nil interface will be return.
func liveSomePeople() SomePeople {
	var s *SomeStudent
	return s
}

func liveSomeStudent() *SomeStudent {
	var s *SomeStudent
	return s
}

// golang will verify weather the struct implement the interface's method
// when the verification is ok,a not nil interface will be return.
/**
BBBBB
CCCCC
EEEEE
GGGGG
*/
func PracticeInterface() {
	p := liveSomePeople()
	if p == nil {
		fmt.Println("AAAAA")
	} else {
		fmt.Println("BBBBB")
	}

	s := liveSomeStudent()
	if s == nil {
		fmt.Println("CCCCC")
	} else {
		fmt.Println("DDDDD")
	}

	var vi interface{}
	if vi == nil {
		fmt.Println("EEEEE")
	} else {
		fmt.Println("FFFFF")
	}

	var ep SomePeople
	if ep == nil {
		fmt.Println("GGGGG")
	} else {
		fmt.Println("HHHHH")
	}
}
