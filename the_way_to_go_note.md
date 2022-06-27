> 前言：这几天想系统的学习一下go，所以一直在找一本合适都书，发现都不是很好。直到看到《the way to go》,发现它如此但适合自己。
> 为什么这么推荐这本书：1. 这本书写但很详细,原理方面多有提及  2. 结构构清晰


#### 第三章、编辑器、ID、开发等工具

3.3 调试器 
在 fmt.Printf 中使用下面的说明符来打印有关变量的相关信息：（[详细参考手册：](https://learnku.com/docs/the-way-to-go/33-debugger/3575)）
`fmt.Printf("%+v", market)` %+v 打印包括字段在内的实例的完整信息
`fmt.Printf("%#v", market)` %#v 打印包括字段和限定类型名称在内的实例的完整信息
`fmt.Printf("%T", market)`  %T  打印某个类型的完整说明

3.2 构建并运行 Go 程序
```
gofmt -w //格式化原文件
gofmt -r '(a) -> a' –w *.go  // gofmt 也可以通过在参数 -r 后面加入用双引号括起来的替换规则实现代码的简单重构
go build 编译并安装自身包和依赖包
go install 安装自身包和依赖包
```

3.3-3.9 其它
```
go doc 工具会从 Go 程序和包文件中提取顶级声明的首行注释以及每个对象的相关注释，并生成相关文档。
go install 是安装 Go 包的工具，类似 Ruby 中的 rubygems。主要用于安装非标准库的包文件，将源代码编译成对象文件。
go fix  用于将你的 Go 代码从旧的发行版迁移到最新的发行版
go test 是一个轻量级的单元测试框架（第 13 章）
```

#### 第四章 基本结构和基本数据类型
4.24 类型
类型可以是基本类型，如：int、float、bool、string；结构化的（复合的），如：struct、array、slice、map、channel；只描述类型的行为的，如：interface。
结构化的类型使用nil作为默认值

4.3 常量
```
常量使用关键字 const 定义，用于存储不会改变的数据。
存储在常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型。
```

4.4-4.6 变量
1. 变量被声明后系统会赋予默认值：：int 为 0，float 为 0.0，bool 为 false，string 为空字符串，指针为 nil

2. 变量的作用域：
全局变量在函数外，局部变量一般可以通过代码块来判断   
当代码块内外有一个相同变量名时，代码块内的变量临时会起作用外部的会临时隐藏 结束时：内部的同名变量被释放，外部的又会重新显示

3. 变量转换类型转换（参考4.26）valueOfTypeB = typeB(valueOfTypeA)
```
a = string(s)  //string 转换
b := uint64(0) //来同时完成类型转换和赋值操作
```

4. int/unit 和float
```
int/unit(8-64)int类型默认是int64/32 依机器类型而定
float：
只有float32 和 float64
float32（+- 1e-45 -> +- 3.4 * 1e38）
float64（+- 5  1e-324 -> 107  1e308）

注意：当从一个取值范围较大的转换到取值范围较小的类型时（例如将 int32 转换为 int16 或将 float32 转换为 int），会发生精度丢失（截断）的情况
```

5. 字符串处理包 strings 和 strconv包 （[用法参考此链接或官方手册](https://learnku.com/docs/the-way-to-go/strings-and-strconv-packages/3588))



#### 第九章 包

注：原文这一块讲的不是很好并且过时不适用了-- 所以原文可以忽略不看
所以参考的网络文章：https://cloud.tencent.com/developer/article/1859833

Go1.12 版本后, 使用go modules进行包管理


#### 第十章、十一章 <结构体方法>和<接口与反射>观看视频更加高效
主要了解go里面"类"、方法、继承等如何使用


#### 第十三章 错误处理与测试
> 通过学习这一章，我们得会知道如何优雅的处理程序的错误。

**13.1 错误处理**
go中有一个预定义的错误
```
type error interface {
    Error() string
}
//自定义错误
var errNotFound error = errors.New("Not found error")


```



**13.2-13.3: 两个关键字 panic和recover**

panic 能够改变程序的控制流，调用 panic 后会立刻停止执行当前函数的剩余代码，并在当前 Goroutine 中递归执行调用方的 defer；
recover 可以中止 panic 造成的程序崩溃。它是一个只能在 defer 中发挥作用的函数,并且可以接收panic传过来的参数信息；

来看下面这个例子，当我们运行这段代码时会发现 main 函数中的 defer 语句并没有执行，执行的只有当前 Goroutine 中的 defer。
```
func main() {
	defer println("in main")
	go func() {
		defer println("in goroutine")
		panic("")
	}()

	time.Sleep(1 * time.Second)
}

$ go run main.go
in goroutine
panic:
```



