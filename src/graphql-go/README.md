- graphql-go/graphql

- graphql-go/handler<-Schema(RootSchema)<-SchemaConfig<-Query<-ObjectConfig<-Fields

```
rootQuery := graphql.ObjectConfig{
  Name: "someQuery"
  Fields: graphql.Fields{
    "query1": filed.doQuery1(),
    "query2": filed.doQuery2(),
    "query3": filed.doQuery3(),
  }
}

newSchemConfig := graphql.SchemaConfig{
  Query:graphql.newObject(rootQuery)
}
newSchema := graphql.NewSchema(schemaConfig)

newHandler := handler.new(
  &handler.Config{
    Schema: &newSchema,
    Pretty:true
    Graphql:true
  }
)

```
