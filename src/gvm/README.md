GoのVersion管理を楽に

```
<(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
source /Users/xxx/.gvm/scripts/gvm
gvm version
gvn listall
//-B binary only,no source file
gvm install go1.9.5 -B
gvm install go1.11.1 -B
gvm use go1.11.1 --default

go version

///Users/kouha.shu/.gvm/gos/go1.11.1
echo $GOROOT

///Users/xxx/.gvm/pkgsets/go1.11.1/global
echo $GOPATH
```
