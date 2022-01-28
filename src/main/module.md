- lists the current module and all its dependencies  
  go list -m all  
  
  golang.org/x/text v0.3.0 // indirect  
  => The indirect comment indicates a dependency is not used directly by this module  
  
  rsc.io/quote v1.5.2  
  rsc.io/quote/v3 v3.1.0  
  Each different major version (v1, v2, and so on) of a Go module uses a different module path
  
      
- go get xxxv0.1.1   download upgrade
  go get rsc.io/sampler@v1.3.1
  
- go list -m -versions rsc.io/sampler  
  list version
  
- go doc rsc.io/quote/v3  
  show package doc
  
-  go mod tidy  
   remove unused.(tidy:綺麗な、きちんと) 
-  
