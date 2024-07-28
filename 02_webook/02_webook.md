## WeBook

### Login

对于一个用户模块来说，最先要设计的接口就是：**注册 & 登录 Registration & Login**。之后再考虑编辑/查看用户信息。

定义 UserHandler 将所有和用户有关的路由都定义在了这个 Handler 上，同时定义方法注册路由到 Engine 中。

**采用分散注册使 main 简洁。使用分组注册避免 path 写错。**

```bash
.
├── internal         # 业务代码
│   └── web
│       └── user.go
├── main.go          # 入口
└── pkg              # 归档
```

:warning: http://localhost:3000/users/signup

```bash
$ git clone https://gitee.com/geektime-geekbang_admin/geektime-basic-go.git
$ git checkout week3
$ cd webook-fe
$ npm install
$ npm run dev
```

:construction_worker: 需要和前端确认是用什么格式进行数据传输的，比如 JSON | XML

使用一个结构体来封装来向请求，通过 `Bind[Content-Type]` 进行反序列化。

```go
type SignUpReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var req SignUpReq
if err := c.BindJSON(&req); err != nil {
	return
}
```

:construction_worker: 校验什么字段、怎么校验，理论上由产品经理决定，即便前端进行了校验，但还是不能完全信任，后端还是要有保障。

:construction_worker: Go 原生的 regex 库支持有限，使用 `regexp2`。Handler 构造器中编译好 regex 提高性能，而不是每次访问的时候都要编译。

```go
type UserHandler struct {
	emailRegexp    *regexp.Regexp
	passwordRegexp *regexp.Regexp
}

func NewUserHandler() *UserHandler {
	// regex pattern const, scope control
	const (
		emailRegexPattern    = "^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$"
		passwordRegexPattern = "^(?=.*[A-Z])(?=.*[a-z])(?=.*\\d)(?=.*[@#$%^&+=]).{8,}$"
	)
	emailRegexp := regexp.MustCompile(emailRegexPattern, regexp.None)
	passwordRegexp := regexp.MustCompile(passwordRegexPattern, regexp.None)
	return &UserHandler{
		emailRegexp:    emailRegexp,
		passwordRegexp: passwordRegexp,
	}
}
```

### CORS

**跨域请求**：访问不同域（protocol/domain/port）的资源时会遇到的安全限制问题，如果不做任何处理，一般无法发送成功。







