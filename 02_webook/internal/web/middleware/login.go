package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginMiddleWareBuilder struct {
	paths []string
}

func NewLoginMiddleWareBuilder() *LoginMiddleWareBuilder {
	return &LoginMiddleWareBuilder{}
}

// IgnorePaths order-agnostic
func (builder *LoginMiddleWareBuilder) IgnorePaths(path string) *LoginMiddleWareBuilder {
	builder.paths = append(builder.paths, path)
	return builder
}

func (builder *LoginMiddleWareBuilder) Build() gin.HandlerFunc {
	return func(c *gin.Context) {
		// no need to validate on paths
		for _, path := range builder.paths {
			if c.Request.URL.Path == path {
				return
			}
		}

		//if c.Request.URL.Path == "/users/login" || c.Request.URL.Path == "/users/signup" {
		//	return
		//}

		sess := sessions.Default(c)
		id := sess.Get("userId")
		if id == nil {
			c.AbortWithStatus(http.StatusUnauthorized) // 401
			return
		}
	}
}
