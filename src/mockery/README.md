```
config-->parameters,use golang's flag to get parameters

vistor-->対象Interfaces

walker.Walk(vistor)->generate mockers

 //歩き回し、必要な情報を収集する
 doWalk{
       Parser(loader,builder)-->use golang's tools(
         os,path,astなど、Programを構造的に構成、解析する
         ) to load and build go source file

      folder continue( if rescure gointo)
      .git,_tempなどcontinue
      !.go,_test.go  continue
      OK＝＞parse(file)
    }
//対象interfaceをParseする

```
