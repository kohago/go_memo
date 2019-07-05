package main

import (
	"fmt"
	"reflect"
)

func PracticeCompare() {

	sn1 := struct {
		age  int
		name string
	}{
		age:  1,
		name: "name1",
	}

	sn2 := struct {
		age  int
		name string
	}{
		age:  1,
		name: "name1",
	}

	sn3 := struct {
		name string
		age  int
	}{
		age:  1,
		name: "name1",
	}
	fmt.Println(sn3)

	if sn1 == sn2 {
		fmt.Println("sn1 equal sn2")
	} else {
		fmt.Println("sn1 not equal sn2")
	}

	/**
	  can't compare with sn1 to sn3.compile error
	  only the same struct can compare with.
	  struct with different order of field is different struct
	*/
	//if sn1 == sn3 {
	//	fmt.Println("sn1 equal sn3")
	//} else {
	//	fmt.Println("sn1 not equal sn3")
	//}

	/**
	can use reflect deeepEqual to compare
	structs with different sequence fields are not deep Equals
	*/
	if reflect.DeepEqual(sn1, sn3) {
		fmt.Println("sn1 deep equals sn3")
	} else {
		fmt.Println("sn1 not deep equals sn3")
	}

	sm1 := struct {
		age   int
		names []string
	}{
		age:   1,
		names: []string{"test1", "test2"},
	}
	fmt.Println(sm1)

	sm2 := struct {
		age   int
		names []string
	}{
		age:   1,
		names: []string{"test1", "test2"},
	}
	fmt.Println(sm2)

	/**
	complie error!
	when contains map,slice field. struct can't be compared.
	*/
	//if sm1 == sm2 {
	//	fmt.Println("sm1 not equal sm2")
	//}

}
