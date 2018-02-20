# GO study

## Setup

 * Goのインストール(OSX)
  [パッケージ]
  https://golang.org/dl/
  [brew]
  [node ,npm]が必要かもしれない！
  `$ brew install go`

  [GOPATHが設定されていない場合は設定する(bash)]
  `$ go env GOPATH`
　go version
  [goclipe]
  　GOROOT、GoPath,Go toolsのインスタンス
 =>eclipseでGoのHello　World
 
  [gdb]
  　Gun　GDB
  　./configure
    make
    make install
    gdb --version
  [gdbのkey-chainの設定]
    システム、コード信頼
    codesign -s gdb-cert `whick gdb`
  [Eclipse Gdbの設定]  
 =>Debugできるようになる
 
 
  [Google App Engine SDK for Goのインストール]
  https://cloud.google.com/appengine/downloads#Google_App_Engine_SDK_for_Go  
 
＝＞LocalのGAEでHelloWorld

[others]
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
  
  https://qiita.com/silverfox/items/3a50cafc44c25e8c52e3
  