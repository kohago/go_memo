# Setup
## set up golang
	https://golang.org/dl/
* in the case OX:<br>

	brew `$ brew install go`
	node ,npm maybe required!
	`$ go env GOPATH`
	go version

##goclipe
* GOROOT、GoPATH,Go tools
* try Hello　World in Eclipse

##gdb
* download Gun　GDB
```
  https://www.gnu.org/software/gdb/download/
  ./configure
  make
  make install
  gdb --version
```

* set gdb's key-chain
	system,code-signed

	codesign -s gdb-cert `whick gdb`

* set Eclipse C++ Gdb
* if you can Debug,Succeed!

* eclipse build Target->build
* run as go application

## Google App Engine SDK for Go
	https://cloud.google.com/appengine/downloads#Google_App_Engine_SDK_for_Go  
 	GAE HelloWorld(app.yaml)

##others
 ```
  $ go get -u github.com/constabulary/gb/...
  $ go get -u github.com/constabulary/gb/cmd/gb-vendor
  $ go get -u github.com/PalmStoneGames/gb-gae
  $ go get -u github.com/themalkolm/gb-run/...
  $ go get -u golang.org/x/tools/cmd/goimports
  $ go get -u github.com/golang/lint/golint
  $ go get -u github.com/favclip/jwg/cmd/jwg
  $ go get -u github.com/favclip/qbg/cmd/qbg
  $ go get -u github.com/favclip/smg/cmd/smg
  $ go get -u github.com/kisielk/errcheck
```

##detail
	https://qiita.com/silverfox/items/3a50cafc44c25e8c52e3

##More party tricks
There are endless ways to use these tools, many presented elsewhere.

```
Chatroulette toy:
golang.org/s/chat-roulette

Load balancer:
golang.org/s/load-balancer

Concurrent prime sieve:
golang.org/s/prime-sieve

Concurrent power series (by McIlroy):
golang.org/s/power-series
```

## Intellij
### plugIn
 - File Watchers
 - CamelCasePlugin
 - go
 - googleCloud
