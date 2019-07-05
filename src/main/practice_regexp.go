package main

import (
	"fmt"
	"regexp"
)

func PracticeRex() {
	var validID = regexp.MustCompile(`^\s*(=|>|<|>=|<=)\s*[0-9]+\s*$`)

	fmt.Println(validID.MatchString("= 100"))
	fmt.Println(validID.MatchString(" =  100 "))
	fmt.Println(validID.MatchString("=100"))
}
