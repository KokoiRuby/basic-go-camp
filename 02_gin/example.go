package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// a gin engine
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "hello go")
	})

	// yet another engine
	go func() {
		r1 := gin.Default()
		err := r1.Run(":8081")
		if err != nil {
			return
		}
	}()

	r.POST("/post", func(c *gin.Context) {
		c.String(http.StatusOK, "hello post")
	})

	// http://localhost:8080/users/kokoi
	r.GET("/users/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "hello "+name)
	})

	// http://localhost:8080/views/index.html
	r.GET("/views/*.html", func(c *gin.Context) {
		html := c.Param(".html")
		c.String(http.StatusOK, "hello "+html)
	})

	// http://localhost:8080/order?id=2233
	r.GET("/order", func(c *gin.Context) {
		id := c.Query("id")
		c.String(http.StatusOK, "hello "+id)
	})

	// listen on 0.0.0.0:8080 & start
	// http://localhost:8080/hello
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
