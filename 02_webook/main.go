package main

import (
	"geekbang/basic-go/02_webook/internal/repository"
	"geekbang/basic-go/02_webook/internal/repository/dao"
	"geekbang/basic-go/02_webook/internal/service"
	"geekbang/basic-go/02_webook/internal/web"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"time"
)

func main() {
	db := initDB()
	server := initServer()
	user := initUser(db)
	user.RegisterRoutesV1(server.Group("/users"))
	//user.RegisterRoutes(s)
	err := server.Run(":8080")
	if err != nil {
		return
	}
}

func initServer() *gin.Engine {
	s := gin.Default()

	// CORS ← fe as input
	s.Use(cors.New(cors.Config{
		//AllowOrigins: []string{"http://localhost:3000"},
		//AllowMethods: []string{"POST", "GET"},
		AllowHeaders: []string{"Content-Type", "Authorization"},
		// headers exposed to clients
		ExposeHeaders: []string{"Content-Type", "Authorization"},
		// cookie like
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") {
				// dev
				return true
			}
			return strings.Contains(origin, "company.domain.name.com")
		},
		MaxAge: 12 * time.Hour,
	}))
	return s
}

func initUser(db *gorm.DB) *web.UserHandler {
	ud := dao.NewUserDAO(db)
	repo := repository.NewUserRepository(ud)
	svc := service.NewUserService(repo)
	u := web.NewUserHandler(svc)
	return u
}

func initDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:13306)/webook"))
	if err != nil {
		// only panic during init
		panic(err)
	}
	err = dao.InitTable(db)
	if err != nil {
		panic(err)
	}
	return db
}
