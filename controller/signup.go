package controller

import (
	"gin_app/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignUpForm(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", gin.H{})
}

func Signup(c *gin.Context) {
	var signupService = service.SignupService{}
	var userid = signupService.Signup(c.PostForm("userid"), c.PostForm("password"))
	c.HTML(http.StatusCreated, "success.html", gin.H{
		"userid": userid,
	})
}
