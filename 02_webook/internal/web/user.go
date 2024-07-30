package web

import (
	"errors"
	"fmt"
	"geekbang/basic-go/02_webook/internal/domain"
	"geekbang/basic-go/02_webook/internal/service"
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	svc            *service.UserService
	emailRegexp    *regexp.Regexp
	passwordRegexp *regexp.Regexp
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	// regex pattern const, scope control
	const (
		emailRegexPattern    = "^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$"
		passwordRegexPattern = "^(?=.*[A-Z])(?=.*[a-z])(?=.*\\d)(?=.*[@#$%^&+=]).{8,72}$"
	)
	return &UserHandler{
		svc:            svc,
		emailRegexp:    regexp.MustCompile(emailRegexPattern, regexp.None),
		passwordRegexp: regexp.MustCompile(passwordRegexPattern, regexp.None),
	}
}

func (u *UserHandler) RegisterRoutesV1(ug *gin.RouterGroup) {
	ug.POST("/signup", u.SignUp)
	ug.POST("/login", u.Login)
	ug.POST("/edit", u.Edit)
	ug.GET("/profile", u.Profile)
}

func (u *UserHandler) RegisterRoutes(s *gin.Engine) {
	// group reg
	ug := s.Group("/users")
	ug.POST("/signup", u.SignUp)
	ug.POST("/login", u.Login)
	ug.POST("/edit", u.Edit)
	ug.GET("/profile", u.Profile)
}

func (u *UserHandler) SignUp(c *gin.Context) {

	// req struct
	type SignUpReq struct {
		Email             string `json:"email"`
		Password          string `json:"password"`
		ConfirmedPassword string `json:"confirmPassword"`
	}

	// bind: unmarshal by content-type
	var req SignUpReq
	if err := c.BindJSON(&req); err != nil {
		return
	}

	// validation for email & pwd
	isMatch, err := u.emailRegexp.MatchString(req.Email)
	if err != nil {
		c.String(http.StatusInternalServerError, "System Error.") // 500
		return
	}
	if !isMatch {
		c.String(http.StatusBadRequest, "Invalid Email.") // 400
		return
	}

	if req.Password != req.ConfirmedPassword {
		c.String(http.StatusBadRequest, "Confirmed Password does not match.")
		return
	}

	isMatch, err = u.passwordRegexp.MatchString(req.Password)
	if err != nil {
		c.String(http.StatusInternalServerError, "System Error.") // 500
		return
	}
	if !isMatch {
		c.String(http.StatusBadRequest, "Invalid Password. The password must be greater than 8 characters and include numbers and special characters.") // 400
		return
	}

	// call svc with domain obj
	// mutation no need ConfirmedPassword anymore
	err = u.svc.SignUp(c, domain.User{
		Email:    req.Email,
		Password: req.Password,
	})
	if errors.Is(err, service.ErrUserDuplicateEmail) {
		c.String(http.StatusConflict, "Email Conflict.")
		return
	}

	if err != nil {
		c.String(http.StatusInternalServerError, "System Error.") // 500
		return
	}

	c.String(http.StatusOK, "Sign Up Successfully.")
	fmt.Printf("%+v\n", req)
}

func (u *UserHandler) Login(c *gin.Context) {

}

func (u *UserHandler) Edit(c *gin.Context) {

}

func (u *UserHandler) Profile(c *gin.Context) {

}
