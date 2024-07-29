package main

import (
	"geekbang/basic-go/02_webook/internal/web"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

func main() {
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

	u := web.NewUserHandler()
	u.RegisterRoutesV1(s.Group("/users"))
	//u.RegisterRoutes(s)
	err := s.Run(":8080")
	if err != nil {
		return
	}

}
