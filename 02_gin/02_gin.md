## [Gin](https://github.com/gin-gonic/gin)

### Basics

高性能 HTTP 框架。要点：

- 如何定义路由：参数/通配符
- 如何处理输入输出
- 如何使用 middleware 解决 AOP 问题

**面向切面编程 AOP Aspect-Oriented Programming**：一种编程范式，通过将**横切关注点 cross-cutting concerns** 即一些通用功能 (日志/验证鉴权/限流/断路) 从主要业务逻辑中分离出来，以**提高代码的模块化性、可维护性和重用性**。

TODO: MVC

```go
import (
	"github.com/gin-gonic/gin"
)
```

**Engine**: 逻辑 web 服务器；一个 Go 程序可以创建多个 Engine。Engine 负责**路由注册 (HTTP endpoint → logic) + middleware 接入**。

:construction_worker: 路由注册时使用 IDE 代码提示和源码查看。

:confused: **如何设计路由？使用什么方法？使用路径 or 查询参数？**

:bookmark_tabs: 初学

- 用户是查询数据时使用 GET，参数放到查询参数里，即 `?` 后 k=v。
- 用户是提交数据时使用 POST，参数全部放到 Body 里面。

```go
gin.Default()

go func() {
	r1 := gin.Default()
	r1.Run(":8081")
}()
```

```go
type Engine struct {
	RouterGroup
    ...
}
```

**Context**：上下文，负责**处理请求 + 返回响应**。

```go
type Context struct {
	...
	Request   *http.Request  // HTTP req
	Writer    ResponseWriter // HTTP res
    ...
}
```

**RouterGroup**

```go
// 静态路由：完全匹配的路由。
router.GET("/hello", func(c *gin.Context) {
    c.String(http.StatusOK, "Hello, Static Route!")
})

// 参数路由：在路径中带上了参数的路由。
router.GET("/hello/:name", func(c *gin.Context) {
    name := c.Param("name")
    c.String(http.StatusOK, "Hello, Parameter Route! Your name is "+name)
})

// 通配符路由：任意匹配的路由。
// 注：* 不能单独出现
// router.GET("/hello/*", func(c *gin.Context) {    // nok
router.GET("/hello/*action", func(c *gin.Context) {
    action := c.Param("action")
    c.String(http.StatusOK, "Hello, Wildcard Route! Your action is "+action)
})
```

**获取参数**

```go
router.GET("/user/:id", func(c *gin.Context) {
	// a GET request to /user/john
	id := c.Param("id") // id == "john"
	// a GET request to /user/john/
	id := c.Param("id") // id == "/john/"
})
```

**查询参数**：URL `?` 后的 k=v

```go
r.GET("/order", func(c *gin.Context) {
	id := c.Query("id")
	c.String(http.StatusOK, "hello "+id)
})
```

### Web Interface

对于一个用户模块来说，最先要设计的接口就是：**注册 & 登录 Registration & Login**。之后再考虑编辑/查看用户信息。

定义 UserHandler 将所有和用户有关的路由都定义在了这个 Handler 上，同时定义方法注册路由到 Engine 中。

**Handler**: 本质是一个函数，用于处理 HTTP 请求并生成 HTTP 响应。

**采用分散注册使 main 简洁。使用分组注册避免 path 写错。**





