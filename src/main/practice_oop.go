package main

import "fmt"

type People struct{}

func (p *People) ShowA() {
	fmt.Println("i am people's showA")
}

func (p *People) ShowB() {
	p.ShowA()
	fmt.Println("i am people's showB")
}

//Teacher will inherit funcs from People
type Teacher struct {
	People
}

func (t *Teacher) ShowA() {
	fmt.Println("i am teacher's showA")
}

func PracticeOOP() {
	t := Teacher{}
	t.ShowA()
	t.ShowB()
}
