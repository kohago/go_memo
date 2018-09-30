package main

import (
	"firebase.google.com/go"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

var (
	curdTemplate = template.Must(template.ParseFiles("./template/crud.html"))

	signinTemplate = template.Must(template.ParseFiles("./template/signin.html"))

	firebaseConfig = &firebase.Config{
		DatabaseURL:"xxx",
		ProjectID:"xxx",
		StorageBucket:"xx",
	}
)

type templateParams struct {
	Notice  string
	Name    string
	Message string
	Posts   []Post
}

func hello(w http.ResponseWriter, r *http.Request) {
	////redirect
	//if r.URL.Path != "/" {
	//	http.Redirect(w,r,"/",http.StatusFound)
	//	return
	//}

	fmt.Fprintln(w, "Hello World!")

	if envVar := os.Getenv("MY_VAR"); envVar != "" {
		fmt.Fprint(w, "I get an envioremnt var :"+envVar)
	} else {
		fmt.Fprint(w, "I get none envioremnt var")
	}
}

func crud(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	//GET
	params := templateParams{}

	q := datastore.NewQuery("Post").Order("-Posted").Limit(20)
	if _, err := q.GetAll(ctx, &params.Posts); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		params.Notice = "Couldn't get lastest posts. Refresh?"
		curdTemplate.Execute(w, params)
		return
	}

	if r.Method == "GET" {
		curdTemplate.Execute(w, params)
		return
	}

	//POST
	name := r.FormValue("name")
	if name == "" {
		name = "Anonymous Gopher"
	}
	params.Name = name

	if r.FormValue("message") == "" {
		w.WriteHeader(http.StatusBadRequest)
		params.Notice = "No Message!"
		curdTemplate.Execute(w, params)
		return
	}
	params.Message = r.FormValue("message")
	params.Notice = fmt.Sprintf("Thank you for your submissioner, %s!", name)

	//save to db
	post := Post{
		Author:  params.Name,
		Message: params.Message,
		Posted:  time.Now(),
	}

	key := datastore.NewIncompleteKey(ctx, "Post", nil)

	if _, err := datastore.Put(ctx, key, &post); err != nil {
		log.Errorf(ctx, "datastore.Put: %v", err)

		w.WriteHeader(http.StatusInternalServerError)
		params.Notice = "Couldn't add new post. Try again?"
		params.Message = post.Message
		curdTemplate.Execute(w, params)
		return
	}

	params.Posts = append([]Post{post}, params.Posts...)
	curdTemplate.Execute(w, params)
}

func signin(w http.ResponseWriter, r *http.Request)  {

	params := templateParams{}
	ctx := appengine.NewContext(r)

	if r.Method == "GET" {
		signinTemplate.Execute(w, params)
		return
	}

	app,err := firebase.NewApp(ctx,firebaseConfig)
	if err!=nil {
		params.Notice ="Couldn't authenticate.Try logging in again?"
		params.Message = r.FormValue("message")
		signinTemplate.Execute(w,params)
		return
	}
	auth,err := app.Auth(ctx)
	if err!=nil {
		params.Notice ="Couldn't authenticate.Try logging in again?"
		params.Message = r.FormValue("message")
		signinTemplate.Execute(w,params)
		return
	}
	token,err := auth.VerifyIDTokenAndCheckRevoked(ctx,r.FormValue("token"))
	if err!=nil {
		params.Notice ="Couldn't authenticate.Try logging in again?"
		params.Message = r.FormValue("message")
		signinTemplate.Execute(w,params)
		return
	}
	user,err := auth.GetUser(ctx,token.UID)
	if err!=nil {
		params.Notice ="Couldn't authenticate.Try logging in again?"
		params.Message = r.FormValue("message")
		signinTemplate.Execute(w,params)
		return
	}
	params.Notice ="enticated!" + user.DisplayName
	params.Message = r.FormValue("message")
	signinTemplate.Execute(w,params)
	return
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/crud", crud)
	http.HandleFunc("/signin",signin)
	appengine.Main()
}
