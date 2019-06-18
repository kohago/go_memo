package main

import (
	"fmt"
)

type Student struct {
	Name string
	Age  int
}

func PracticeForEach() {
	m := make(map[string]*Student)

	stus := []Student{
		{"one", 1},
		{"two", 2},
		{"three", 3},
	}
	//s is a temp variable,and just a copy
	//can't change the origin value via s
	for _, s := range stus {
		s.Age = s.Age + 10
		m[s.Name] = &s
	}

	for k, v := range m {
		fmt.Printf("key:%v,value:%v\n", k, v)
	}

	for k, v := range stus {
		fmt.Printf("key:%v,value:%v\n", k, v)
	}
}
