## 反射的应用

### 总结一些反射的知识点

- reflect.TypeOf(var)返回值为reflect.Type 类型
- reflect.Type 类型的值的Name()方法，会返回类型的名称，有些没有名称，如指针或切片，会返回空字符串
- reflect.Type 类型的值的Kind()方法，会返回数据类型的名称，如切片、map、指针、结构体、接口、字符串、数组、函数、 数值类型等。
> If you define a struct named Foo, the kind is struct and the type is Foo.
- 如果变量是指针、 map、 slice、 channel 或 array，则可以使用 varType.Elem()查找所包含的类型。
- 如果变量是结构体，可以通过反射来获取结构体中属性的数量，得到reflect.StructField{},StructField中提供了属性的名称、顺序、类型和标签。
> 以上几点可以在type_resolver中查看到相关知识点
- reflect.ValueOf(var) 能够得到变量的实例，如果要修改值的话，参数需为指针，然后调用Elem().Set(newRefVal)来修改值，其中Set()的参数也必须是reflect.Value类型
> 这一点可以在change_value中看到。
- 创建新实例：reflect.New(varType)，参数为reflect.Type类型，然后通过Elem().Set()赋值
- 可以使用 reflect.makelice、 reflect.MakeMap 和 reflect.MakeChan 函数生成切片、map或通道
> 这一点可以在making_without_make中看到。
- 使用 reflect.MakeFunc 创建新函数
> 可以在make_func中查看
- 创建结构体：通过向 reflect.StructOf 函数传递一个 reflect.StructField 实例切片
> 可以在make_struct中查看
- 不能创建新方法，也就意味着不能在运行时使用反射来实现接口。

### 参考资料(代码来源)
- https://medium.com/capital-one-tech/learning-to-use-go-reflection-822a0aed74b7
- https://medium.com/capital-one-tech/learning-to-use-go-reflection-part-2-c91657395066

> 考虑到反射在实际项目中的应用较少，在这里做记录，以防忘掉如何处理。

