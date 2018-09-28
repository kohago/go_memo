package main

import (
	"fmt"
	"google.golang.org/appengine"
	"net/http"
)

func handle(w http.ResponseWriter,r *http.Request) {
	if r.URL.Path != "/" {
		http.Redirect(w,r,"/",http.StatusFound)
		return
	}

	fmt.Fprintln(w,"Hello Wolrd!")
}

func main() {
	http.HandleFunc("/",handle)
	appengine.Main()
}