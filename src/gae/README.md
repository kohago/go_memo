# gae demo

## create  gcp project

## setup goland,eclipse...IDE
 https://www.jetbrains.com/go/
 
## install golang,set GOROOT,GOPATH,download appengine-go
```
  export GOPATH=$HONE/.go
  go version
  go get -u google.golang.org/appengine/...
```

## download gcloud,authorize it ,and it's components
```
 gcloud init
 gcloud components list
 gcloud components install app-engine-go
```

## run helloWorld, http://localhsot:8080
```
dev_admin.py app.yaml
```

## deploy it to gae
```
gcloud app deploy
gcloud app browser
gcloud app logs tail -s default
```

#Run And Debug in IntelliJ(Can't debug.....)
## Setting:install google clouds tool plugin 
## Run Menu:set up run configuration go app Engine
## Set GOROOT to googeAppEngine GoRoot
