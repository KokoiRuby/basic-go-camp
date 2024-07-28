:construction_worker: **目标**

- 基础目标：切片的辅助方法、map 辅助方法，以及用内置 map 封装一个Set 出来。
- 中级目标：设计 List、普通队列、HashMap 和 LinkedMap
- 高级目标：基于树形结构衍生出来的类型、基于跳表衍生出来的类型、beancopier 机制。
- 中级目标：并发工具。

:construction_worker: **项目经历**：Go 泛型工具库，提供高性能的辅助方法 & 数据结构，复用以减少业务代码中的冗余。

1. slice 辅助方法：添加/删除/查找/并集/交集/Map Reduce API
2. map 辅助方法
3. 扩展 map：可接受任意类型 HashMap/TreeMap/LinkedMap
4. List 实现：LinkedList, ArrayList, SkipList
5. Set：HashSet, TreeSet, SortedSet
6. 队列：普通队列，优先级队列
7. Bean 操作辅助类：高性能高扩展 Bean Copier 机制
8. 并发扩展工具：并发队列/并发阻塞队列/并发阻塞优先级队列
9. 协程池

## Hello Go!

`main()` 函数 = 程序的入口，且必须属于 "main" package，否则程序无法运行。

`go run main.go` 执行。

## Basics

### main

每一个 `.go` 文件都需要有一个 package 声明（它的归属）。

- 同一个目录下，`.go` package 声明必须一致。
- 同一个目录下， `.go` 和 `_test.go` 文件的 package 声明可以不一致。

Go 无需 `;` 结束

`main()` 无参无返回。

如果 main() 里面**引用了同一个包的其它方法/类型**，那么要加上对应的文件：`go run main.go hello.go` or `go run .`

### type

**int**, int8, int16, int32, int64；int 多少字节取决于 CPU 

**uint**, uint8, 16, 32, 64；无符号

float32, **64**

**只有同类型才可以进行运算，Go 强类型，没有自动类型转换，必须显式转换。**

`math` 复杂数学计算：**常量极值**，float 只有最大值 & 最小正数。

string 双引号转义，反引号表示 raw 不转义。**不建议手写转义，先写好然后赋值到 IDE 中自动完成转义。**

`utf8.RuneCountInString()` 统计 string 中 rune 数量。

`strings` 包操作字符串: find/replace/cap/sub...

byte alias to uint8，表示一个 ASCII 码，格式化 `%c`，转字节切片 []byte("string")，`bytes` 包

bool true/false `&&`, `||`, `!`

- `!(a && b)` 等价于 `!a || !b`（德·摩根定律）
- `!(a || b)` 等价于 `!a && !b`（德·摩根定律）

变量名大小写控制访问性；类型推断：整数默认 int，浮点默认 float64；`:=` **短变量声明，只能用于局部变量 & 方法内部，左边必须有一个新变量**

常量无法修改值。`iota` 随主动赋值而递增的计数。编译器会自动帮忙计算。

方法/函数签名 sig：name + param_list + ret_list (支持多返回 + 返回值命名 all or nothing)

Go 不支持重载，同一个包内不允许重名

`_` 可用于忽略返回值

对于栈溢出，治标：增加栈大小；治本：优化代码，尽可能快地到递归出口

函数“一等公民”

- 支持赋值给变量（使用函数名） `varName := myFunc` 使用变量名发起用 `varName()`
- 可作为返回值
- 函数内定义局部匿名函数，作用域仅限函数内 (可立刻调用) → 闭包

**闭包：匿名函数 + 关联上下文；使用不当可能会内存泄露，一个对象被闭包引用是不会被 GC 的。**

不定长参数：类型切片。

类栈 defer 延迟调用，在函数退出前必定会执行。

**defer 与闭包：defer 执行的时候怎么确定里面的值？**

- 作为参数传入的：定义 defer 的时候就确定了。
- 作为闭包引入的：执行 defer 对应的方法的时候才确定。

defer 支持修改**命名返回值**

control flow: if/else (局部变量作用域), **for (不要迭代参数地址，循环开始时就确定，拿到的是迭代结束后的值)**/break/continue, switch/case/default (如果不接 bool 表达式，保证每个 case 互斥)

### Bult-in

array/slice，推荐 `sl := make([]type, 0, cap)`，提前预估容量，避免扩容开销（重新开辟内存并搬迁，底层数组发生改变）。遇事不决用切片。

子切片 [start:end) 

map 同样提前预估容量，避免扩容开销

comparable：在编译期、运行的时候能够判断出来元素是不是相等。基本类型 & string 可比；如果元素以及结构体字段可比，那么数组和结构体也可比。

### Interface

接口：一组行为的抽象/规范

面向接口编程，即业务层和实现层都面向抽象

结构体，字段大小写控制访问性

结构体初始化 &Struct{} 或 new(Struct) 返回一个指针 `%p`，使用命名字段初始化

结构体打印 `fmt.Printf("%+v", struct)`

接收者 → 值/指针类型；方法传参都是传值，传指针指向的结构体还是同一个；初学：遇事不决用指针。

初始化声明为值/指针都可以调用结构体方法

结构体字段自引用，只能用指针

衍生类型 = 类型重定义 `type TypeA TypeB` 本质上是两个类型，需要强制类型转换；一般想使用第三方库，但无法修改源码，想扩展该库的结构体，可以使用

类型别名 type TypeA = TypeB 本质上同一个类型，无需强制类型转换

鸭子类型：态类型语言中的一种类型风格，对象的适用性不是基于继承特定的类或实现接口，而是取决于对象的方法和属性。

Goland IDE → 右键 Generate → Implement Method → 选择接口 or 快捷键 `Ctrl + I`

组合：类型作为结构体的匿名成员；类继承，没有多态

- 接口组合接口
- 结构体组合结构体
- 结构体组合结构体指针
- 结构体组合接口

A 组合 B 后，可直接在 A 上调用 B 方法；B 已实现的接口，A 也相应实现；同名方法 B 会覆盖 A

泛型：`type Name[T any] interface {}` where T 表示类型参数，初始化时决定具体的类型

结构体/接口/函数都支持泛型

++ 泛型约束类型，约束只能作用在变量，而不是在参数和返回值类型上。







