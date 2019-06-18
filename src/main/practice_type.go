package main

import "fmt"

//.(type) only can be used with interface!
func PracticeType() {

	v := getValue()

	//switch v.(type) {
	//case int:
	//	fmt.Println("i am int")
	//case string:
	//	fmt.Println("i am int")
	//}

	fmt.Println(v)
}

func getValue() int {
	return 1
}
