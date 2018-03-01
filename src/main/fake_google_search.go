package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	Web   = fakeSearch("web")
	Video = fakeSearch("video")
	Image = fakeSearch("image")

	Web1   = fakeSearch("web1")
	Video1 = fakeSearch("video1")
	Image1 = fakeSearch("image1")

	Web2   = fakeSearch("web2")
	Video2 = fakeSearch("video2")
	Image2 = fakeSearch("image2")
)

func googleFakeSearch() {
	//to get different random value
	rand.Seed(time.Now().UnixNano())
	start := time.Now()

	//results := googleVersion0("golang")
	//results := googleVersion1("goLang")
	//results := first("goLang", fakeSearch("replica 1"), fakeSearch("replica 2"))
	results := googleVersion2("golang")

	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}

//do the same thing,take the first
func first(query string, replicas ...Search) Result {
	c := make(chan Result)
	searchReplica := func(i int) {
		c <- replicas[i](query)
	}
	for i := range replicas {
		go searchReplica(i)
	}
	return <-c
}

type Result struct {
	msg string
}

type Search func(query string) Result

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Microsecond)
		return Result{fmt.Sprintf("%s result for %q\n", kind, query)}
	}
}

// append is slow
func googleVersion0(query string) (results []Result) {
	results = append(results, Web(query))
	results = append(results, Video(query))
	results = append(results, Image(query))
	return results
}

func googleVersion1(query string) (results []Result) {
	c := make(chan Result)
	go func() {
		c <- Web(query)
	}()
	go func() {
		c <- Video(query)
	}()
	go func() {
		c <- Image(query)
	}()

	timeout := time.After(250 * time.Microsecond)
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("Time out !")
			return
		}
	}
	//it is very obvious to return results
	return
}

func googleVersion2(query string) (results []Result) {
	c := make(chan Result)
	go func() {
		c <- first(query, Web1, Web2)
	}()
	go func() {
		c <- first(query, Video1, Video2)
	}()
	go func() {
		c <- first(query, Image1, Image2)
	}()

	timeout := time.After(250 * time.Microsecond)
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("Time out !")
			return
		}
	}
	//it is very obvious to return results
	return
}
