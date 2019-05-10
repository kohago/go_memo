package main

import (
	"fmt"
	"sort"
)

type hasArray struct {
	Members []string
}

func testCollection()  {
	//testZeroValue()
	//testArrayFor()
	//testSort()
}

type Zero struct {
	Name string
	Age int
	Married bool
}

func testZeroValue()  {
	var z Zero
	//name:,age:0,married:false
	fmt.Printf("name:%v,age:%v,married:%v",z.Name,z.Age,z.Married)
}

func testArrayFor()  {
	//no necessary to check then nil or length of array or the array field
	var ha1 hasArray
	for i:=range ha1.Members{
		fmt.Printf("ha1 %v:%v ",i,ha1.Members[i])
	}

	var ha2 hasArray
	ha2.Members=[]string{"1","2",}
	for i:=range ha2.Members{
		fmt.Printf("ha2 %v:%v ",i,ha2.Members[i])
	}
}


//Sort interface( Len Swap Less)
//sort.Sort()
type person struct {
	Name string
	Age int
}

type persons []person

func (p persons) Len()  int{
	return len(p)
}

func (p persons) Swap(i,j int)  {
	p[i],p[j]=p[j],p[i]
}

func (p persons) Less(i,j int) bool  {
	return  p[i].Age< p[j].Age
}

func testSort()  {
	ps:=persons{
		{Name:"001",Age:11},
		{Name:"002",Age:9},
		{Name:"003",Age:10},
		{Name:"003",Age:8},
	}
	sort.Sort(ps)
	fmt.Printf("%v",ps)
}