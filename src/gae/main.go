package main

import (
	"fmt"
	"google.golang.org/appengine"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, r *http.Request) {
	////redirect
	//if r.URL.Path != "/" {
	//	http.Redirect(w,r,"/",http.StatusFound)
	//	return
	//}

	fmt.Fprintln(w, "Hello Wolrd!")

	if envVar := os.Getenv("MY_VAR"); envVar != "" {
		fmt.Fprint(w, "I get an envioremnt var :"+envVar)
	} else {
		fmt.Fprint(w, "I get none envioremnt var")
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	appengine.Main()
}
