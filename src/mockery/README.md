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
  visitor.VisitWalk(iface)
   ->generator.Generate()
    generate prologeNote
    generate packageName
    generate imports:nameToPackagePath
        {
          method's parameters and return results を見て、必要なものをimportする
          g.addImportsFromTuple(ftype.Params())
		      g.addImportsFromTuple(ftype.Results())
        }
```

- mock生成の注意点
 packageが違うから、ぜんぜん別物と認識される。
 have mocks.fromRecord(*repository.TaskRecord) *domain.Task
 want repository.fromRecord(*repository.TaskRecord) *domain.Task


- when cann't find interface,mockery will be stuck....
- when cann't found some import path ,mockery will print the real path for import
  that will be a compile error

  import xxx "_/Users/xx/....."
