### Prerequisite [Win11]

Install [Chocolatey](https://chocolatey.org/install)

```go
$ go env -w CGO_ENABLED=1
$ choco install mingw
$ gcc -v
```

## [GORM](https://github.com/go-gorm/gorm)

Object-Relational Mapping 一种编程技术，用于**在关系数据库和面向对象编程语言之间建立映射**，从而实现对象和数据库之间的转换。

ORM library for Golang, features

- 支持多数据库（驱动）：MySQL, PostgreSQL, SQLite, SQL Server & TiDB
- 支持简单查询/事务/关联
- 支持 Hook
- 支持自动迁移

:cry: gorm.Open(...).* CRUD 接口

[Model Definition](https://gorm.io/zh_CN/docs/models.html): GORM 通过将 Go struct 映射到 DB Table 来简化数据库交互。

```go
type MyModel struct {
	gorm.Model // extra fields
	...
}

// struct tag, reflect to get
type Model struct {
	ID        uint `gorm:"primarykey"`   // primary key
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt DeletedAt `gorm:"index"`   // soft deletion as tombstone & index on
}
```