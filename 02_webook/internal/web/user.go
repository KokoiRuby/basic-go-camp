package web

import "github.com/gin-gonic/gin"

type UserHandler struct {
}

func (u *UserHandler) RegisterRoutes(s *gin.Engine) {
	// group reg
	ug := s.Group("/users")
	ug.POST("/signup", u.SignUp)
	ug.POST("/login", u.Login)
	ug.POST("/edit", u.Edit)
	ug.POST("/profile", u.Profile)
}

func (u *UserHandler) SignUp(c *gin.Context) {

}

func (u *UserHandler) Login(c *gin.Context) {

}

func (u *UserHandler) Edit(c *gin.Context) {

}

func (u *UserHandler) Profile(c *gin.Context) {

}
