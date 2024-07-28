package main

import (
	"geekbang/basic-go/02_webook/internal/web"
	"github.com/gin-gonic/gin"
)

func main() {
	s := gin.Default()
	u := &web.UserHandler{}
	u.RegisterRoutes(s)
	err := s.Run(":8080")
	if err != nil {
		return
	}
}
