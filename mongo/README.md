# mongodb 使用笔记

Go Driver有两个系列的类型表示BSON数据：**D系列类型**和**Raw系列类型**。一般使用**D类型**

**D系列**的类型使用原生的Go类型简单地构建BSON对象。这可以非常有用的来创建传递给MongoDB的命令。 D系列包含4种类型：  
\- **D**：一个BSON文档。这个类型应该被用在顺序很重要的场景， 比如MongoDB命令。  
\- **M**: 一个无序map。 它和D是一样的， 除了它不保留顺序。  
\- **A**: 一个BSON数组。  
\- **E**: 在D里面的一个单一的子项。  

```
bson.D{
		bson.E{
			Key: "name",
			Value: bson.D{
				{"$in", bson.A{"Alice", "Bob", "Cat"}}},
		},
	}
bson.M{
		"name":"Alice",
		"age": 18,
	}
```

- [mongodb官方文档](https://docs.mongodb.com/manual/)
- [mongo-driver官方文档](https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo?tab=doc)
- [实时聚合千万数据](https://juejin.im/post/5e0b58de6fb9a0481467cc5f)