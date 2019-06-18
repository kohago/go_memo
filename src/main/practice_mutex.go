package main

import (
	"fmt"
	"sync"
)

type UserAges struct {
	ages map[string]int
	sync.Mutex
}

func (u *UserAges) add(name string, age int) {
	u.Lock()
	defer u.Unlock()
	u.ages[name] = age
}

func (u *UserAges) get(name string) int {
	if age, ok := u.ages[name]; ok {
		return age
	}
	return -1
}

/**
get doesn't use mutex lock
there maybe concurrent map read or write error.

but i found none error...
*/
func PracticeMutex() {
	us := UserAges{
		ages: make(map[string]int, 10),
	}

	wg := sync.WaitGroup{}
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(i int) {
			n := string(i)
			us.add(n, i)
			wg.Done()
		}(i)
	}
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(i int) {
			n := string(i)
			fmt.Println(i, us.get(n))
			wg.Done()
		}(i)
	}
	wg.Wait()
}
