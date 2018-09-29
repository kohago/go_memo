package main

import (
	"fmt"
)

type TestStruct struct {
	x int
	y int
}

func testStructs() {
	testPointer()

	testStruct()

	testArray()

	testMap()
}

//test Map
func testMap() {
	m := make(map[string]TestStruct)
	m["test1"] = TestStruct{1, 2}
	fmt.Println(m["test1"])

	n := map[string]TestStruct{
		"test11": TestStruct{3, 4},
		"test12": TestStruct{5, 6},
	}
	fmt.Println(n)

	//abrigement target type when
	o := map[string]TestStruct{
		"test13": {13, 14},
		"test14": {15, 16},
	}
	fmt.Println(o["test13"])

	m2 := map[string]int{"one": 1, "two": 2}
	m2["one"] = 100
	fmt.Println("one is:", m2["one"])
	fmt.Println(m2)
	delete(m2, "two")
	fmt.Println(m2)
	fmt.Println("the value: ", m2["two"])
	elem1, ok1 := m2["one"]
	fmt.Println("elem1:", elem1, ",ok1", ok1)

	elem2, ok2 := m2["two"]
	fmt.Println("elem2:", elem2, ",ok2", ok2)
}

//testArray
func testArray() {
	var array [2]string
	array[0] = "Hello"
	array[1] = "World"
	fmt.Println(array[0], array[1])
	fmt.Println(array)

	manyValues := [7]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(manyValues)

	for i, v := range manyValues {
		fmt.Printf("%d,%d\n", i, v)
	}

	slice := manyValues[1:4]
	fmt.Println(slice)
	fmt.Printf("Type:%T,value:%v\n", slice, slice)
	fmt.Printf("Type:%T,value:%v\n", manyValues, manyValues)

	var emptySlice []int
	fmt.Printf("len=%d,cap=%d,value:%v\n", len(emptySlice), cap(emptySlice), emptySlice)

	emptySlice = append(emptySlice, 1, 2, 3)
	fmt.Printf("len=%d,cap=%d,value:%v\n", len(emptySlice), cap(emptySlice), emptySlice)

	pow := make([]int, 10, 20)
	fmt.Printf("len=%d,cap=%d,value:%v\n", len(pow), cap(pow), pow)
	for i := range pow {
		pow[i] = 1 << uint(i)
	}

	//take value only
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}

}

//try pointer
func testPointer() {
	i, j := 42, 2000
	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	fmt.Println(*p)
	*p = *p / 10
	fmt.Println(j)

	fmt.Println(p)
}

//try struct
func testStruct() {
	v := TestStruct{1, 2}
	v1 := TestStruct{x: 1}
	v2 := TestStruct{}

	p := &v
	fmt.Println(p.x)
	fmt.Println(v.x)
	fmt.Println(p)
	fmt.Println(v)
	fmt.Println(v1)
	fmt.Println(v2)
}
