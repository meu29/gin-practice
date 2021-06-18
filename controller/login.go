package controller

import (
	"gin_app/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginForm(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func Login(c *gin.Context) {
	var loginService = service.LoginService{}
	var user = loginService.Login(c.PostForm("userid"), c.PostForm("password"))
	log.Println(user)
	c.Redirect(http.StatusMovedPermanently, "/login")
}
